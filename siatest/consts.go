package siatest

import (
	"github.com/HungMingWu/Sia/crypto"
	"github.com/HungMingWu/Sia/modules"
)

// ChunkSize is a helper method to calculate the size of a chunk depending on
// the minimum number of pieces required to restore the chunk.
func ChunkSize(minPieces uint64) uint64 {
	return (modules.SectorSize - crypto.TwofishOverhead) * minPieces
}
