package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zsbrown97/countoff/song"
)

var rootCmd = &cobra.Command{
	Use:   "countoff",
	Short: "A simple CLI for generating chord progressions",
	Run: func(cmd *cobra.Command, args []string) {
		progression := song.ChordProgression(Key, 4, MajorMinor)[0]

		fmt.Println(Key)
		fmt.Println(Stringify(progression))
	},
}

func Execute() {
	rootCmd.Execute()
}