package ap

type Signal struct {
	Weight int `json:"weight"` // Weight of the criteria
}

func (s *Signal) SetWeight(weight int) {
	s.Weight = weight
}

func (s Signal) Criteria() string {
	return "Score channel based on the signal strength"
}

func (s Signal) Name() string {
	return "Signal"
}

func (s Signal) Select(ScoredChannels map[int]int, UsedChannels map[int]*Channel) (map[int]int, error) {
	for channel, channelInfo := range UsedChannels {
		ScoredChannels[channel] -= (channelInfo.overallQuality() * (s.Weight / 100)) / 100
	}
	return ScoredChannels, nil
}

func (c *Channel) overallQuality() int {
	return c.SignalStrength / c.NbOfAccessPoints
}
