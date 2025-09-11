package main

import (
	"github.com/zsbrown97/countoff/cmd"
)

func main() {
	cmd.Execute()
	// var majorCmd = &cobra.Command{
	// 	Use: "major",
	// 	Short: "Generate a major key progression",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		songStart(0,4)
	// 	},
	// }

	// var minorCmd = &cobra.Command{
	// 	Use: "minor",
	// 	Short: "Generate a minor key progression",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		songStart(1,4)
	// 	},
	// }


}

// func chord(majorMinor int) string {
// 	chords := [7][2]string{
// 		{"I", "i"},
// 		{"ii", "ii°"},
// 		{"iii", "III"},
// 		{"IV", "iv"},
// 		{"V", "V"},
// 		{"vi", "VI"},
// 		{"vii°", "VII"},
// 	}
// 	chordIndex := rand.Intn(len(chords))
// 	return chords[chordIndex][majorMinor]
// }

// func keySignature(majorMinor int) string {
// 	keys := [12][2]string{
// 		{"C", "Am"},
// 		{"G", "Em"},
// 		{"D", "Bm"},
// 		{"A", "F#m"},
// 		{"E", "C#m"},
// 		{"B", "G#m"},
// 		{"F#", "D#m"},
// 		{"C#", "A#m"},
// 		{"Bb", "Gm"},
// 		{"Eb", "Cm"},
// 		{"Ab", "Fm"},
// 		{"F", "Dm"},
// 	}
// 	keyIndex := rand.Intn(len(keys))
// 	return keys[keyIndex][majorMinor]
// }

// func progression(len int, majorMinor int) string {
// 	progression := ""
// 	for i := 0; i < len; i++ {
// 		progression += chord(majorMinor) + " "
// 	}
// 	return strings.TrimSpace(progression)
// }

// func songStart(m int, i int) {
// 	fmt.Println("Key: ", keySignature(m))
// 	switch m {
// 	case 0:
// 		fmt.Println("Chord Progression: ", progression(i, 0))
// 	case 1:
// 		fmt.Println("Chord Progression: ", progression(i, 1))
// 	}
// }