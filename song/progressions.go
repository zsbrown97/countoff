package song

import (
	"math/rand"
)

func ChordProgression(key string, length int, mode int) [][]string {	
	var progressionIndexes = []int{rand.Intn(7)} // Contains index values for lettered and numbered chords
	var chordLetters         []string 			 // Will contain ordered chords for a given key
	var chordNumerals        []string 			 // Will contain roman numerals for the given mode 
	var progressionLetters   []string 			 // Will contain lettered chords based on progressionIndexes
	var progressionNumerals  []string 			 // Will contain chord numbers based on progressionIndexes
	// Return value
	var progressions       [][]string 			 // Will return both progressions

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
		lastChord := progressionIndexes[len(progressionIndexes)-1]
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

		// Selects a random index from nextChords, and pushes it into progressionIndexes
		nextChordIndex := nextChords[rand.Intn(len(nextChords))]
		progressionIndexes = append(progressionIndexes, nextChordIndex)
	}

	// Populates progressionLetters and progressionNumerals based on progressionIndexes
	for _, ind := range progressionIndexes {
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
