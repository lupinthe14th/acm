package ds18b20

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"periph.io/x/conn/v3/onewire"
	"periph.io/x/devices/v3/ds18b20"
	host "periph.io/x/host/v3"
	"periph.io/x/host/v3/netlink"
)

type Env struct {
	Temperature float64   `json:"temperature"`
	Timestamp   time.Time `json:"timestamp"`
}

type Dev struct {
	dev  *ds18b20.Dev
	bus  *netlink.OneWire
	addr onewire.Address
}

type Devs struct {
	devs []Dev
}

func New() (*Devs, error) {
	ds := &Devs{}

	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		return ds, err
	}

	// get 1-wire bus
	oneBus, err := netlink.New(001)
	if err != nil {
		return ds, err
	}
	log.Debug().Msgf("1wire bus (%#+v)", oneBus)

	// get 1wire address
	addrs, err := oneBus.Search(false)
	if err != nil {
		return ds, err
	}
	log.Debug().Msgf("1wire address (%#+v)", addrs)

	for _, addr := range addrs {
		// Open a handle to a ds18b20 connected on the 1-wire bus using default settings
		dev, err := ds18b20.New(oneBus, addr, 10)
		if err != nil {
			return ds, fmt.Errorf("ds18b20 init: %w", err)
		}
		log.Debug().Msgf("ds18b20 (%#+v)", dev)
		ds.devs = append(ds.devs, Dev{dev: dev, bus: oneBus, addr: addr})
	}
	return ds, nil
}

func (d *Dev) Read() (Env, error) {
	e := Env{}
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		if err := ds18b20.ConvertAll(d.bus, 10); err != nil {
			return e, fmt.Errorf("device %v failed to convert: %w", d, err)
		}
		temp, err := d.dev.LastTemp()
		if err == nil {
			e.Temperature = temp.Celsius()
			e.Timestamp = time.Now()
			return e, nil
		}
		log.Info().Msgf("device not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return e, fmt.Errorf("device %v failed to respond after %s", d, timeout)
}

func (d *Dev) String() string {
	return fmt.Sprintf("%x", d.addr)
}

func (ds *Devs) GetDevs() []Dev {
	return ds.devs
}
