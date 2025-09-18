package song

import (
	"math/rand"
	"strings"
)

func LetteredProgression(key string, length int, mode int) string {
	var chordMap map[string][]string 

	// Sets chordMap to major or minor chord names
	if mode == 0 {
		chordMap = MajorKeys 
	} else {
		chordMap = MinorKeys
	}
	// Sets the list of chords based on the key
	chords := chordMap[key]

	// Starts the progression slice. This will be what the function returns
	progression := []string {chords[rand.Intn(7)]}

	// Adds chords until progression is of size length
	for i := 1; i < length; i++ {
		lastChord := progression[len(progression)-1]
		var nextChords []string

		switch lastChord {
		case chords[0]: // I | i
			nextChords = append(nextChords, chords[rand.Intn(7)])
		case chords[1]: // ii | ii°
			nextChords = append(nextChords, chords[4], chords[6])
		case chords[2]: // iii | III
			nextChords = append(nextChords, chords[1], chords[3], chords[5])
		case chords[3]: // IV | iv
			nextChords = append(nextChords, chords[1], chords[4], chords[6])
		case chords[4]: // V | v
			nextChords = append(nextChords, chords[0], chords[5], chords[6])
		case chords[5]: // vi | VI
			nextChords = append(nextChords, chords[1], chords[3])
		case chords[6]: // vii° | VII
			nextChords = append(nextChords, chords[0], chords[4], chords[5])
		}

		// Picks a random chord from nextChords and adds it to progression
		nextChordIndex := rand.Intn(len(nextChords))
		progression = append(progression, nextChords[nextChordIndex])
	}

	return strings.Join(progression, " ")
}


func RomanProgression(length int, mode int) string {
	var chords []string

	// Sets chords to major or minor roman numerals
	if mode == 0 {
		chords = RomanMajorChords
	} else {
		chords = RomanMinorChords
	}
	
	// Starts the progression slice. This will be what the function returns 
	progression := []string {chords[rand.Intn(7)]}

	// Adds chords until progression is of size length
	for i := 1; i < length; i++ {
		lastChord := progression[len(progression)-1]
		var nextChords []string

		switch lastChord {
		case chords[0]: // I | i
			nextChords = append(nextChords, chords[rand.Intn(7)])
		case chords[1]: // ii | ii°
			nextChords = append(nextChords, chords[4], chords[6])
		case chords[2]: // iii | III
			nextChords = append(nextChords, chords[1], chords[3], chords[5])
		case chords[3]: // IV | iv
			nextChords = append(nextChords, chords[1], chords[4], chords[6])
		case chords[4]: // V | v
			nextChords = append(nextChords, chords[0], chords[5], chords[6])
		case chords[5]: // vi | VI
			nextChords = append(nextChords, chords[1], chords[3])
		case chords[6]: // vii° | VII
			nextChords = append(nextChords, chords[0], chords[4], chords[5])
		}

		// Picks a random chord from nextChords and adds it to progression
		nextChordIndex := rand.Intn(len(nextChords))
		progression = append(progression, nextChords[nextChordIndex])
	}

	return strings.Join(progression, " ")
}
