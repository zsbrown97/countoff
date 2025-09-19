package song

import (
	"math/rand"
)

func ChordProgression(key string, length int, mode int) [][]string {
	var progressionIndices = []int{rand.Intn(7)}
	var chordLetters         []string
	var chordNumerals        []string

	// Return values
	var progressions       [][]string
	var progressionLetters   []string
	var progressionNumerals  []string

	// Sets chordMap and chordNumerals to major or minor chord names
	if mode == 0 {
		chordLetters = MajorKeys[key]
		chordNumerals = RomanMajorChords
	} else {
		chordLetters = MinorKeys[key]
		chordNumerals = RomanMinorChords
	}
	
	// Adds chords until progression is of size length
	for i := 1; i < length; i++ {
		lastChord := progressionIndices[len(progressionIndices)-1]
		var nextChords []int

		switch lastChord {
		case 0: // I | i
			nextChords = append(nextChords, rand.Intn(7))
		case 1: // ii | ii°
			nextChords = append(nextChords, 4, 6)
		case 2: // iii | III
			nextChords = append(nextChords, 1, 3, 5)
		case 3: // IV | iv
			nextChords = append(nextChords, 1, 4, 6)
		case 4: // V | v
			nextChords = append(nextChords, 0, 5, 6)
		case 5: // vi | VI
			nextChords = append(nextChords, 1, 3)
		case 6: // vii° | VII
			nextChords = append(nextChords, 0, 4, 5)
		}

		// Selects a random index from nextChords, and pushes it into progressionIndices
		nextChordIndex := nextChords[rand.Intn(len(nextChords))]
		progressionIndices = append(progressionIndices, nextChordIndex)
	}

	for _, ind := range progressionIndices {
		progressionLetters = append(progressionLetters, chordLetters[ind])
		progressionNumerals = append(progressionNumerals, chordNumerals[ind])
	}

	progressions = append(
		progressions,
		progressionLetters,
		progressionNumerals,
	)
	
	return progressions
}
