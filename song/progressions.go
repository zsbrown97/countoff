package song

import (
	"math/rand"
	"strings"
)

func Progression(len int, majorMinor int) string {
	progression := ""
	for i := 0; i < len; i++ {
		progression += Chord(majorMinor) + " "
	} 
	return strings.TrimSpace(progression)
}

func MajorProgression(length int) string {
	// I ii iii IV V vi vii°
	progression := []string {Chord(0)} 

	for i := 1; i < length; i++ {
		lastChord := progression[len(progression)-1]
		var nextChords []string

		switch lastChord{
		case "I":
			nextChords = append(nextChords, Chord(0))
		case "ii":
			nextChords = append(nextChords,"V", "vii°")
		case "iii":
			nextChords = append(nextChords, "ii", "IV", "vi")
		case "IV":
			nextChords = append(nextChords, "ii", "V", "vii°")
		case "V":
			nextChords = append(nextChords, "I", "vi", "vii°")
		case "vi":
			nextChords = append(nextChords, "ii", "IV")
		case "vii°":
			nextChords = append(nextChords, "I", "V", "vi")
		}

		nextChordIndex := rand.Intn(len(nextChords))
		progression = append(progression, nextChords[nextChordIndex])
	}

	return strings.Join(progression, " ")
}

func MinorProgression(length int) string {
	// i ii° III iv v VI VII 
	progression := []string {Chord(1)} 

	for i := 1; i < length; i++ {
		lastChord := progression[len(progression)-1]
		var nextChords []string

		switch lastChord{
		case "i":
			nextChords = append(nextChords, Chord(1))
		case "ii°":
			nextChords = append(nextChords,"v", "VII")
		case "III":
			nextChords = append(nextChords, "ii°", "iv", "VI")
		case "iv":
			nextChords = append(nextChords, "ii°", "v", "VII")
		case "v":
			nextChords = append(nextChords, "i", "VI", "VII")
		case "VI":
			nextChords = append(nextChords, "ii°", "iv")
		case "VII":
			nextChords = append(nextChords, "i", "v", "VI")
		}

		nextChordIndex := rand.Intn(len(nextChords))
		progression = append(progression, nextChords[nextChordIndex])
	}

	return strings.Join(progression, " ")
}