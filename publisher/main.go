package main

import (
	"context"
	"encoding/json"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"github.com/lupinthe14th/acm/publisher/ds18b20"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Connect to the broker and publish a message periodically
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cfg, err := getConfig()
	if err != nil {
		log.Fatal().Err(err)
	}

	cliCfg := autopaho.ClientConfig{
		BrokerUrls:        []*url.URL{cfg.serverURL},
		KeepAlive:         cfg.keepAlive,
		ConnectRetryDelay: cfg.connectRetryDelay,
		OnConnectionUp:    func(*autopaho.ConnectionManager, *paho.Connack) { log.Info().Msg("mqtt connection up") },
		OnConnectError:    func(err error) { log.Error().Err(err).Msgf("error whilst attempting connection: %s\n", err) },
		Debug:             paho.NOOPLogger{},
		ClientConfig: paho.ClientConfig{
			ClientID:      cfg.clientID,
			OnClientError: func(err error) { log.Printf("server requested disconnect: %s\n", err) },
			OnServerDisconnect: func(d *paho.Disconnect) {
				if d.Properties != nil {
					log.Info().Msgf("server requested disconnect: %s\n", d.Properties.ReasonString)
				} else {
					log.Info().Msgf("server requested disconnect: %d\n", d.ReasonCode)
				}
			},
		},
	}

	if cfg.debug {
		cliCfg.Debug = logger{prefix: "autoPaho"}
		cliCfg.PahoDebug = logger{prefix: "paho"}
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Connect to the broker - this will return immediately after initiating the connection process
	cm, err := autopaho.NewConnection(ctx, cliCfg)
	if err != nil {
		log.Fatal().Err(err)
	}

	var wg sync.WaitGroup

	// Start off a goRoutine that publishes messages
	wg.Add(1)
	go func() {
		defer wg.Done()
		d, err := ds18b20.New()
		if err != nil {
			log.Fatal().Err(err)
		}

		for {
			// AwaitConnection will return immediately if connection is up; adding this call stops publication whilst
			// connection is unavailable.
			err = cm.AwaitConnection(ctx)
			if err != nil { // Should only happen when context is canceled
				log.Info().Msgf("publisher done (AwaitConnection: %s)\n", err)
				return
			}

			e, err := d.Read()
			if err != nil {
				log.Fatal().Err(err)
			}
			// The message could be anything; lets make it JSON containing a simple count (make it simpler to track the messages)
			msg, err := json.Marshal(e)
			if err != nil {
				log.Fatal().Err(err)
			}

			// Publish will block so we run it in a goRoutine
			wg.Add(1)
			go func(msg []byte) {
				defer wg.Done()
				pr, err := cm.Publish(ctx, &paho.Publish{
					QoS:     cfg.qos,
					Topic:   cfg.topic,
					Payload: msg,
				})
				if err != nil {
					log.Error().Err(err).Msg("error publishing")
				} else if pr.ReasonCode != 0 && pr.ReasonCode != 16 { // 16 = Server received message but there are no subscribers
					log.Info().Msgf("reason code %d received\n", pr.ReasonCode)
				} else if cfg.printMessage {
					log.Info().Msgf("sent message: %s\n", msg)
				}
			}(msg)

			select {
			case <-time.After(cfg.delayBetweenMessages):
				log.Info().Msg("delay between messages")
			case <-ctx.Done():
				log.Info().Msg("publisher done")
				return
			}
		}
	}()

	// Wait for a signal before exiting
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	<-sig
	log.Info().Msg("signal caught - exiting")
	cancel()

	wg.Wait()
	log.Info().Msg("shutdown complete")
}
