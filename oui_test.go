package oui_test

import (
	"net"
	"testing"

	"github.com/shadowy-pycoder/oui"
)

func TestVendor(t *testing.T) {
	tests := map[string]bool{
		"ff:ff":             false,
		"ff:ff:ff":          false,
		"00:00:0F":          true,
		"00:00:0f":          true,
		"00:00:00":          true,
		"00:00:00:00:00:00": true,
	}

	for s, ok := range tests {
		v := oui.Vendor(s)
		if ok && v == "" {
			t.Errorf("%q should have vendor but doesn't", s)
		}

		if !ok && v != "" {
			t.Errorf("%q should not have vendor but has %q", s, v)
		}
	}
}

func TestVendorFromMAC(t *testing.T) {
	mac, err := net.ParseMAC("00:00:0f:01:02:03")
	if err != nil {
		t.Errorf("%v", mac)
	}

	if oui.VendorFromMAC(mac) == "" {
		t.Errorf("vendor for %q not found", mac)
	}
}

func TestVendorWithMAC(t *testing.T) {
	mac, err := net.ParseMAC("00:00:0f:01:02:03")
	if err != nil {
		t.Errorf("%v", mac)
	}
	if oui.VendorWithMAC(mac) != "Next_01:02:03" {
		t.Errorf("vendor for %q not found", mac)
	}
	mac, err = net.ParseMAC("00:05:ee:01:02:03")
	if err != nil {
		t.Errorf("%v", mac)
	}
	if oui.VendorWithMAC(mac) != "Vanderbilt_International_Swe_01:02:03" {
		t.Errorf("vendor for %q not found", mac)
	}
}
