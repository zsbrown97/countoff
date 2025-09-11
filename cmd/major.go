package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zsbrown97/countoff/song"
)

func init() {
	var majorCmd = &cobra.Command{
		Use: "major",
		Short: "Generate a major key progression",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(song.KeySignature(0))
			fmt.Println(song.MajorProgression(4))
		},
	}
	rootCmd.AddCommand(majorCmd)
}
