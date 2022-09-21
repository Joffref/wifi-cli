package ap

type AccessPointNumber struct {
	Weight int `json:"weight"` // Weight of the criteria
}

func (c *AccessPointNumber) SetWeight(weight int) {
	c.Weight = weight
}

func (c AccessPointNumber) Criteria() string {
	return "Score channel based on the number of access points"
}

func (c AccessPointNumber) Name() string {
	return "AccessPointNumber"
}

func (c AccessPointNumber) Select(ScoredChannels map[int]int, UsedChannels map[int]*Channel) (map[int]int, error) {
	for channel, channelInfo := range UsedChannels {
		ScoredChannels[channel] -= (channelInfo.NbOfAccessPoints * (c.Weight / 100)) / 100
	}
	return ScoredChannels, nil
}
