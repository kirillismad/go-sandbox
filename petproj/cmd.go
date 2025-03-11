package petproj

import "github.com/spf13/cobra"

var RunServer = &cobra.Command{
	Use:   "petproj",
	Short: "Run petproj",
	Run: func(cmd *cobra.Command, args []string) {
		// Do something
	},
}
