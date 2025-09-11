package song

import "math/rand"

func Chord(majorMinor int) string {
	chords := [7][2]string {
		{"I", "i"},
		{"ii", "ii°"},
		{"iii", "III"},
		{"IV", "iv"},
		{"V", "v"},
		{"vi", "VI"},
		{"vii°", "VII"},
	}

	chordIndex := rand.Intn(len(chords))
	return chords[chordIndex][majorMinor]
}