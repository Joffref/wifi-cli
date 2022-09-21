package ap

// UnoccupiedChannel is a criteria that select the channel free of other channel coverage?
type UnoccupiedChannel struct {
	Weight int `json:"weight"` // Weight of the criteria. This criteria shall be always 1000000000.
}

// SetWeight sets the weight of the criteria
func (e *UnoccupiedChannel) SetWeight(weight int) {
	e.Weight = weight
}

// Criteria returns the description of the criteria
func (e UnoccupiedChannel) Criteria() string {
	return "Select the channel free of other channel coverage"
}

// Name returns the name of the criteria
func (e UnoccupiedChannel) Name() string {
	return "UnoccupiedChannel"
}

func (e UnoccupiedChannel) Select(ScoredChannels map[int]int, UsedChannels map[int]*Channel) (map[int]int, error) {
	CoveredChannels := make(map[int]bool)
	for i := 1; i <= NumOfChannels; i++ {
		CoveredChannels[i] = false
	}
	for channel, channelInfo := range UsedChannels {
		for _, channelCovered := range channelInfo.ChannelsCovered {
			CoveredChannels[channelCovered] = true
		}
		CoveredChannels[channel] = true
	}
	for channel := range ScoredChannels {
		if !CoveredChannels[channel] {
			ScoredChannels[channel] = e.Weight // This channel is free of other channel coverage
		}
	}
	return ScoredChannels, nil
}
