package ds18b20

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"periph.io/x/devices/v3/ds18b20"
	host "periph.io/x/host/v3"
	"periph.io/x/host/v3/netlink"
)

type Env struct {
	Temperature float64   `json:"temperature"`
	Timestamp   time.Time `json:"timestamp"`
}

type Dev struct {
	dev *ds18b20.Dev
	bus *netlink.OneWire
}

func New() (*Dev, error) {
	d := &Dev{dev: &ds18b20.Dev{}, bus: &netlink.OneWire{}}

	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		return d, err
	}

	// get 1-wire bus
	oneBus, err := netlink.New(001)
	if err != nil {
		return d, err
	}
	log.Debug().Msgf("1wire bus (%#+v)", oneBus)

	// get 1wire address
	addr, err := oneBus.Search(false)
	if err != nil {
		return d, err
	}
	log.Debug().Msgf("1wire address (%#+v)", addr)

	// Open a handle to a ds18b20 connected on the 1-wire bus using default settings
	dev, err := ds18b20.New(oneBus, addr[0], 10)
	if err != nil {
		return d, fmt.Errorf("ds18b20 init: %w", err)
	}
	log.Debug().Msgf("ds18b20 (%#+v)", dev)
	return &Dev{dev: dev, bus: oneBus}, nil
}

func (d *Dev) Read() (Env, error) {
	e := Env{}
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		ds18b20.ConvertAll(d.bus, 10)
		temp, err := d.dev.LastTemp()
		if err == nil {
			e.Temperature = temp.Celsius()
			e.Timestamp = time.Now()
			return e, nil
		}
		log.Info().Msgf("device not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return Env{}, fmt.Errorf("device %v failed to respond after %s", d, timeout)
}
