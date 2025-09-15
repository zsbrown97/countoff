package song

import (
	"math/rand"
	"strings"
)

func MajorProgressionLetters(key string, length int) string {
	// root +2 +2 +1 +2 +2 +2 +1
	chordMap := map[string][]string {
		"C": {"C","Dm","Em","F","G","Am","Bdim"},
		"G": {"G","Am","Bm","C","D","Em","F#dim"},
		"D": {"D","Em","F#m","G","A","Bm","C#dim"},
		"A": {"A","Bm","C#m","D","E","F#m","G#dim"},
		"E": {"E","F#m","G#m","A","B","C#m","D#dim"},
		"B": {"B","C#m","D#m","E","F#","G#m","A#dim"},
		"F#/Gb": {"F#","G#m","A#m","B","C#","D#m","E#dim"},
		"C#/Db": {"C#","D#m","E#m","F#","G#","A#m","B#dim"},
		"Ab": {"Ab","Bbm","Cm","Db","Eb","Fm","Gdim"},
		"Eb": {"Eb","Fm","Gm","Ab","Bb","Cm","Ddim"},
		"Bb": {"Bb","Cm","Dm","Eb","F","Gm","Adim"},
		"F": {"F","Gm","Am","Bb","C","Dm","Edim"},
	}

	chords := chordMap[key]

	progression := []string {chords[rand.Intn(7)]}

	for i := 1; i < length; i++ {
		lastChord := progression[len(progression)-1]
		var nextChords []string

		switch lastChord {
		case chords[0]: // I
			nextChords = append(nextChords, chords[rand.Intn(7)])
		case chords[1]: // ii
			nextChords = append(nextChords, chords[4], chords[6])
		case chords[2]: // iii
			nextChords = append(nextChords, chords[1], chords[3], chords[5])
		case chords[3]: // IV
			nextChords = append(nextChords, chords[1], chords[4], chords[6])
		case chords[4]: // V
			nextChords = append(nextChords, chords[0], chords[5], chords[6])
		case chords[5]: // vi
			nextChords = append(nextChords, chords[1], chords[3])
		case chords[6]: // vii°
			nextChords = append(nextChords, chords[0], chords[4], chords[5])
		}

		nextChordIndex := rand.Intn(len(nextChords))
		progression = append(progression, nextChords[nextChordIndex])
	}

	return strings.Join(progression, " ")
}