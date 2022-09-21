package ap

// Coverage is a criteria that select the channels with the least coverage.
type Coverage struct {
	Weight int `json:"weight"` // Weight of the criteria
}

func (r *Coverage) SetWeight(weight int) {
	r.Weight = weight
}

func (r Coverage) Criteria() string {
	return "Decrease score by radiation power on other channels and its own channel"
}

func (r Coverage) Name() string {
	return "Coverage"
}

func (r Coverage) Select(ScoredChannels map[int]int, UsedChannels map[int]*Channel) (map[int]int, error) {
	for channel, channelInfo := range UsedChannels {
		for _, channelCovered := range channelInfo.ChannelsCovered {
			ScoredChannels[channelCovered] -= (channelInfo.overallQuality() * (r.Weight / 100)) / 100
		}
		ScoredChannels[channel] -= (channelInfo.overallQuality() * (r.Weight / 100)) / 100
	}
	return ScoredChannels, nil
}
