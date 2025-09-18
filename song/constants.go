package song

import "math/rand"

// Functions
func GetKeySignature(majorMinor int) string {
	keyIndex := rand.Intn(len(KeySignatures))
	return KeySignatures[keyIndex][majorMinor]
}

// Variables
var KeySignatures = [12][2]string {
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

var MajorKeys = map[string][]string {
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

var MinorKeys = map[string][]string {
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

var RomanChords = [7][2]string {
	{"I", "i"},
	{"ii", "ii°"},
	{"iii", "III"},
	{"IV", "iv"},
	{"V", "v"},
	{"vi", "VI"},
	{"vii°", "VII"},
}