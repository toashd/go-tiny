// Go implementation for generating Tiny URL- and bit.ly-like URLs.
// A bit-shuffling approach is used to avoid generating consecutive, predictable URLs.
// However, the algorithm is deterministic and will guarantee that no collisions will occur.
package tiny

import (
	"fmt"
	"math"
	"strings"
)

const (
	defaultAlphabet  = "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"
	defaultBlockSize = 24
	defaultMinLength = 5
)

// TinyUrl generates a tiny url by using a bit-shuffling approach
// to avoid generating consecutive, predictable URLs.
type TinyUrl struct {
	alphabet string
	blockSize int // specifies how many bits will be shuffled
	mask int
	mapping []int
}

// Initializes a new TinyUrl
func New() *TinyUrl {
	t := &TinyUrl{
		alphabet: defaultAlphabet,
		blockSize: defaultBlockSize,
		mask: mask(defaultBlockSize),
		mapping: mapping(defaultBlockSize),
	}
	return t
}

// Calculates mask depending on block size
func mask(blockSize int) int {
	return  (1 << uint64(blockSize)) - 1
}

// Calculates mapping depending on block size
func mapping(blockSize int) []int {
	n := blockSize
	mapping := make([]int, n)
	for i, _ := range mapping {
		n--
		mapping[n] = i
	}
	return mapping
}

// Ecodes the url (id) with min length
func (t *TinyUrl) EncodeUrl(n int, minLength int) string {
	if minLength < defaultMinLength {
		minLength = defaultMinLength
	}
	return t.enbase(t.encode(n), minLength)
}

// Decodes the url
func (t *TinyUrl) DecodeUrl(n string) int {
	return t.decode(t.debase(n))
}

func (t *TinyUrl) encode(n int) int {
	return (n & ^t.mask) | t._encode(n & t.mask)
}

func (t *TinyUrl) _encode(n int) int {
	var result = 0
	for i, m := range t.mapping {
		if n & (1 << uint(i)) != 0 {
			result |= (1 << uint(m))
		}
	}
	return result
}

func (t *TinyUrl) decode(n int) int {
	return (n & ^t.mask) | t._decode(n & t.mask)
}

func (t *TinyUrl) _decode(n int) int {
	var result = 0
	for i, m := range t.mapping {
		if n & (1 << uint(m)) != 0 {
			result |= (1 << uint(i))
		}
	}
	return result
}

// Converts the given integer to a different base
func (t *TinyUrl) enbase(x int, minLength int) string {
	var result = t._enbase(x)
	var padding = strings.Repeat(string(t.alphabet[0]), (minLength - len(result)))
	return fmt.Sprintf("%s%s", padding, result)
}

func (t *TinyUrl) _enbase(x int) string {
	var n = len(t.alphabet)
	if x < n {
		return string(t.alphabet[x])
	}
	return t._enbase(x/n) + string(t.alphabet[x%n])
}

func (t *TinyUrl) debase(x string) int {
	var n = len(t.alphabet)
	var result = 0
	for i, rune := range reverseString(x) {
		result += strings.Index(t.alphabet, string(rune)) * int(math.Pow(float64(n), float64(i)))
	}
	return result
}

// Reverses the given string
func reverseString(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

