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
			key := song.KeySignature(0)
			
			fmt.Println(key)
			fmt.Println(song.MajorProgressionLetters(key,4))
		},
	}

	rootCmd.AddCommand(majorCmd)
}
