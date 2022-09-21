package ap

import (
	"fmt"
	"github.com/Joffref/wifi-cli/pkg/wlist"
	log "github.com/sirupsen/logrus"
)

// Channel represents a wifi channel.
type Channel struct {
	Number           int          // Channel number.
	SignalStrength   int          // Sum of quality of the channel, representing the overall signal usage of the channel.
	NbOfAccessPoints int          // Number of access points on the channel.
	Cells            []wlist.Cell // Cells on the channel.
	ChannelsCovered  []int        // Represents the channels that are covered by the channel.
}

const (
	NumOfChannels        = 13             // Number of channels in the 2.4GHz band
	NumOfChannelsCovered = 5              // Number of channels covered by a channel, including itself. For example, channel 1 covers channels 1, 2, 3, 4, 5.
	minusInfinity        = -1000000000000 // A very small number.
	Infinite             = 1000000000000  // A very large number.
)

// SelectionMiddleware is a middleware that selects the best channel given a criteria.
type SelectionMiddleware interface {
	// Select ranks the channels based on the criteria.
	Select(ScoredChannels map[int]int, UsedChannels map[int]*Channel) (map[int]int, error)
	// Criteria returns the criteria of the middleware.
	Criteria() string
	// Name returns the name of the criteria.
	Name() string
	// SetWeight sets the weight of the criteria.
	// The weight is an int between 0 and 100.
	SetWeight(int)
}

type SelectionChain []SelectionMiddleware // A chain of selection middlewares.

// FindBestChannel returns the best channel to use for the given interface.
func FindBestChannel(ifname string, chain SelectionChain) (int, error) {
	log.Infof("Scanning for access points...")
	cells, err := wlist.Scan(ifname)
	if err != nil {
		return 0, err
	}
	usedChannels := UsedChannels(cells)
	log.Infof("Found %v access points on %v channels", len(cells), len(usedChannels))
	for _, channel := range usedChannels {
		log.Infof("Channel %v: %v access points", channel.Number, channel.NbOfAccessPoints)
	}
	scoredChannels := Select(chain, usedChannels)
	max := minusInfinity
	var bestChannel int
	fmt.Println("Scored channels:" + fmt.Sprint(scoredChannels))
	for channel, score := range scoredChannels {
		if score > max {
			max = score
			bestChannel = channel
		}
	}
	return bestChannel, nil

}

func Select(chain SelectionChain, usedChannels map[int]*Channel) map[int]int {
	scoredChannels := make(map[int]int)
	for i := 1; i <= NumOfChannels; i++ {
		scoredChannels[i] = 0
	}
	for _, middleware := range chain {
		log.Infof("Selecting best channel using %v based on %s...", middleware.Name(), middleware.Criteria())
		scoredChannels, _ = middleware.Select(scoredChannels, usedChannels)
	}
	return scoredChannels
}

// UsedChannels returns a map of channels from a list of cells.
func UsedChannels(cells []wlist.Cell) map[int]*Channel {
	channels := make(map[int]*Channel) // Map of channels
	for _, cell := range cells {
		if cell.Channel < 1 || cell.Channel > 13 { // Access point is on a channel that is not in the 2.4GHz band.
			continue
		}
		if _, ok := channels[cell.Channel]; !ok { // Channel is not in the map.
			channels[cell.Channel] = &Channel{
				Number:           cell.Channel,
				SignalStrength:   cell.SignalQuality,
				NbOfAccessPoints: 1,
				Cells:            []wlist.Cell{cell},
				ChannelsCovered:  []int{cell.Channel},
			}
			for i := 1; i <= NumOfChannelsCovered; i++ {
				if cell.Channel+i <= NumOfChannels {
					channels[cell.Channel].ChannelsCovered = append(channels[cell.Channel].ChannelsCovered, cell.Channel+i)
				}
				if cell.Channel-i >= 1 {
					channels[cell.Channel].ChannelsCovered = append(channels[cell.Channel].ChannelsCovered, cell.Channel-i)
				}
			}
		} else {
			channels[cell.Channel].Cells = append(channels[cell.Channel].Cells, cell)
			channels[cell.Channel].SignalStrength += cell.SignalQuality
			channels[cell.Channel].NbOfAccessPoints++
		}
	}
	return channels
}
