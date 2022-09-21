package ap

import (
	"github.com/Joffref/wifi-cli/pkg/ap"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ifname string               // WLAN Interface name
var AccessPointNumberWeight int // Weight of the number of access point on the channel
var SignalStrengthWeight int    // Weight of the signal strength
var CoverageWeight int          // Weight of the coverage

func NewApCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ap",
		Short: "ap is a command line tool for wifi in ap mode",
		RunE: func(cmd *cobra.Command, args []string) error {
			if AccessPointNumberWeight < 0 || SignalStrengthWeight < 0 || CoverageWeight < 0 {
				log.Errorf("Weights must be positive")
			}
			if AccessPointNumberWeight > 100 || SignalStrengthWeight > 100 || CoverageWeight > 100 {
				log.Errorf("Weights must be inferior to 100")
			}
			chain := ap.SelectionChain{
				&ap.UnoccupiedChannel{
					Weight: ap.Infinite,
				},
				&ap.AccessPointNumber{
					Weight: AccessPointNumberWeight,
				},
				&ap.Signal{
					Weight: SignalStrengthWeight,
				},
				&ap.Coverage{
					Weight: CoverageWeight,
				},
			}
			chanel, err := BestChanel(ifname, chain)
			if err != nil {
				log.Errorf("Error while finding best channel: %v", err)
				return err
			}
			log.Infof("Best channel is %v", chanel)
			return nil
		},
	}
	cmd.Flags().StringVarP(&ifname, "interface", "i", "wlan0", "wifi interface")
	cmd.Flags().IntVarP(&AccessPointNumberWeight, "AccessPointNumberWeight", "a", 1, "AccessPointNumberWeight")
	cmd.Flags().IntVarP(&SignalStrengthWeight, "SignalStrengthWeight", "s", 1, "SignalStrengthWeight")
	cmd.Flags().IntVarP(&CoverageWeight, "CoverageWeight", "r", 1, "CoverageWeight")
	return cmd
}

func BestChanel(ifname string, chain ap.SelectionChain) (int, error) {
	return ap.FindBestChannel(ifname, chain)
}
