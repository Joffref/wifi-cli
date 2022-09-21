package terminal

import (
	"github.com/Joffref/wifi-cli/pkg/terminal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ifname string // WLAN Interface name

func NewTerminalCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "terminal",
		Short: "terminal is a command line tool for wifi in terminal mode",
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := ScanBestAccessPoint(ifname)
			if err != nil {
				log.Errorf("Error while scanning best access point: %v", err)
				return err
			}
			log.Infof("Best access point is %v", t)
			return err
		},
	}
	cmd.Flags().StringVarP(&ifname, "interface", "i", "wlan0", "wifi interface")
	return cmd
}

func ScanBestAccessPoint(ifname string) (terminal.Terminal, error) {
	return terminal.ScanBestAccessPoint(ifname)
}
