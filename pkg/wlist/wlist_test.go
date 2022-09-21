package wlist

import "testing"

func TestScan(t *testing.T) {
	cells, err := Scan("some-wlan")
	if err == nil {
		t.Error("scan on undefined interface must be wrong")
	}

	cells, err = Scan("wlp1s0") // replace with your interface name
	if err != nil {
		t.Error(err)
	}

	if len(cells) <= 0 {
		t.Error("scan error")
	}
}
