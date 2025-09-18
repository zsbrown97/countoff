package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
	"github.com/zsbrown97/countoff/song"
)

var rootCmd = &cobra.Command{
	Use:   "countoff",
	Short: "A simple CLI for generating chord progressions",
	Run: func(cmd *cobra.Command, args []string) {
		// 0 for major, 1 for minor
		//majorMinor := rand.Intn(2)
		//song.KeySignature(majorMinor)
		majorMinor := rand.Intn(2)
		fmt.Println(song.GetKeySignature(majorMinor))
		switch majorMinor{
		case 0:
			fmt.Println(song.MajorProgression(4))
		case 1:
			fmt.Println(song.MinorProgression(4))
		}
	},
}

func Execute() {
	rootCmd.Execute()
}