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
		progression := song.GetSongStarter(Key, 4, MajorMinor)
		fmt.Println(Key)
		fmt.Println(progression.Chords)
	},
}

func Execute() {
	rootCmd.Execute()
}