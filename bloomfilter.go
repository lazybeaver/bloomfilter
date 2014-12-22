// Package bloomfilter implements a standard bloomfilter

package bloomfilter

import (
	"hash/fnv"

	"github.com/lazybeaver/bitset"
)

type BloomFilter struct {
	bits      *bitset.BitSet
	numBits   int
	numHashes int
	numItems  int
}

func hash64(item []byte) uint64 {
	h := fnv.New64()
	h.Write(item)
	return h.Sum64()
}

func getHashes(item []byte) (uint32, uint32) {
	value := hash64(item)
	h1 := uint32(value >> 32)
	h2 := uint32(value & 0xffffffff)
	return h1, h2
}

func getIndex(h1 uint32, h2 uint32, numBits int, hashNum int) int {
	return int((h1 + uint32(hashNum)*h2) % uint32(numBits))
}

func (bf *BloomFilter) Add(item []byte) {
	h1, h2 := getHashes(item)
	for i := 0; i < bf.numHashes; i++ {
		index := getIndex(h1, h2, bf.numBits, i)
		bf.bits.Set(index)
	}
	bf.numItems++
}

func (bf *BloomFilter) Contains(item []byte) bool {
	h1, h2 := getHashes(item)
	for i := 0; i < bf.numHashes; i++ {
		index := getIndex(h1, h2, bf.numBits, i)
		if !bf.bits.Get(index) {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) Size() int {
	return bf.numItems
}

func New(numBits int, numHashes int) *BloomFilter {
	return &BloomFilter{
		bits:      bitset.New(numBits),
		numBits:   numBits,
		numHashes: numHashes,
		numItems:  0,
	}
}
