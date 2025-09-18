package cmd

import (
	"math/rand"

	"github.com/zsbrown97/countoff/song"
)

// Randomizers
var MajorMinor int = rand.Intn(2)
var Key string = song.GetKeySignature(MajorMinor)
