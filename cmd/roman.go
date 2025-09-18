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
			fmt.Println(Key)
			fmt.Println(song.RomanProgression(4, MajorMinor))
		},
	}

	rootCmd.AddCommand(romanCmd)
}