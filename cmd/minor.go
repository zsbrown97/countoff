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
			key := song.GetKeySignature(1)
			progression := song.ChordProgression(key, 4, 1)[0]
			
			fmt.Println(key)
			fmt.Println(Stringify(progression))	
		},
	}
	rootCmd.AddCommand(minorCmd)
}
