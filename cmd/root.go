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
		key := song.GetKeySignature(majorMinor)

		fmt.Println(key)
		fmt.Println(song.LetteredProgression(key, 4, majorMinor))
	},
}

func Execute() {
	rootCmd.Execute()
}