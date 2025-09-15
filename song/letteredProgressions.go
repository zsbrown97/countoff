package song

import (
	"math/rand"
	"strings"
)

var majorMap = map[string][]string {
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

var minorMap = map[string][]string {
	"Am": {"Am","Bdim","C","Dm","Em","F","G"},
	"Em": {"Em","F#dim","G","Am","Bm","C","D"},
	"Bm": {"Bm","C#dim","D","Em","F#m","G","A"},
	"F#m": {"F#m","G#dim","A","Bm","C#m","D","E"},
	"C#m": {"C#m","D#dim","E","F#m","G#m","A","B"},
	"G#m": {"G#m","A#dim","B","C#m","D#m","E","F#"},
	"D#m/Ebm": {"D#m","E#dim","F#","G#m","A#m","B","C#"},
	"A#m/Bbm": {"A#m","B#dim","C#","D#m","E#m","F#","G#"},
	"Fm": {"Fm","Gdim","Ab","Bbm","Cm","Db","Eb"},
	"Cm": {"Cm","Ddim","Eb","Fm","Gm","Ab","Bb"},
	"Gm": {"Gm","Adim","Bb","Cm","Dm","Eb","F"},
	"Dm": {"Dm","Edim","F","Gm","Am","Bb","C"},
}

func ProgressionLetters(key string, length int, mode int) string {
	var chordMap map[string][]string 

	if mode == 0 {
		chordMap = majorMap
	} else {
		chordMap = minorMap
	}

	chords := chordMap[key]

	progression := []string {chords[rand.Intn(7)]}

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

		nextChordIndex := rand.Intn(len(nextChords))
		progression = append(progression, nextChords[nextChordIndex])
	}

	return strings.Join(progression, " ")
}
