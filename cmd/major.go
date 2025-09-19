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
			progression := song.ChordProgression(key, 4, 0)[0]
			
			fmt.Println(key)
			fmt.Println(Stringify(progression))
		},
	}
	rootCmd.AddCommand(majorCmd)
}
