package cmd

import "github.com/spf13/cobra"

var (
	authDisableCmd = &cobra.Command{
		Use:   "disable [value]",
		Short: "Disable http auth for a deta program",
		Args:  cobra.MaximumNArgs(1),
		RunE:  disableAuth,
	}
)

func init() {
	authCmd.AddCommand(authDisableCmd)
}

func disableAuth(cmd *cobra.Command, args []string) error {
	return updateAuth(false, args)
}
