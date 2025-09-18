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
		fmt.Println(Key)
		fmt.Println(song.LetteredProgression(Key, 4, MajorMinor))
	},
}

func Execute() {
	rootCmd.Execute()
}