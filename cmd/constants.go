package cmd

import (
	"math/rand"
	"strings"

	"github.com/zsbrown97/countoff/song"
)

// Functions
func Stringify(progression []string) string {
	return strings.Join(progression, " ")
}

// Randomizers
var MajorMinor int = rand.Intn(2)
var Key        string = song.GetKeySignature(MajorMinor)
