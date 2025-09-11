package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zsbrown97/countoff/song"
)

func init() {
	var minorCmd = &cobra.Command{
		Use: "minor",
		Short: "Generate a minor key progression",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(song.KeySignature(1))
			fmt.Println(song.MinorProgression(4))
		},
	}
	rootCmd.AddCommand(minorCmd)
}
