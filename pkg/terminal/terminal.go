package terminal

import (
	"fmt"
	"github.com/Joffref/wifi-cli/pkg/wlist"
)

// Terminal is a struct for a wifi terminal.
type Terminal struct {
	ESSID   string // ESSID of the terminal.
	Channel int    // Channel of the terminal.
	Quality int    // Quality of the terminal.
}

// ScanBestAccessPoint scans the best access point. Meaning the access point with the best quality.
func ScanBestAccessPoint(ifname string) (Terminal, error) {
	cells, err := wlist.Scan(ifname)
	if err != nil {
		return Terminal{}, err
	}
	best := wlist.Cell{
		SignalQuality: 0,
		Channel:       0,
		ESSID:         "",
	}
	for _, cell := range cells {
		if cell.Channel < 1 || cell.Channel > 13 {
			continue
		}
		if cell.SignalQuality > best.SignalQuality {
			best.ESSID = cell.ESSID
			best.Channel = cell.Channel
			best.SignalQuality = cell.SignalQuality
		}
	}
	if best.SignalQuality == 0 && best.Channel == 0 && best.ESSID == "" {
		return Terminal{}, fmt.Errorf("no access point found")
	}
	return Terminal{
		ESSID:   best.ESSID,
		Channel: best.Channel,
		Quality: best.SignalQuality,
	}, nil
}

// String returns a string representation of the terminal.
// Makes the Terminal type a Stringer.
func (T Terminal) String() string {
	return fmt.Sprintf("ESSID: %v, Channel: %v, SignalStrength: %v", T.ESSID, T.Channel, T.Quality)
}
