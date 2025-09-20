package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zsbrown97/countoff/song"
)

func init() {
	var minorCmd = &cobra.Command{
		Use: "minor",
		Short: "Generates a minor key progression",
		Run: func(cmd *cobra.Command, args []string) {
			progression := song.GetSongStarter(MinorKey, 4, 1)
			fmt.Println(MinorKey)
			fmt.Println(progression.Chords)	
		},
	}
	rootCmd.AddCommand(minorCmd)
}
