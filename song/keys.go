package song

import "math/rand"

func KeySignature(majorMinor int) string {
	keys := [12][2]string {
		{"C", "Am"},
		{"G", "Em"},
		{"D", "Bm"},
		{"A", "F#m"},
		{"E", "C#m"},
		{"B", "G#m"},
		{"F#/Gb", "D#m/Ebm"},
		{"C#/Db", "A#m/Bbm"},
		{"Ab", "Fm"},
		{"Eb", "Cm"},
		{"Bb", "Gm"},
		{"F", "Dm"},
	}

	keyIndex := rand.Intn(len(keys))
	return keys[keyIndex][majorMinor]
}