package main

import (
	"fmt"
	"os"

	"github.com/kilianp07/AthleteIQBox/entrypoint"
	"github.com/spf13/cobra"
)

const (
	softwareName = "athleteiqbox"
)

func main() {
	var confFile string
	// Define the root command
	var rootCmd = &cobra.Command{
		Use:   softwareName,
		Short: "Run the sofware",
		Run: func(_ *cobra.Command, _ []string) {
			entrypoint.Start(confFile)
		},
	}

	// Adding the --conf flag
	rootCmd.PersistentFlags().StringVar(&confFile, "conf", "", "config file (required)")
	_ = cobra.MarkFlagRequired(rootCmd.PersistentFlags(), "conf")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
