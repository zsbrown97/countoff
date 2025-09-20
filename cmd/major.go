package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zsbrown97/countoff/song"
)

func init() {
	var majorCmd = &cobra.Command{
		Use: "major",
		Short: "Generates a major key progression",
		Run: func(cmd *cobra.Command, args []string) {
			key := song.GetKeySignature(0)
			progression := song.GetSongStarter(key, 4, 0)
			
			fmt.Println(key)
			fmt.Println(progression.Chords)
		},
	}
	rootCmd.AddCommand(majorCmd)
}
