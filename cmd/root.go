package cmd

import (
	"github.com/Joffref/wifi-cli/cmd/ap"
	"github.com/Joffref/wifi-cli/cmd/terminal"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wifi-cli",
		Short: "wifi-cli is a command line tool for wifi",
		Long:  `wifi-cli is a command line tool for wifi, it can be used in ap mode or in terminal mode. It is written in go. It based on a practical exercise of the course "RESA" of ESIR.`,
	}
	cmd.AddCommand(terminal.NewTerminalCommand())
	cmd.AddCommand(ap.NewApCommand())
	return cmd
}

func Execute() error {
	return NewRootCommand().Execute()
}
