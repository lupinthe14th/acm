package ds18b20

import (
	"testing"

	"periph.io/x/conn/v3/onewire"
	"periph.io/x/devices/v3/ds18b20"
	"periph.io/x/host/v3/netlink"
)

func TestDevString(t *testing.T) {
	d := Dev{
		&ds18b20.Dev{},
		&netlink.OneWire{},
		onewire.Address(0x293ce10457784c28),
	}
	if s := d.String(); s != "293ce10457784c28" {
		t.Errorf("Dev.String() = %s; want 293ce10457784c28", s)
	}
}

func TestGetDevs(t *testing.T) {
	devsData := []Dev{
		{
			dev:  nil,
			bus:  nil,
			addr: 0x293ce10457784c28,
		},
		{
			dev:  nil,
			bus:  nil,
			addr: 0x293ce10457784c29,
		},
	}

	ds := &Devs{
		devs: devsData,
	}

	returnedDevs := ds.GetDevs()

	if len(returnedDevs) != len(devsData) {
		t.Errorf("expected length %d, got %d", len(devsData), len(returnedDevs))
	}

	for i, dev := range returnedDevs {
		if dev.addr != devsData[i].addr {
			t.Errorf("at index %d: expected addr %v, got %v", i, devsData[i].addr, dev.addr)
		}
	}
}
