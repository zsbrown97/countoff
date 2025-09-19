package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zsbrown97/countoff/song"
)

func init() {
	var romanCmd = &cobra.Command{
		Use: "roman",
		Short: "Generates a chord progression using roman numerals",
		Run: func(cmd *cobra.Command, args []string) {
			progression := song.ChordProgression(Key, 4, MajorMinor)[1]

			fmt.Println(Key)
			fmt.Println(Stringify(progression))
		},
	}

	rootCmd.AddCommand(romanCmd)
}