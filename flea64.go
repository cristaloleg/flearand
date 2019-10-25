package flearand

import "math/bits"

const (
	fleaSeed64   = uint64(0xf1ea5eed)
	flea64Rot1   = 39
	flea64Rot2   = 11
	flea64Rounds = 3
)

// Flea64 is a small 64-bit noncryptographic PRNG.
// See details: http://burtleburtle.net/bob/rand/smallprng.html
//
type Flea64 struct {
	a, b, c, d uint64
}

// New64 instantiates a new 64-bit FLEA generator with a given seed.
func New64(seed uint64) *Flea64 {
	f := &Flea64{
		a: fleaSeed64,
		b: seed,
		c: seed,
		d: seed,
	}

	// Functions with for-loops aren't inlined.
	// See https://github.com/golang/go/issues/14768
	i := 0
loop:
	e := f.a - bits.RotateLeft64(f.b, flea64Rot1)
	f.a = f.b ^ bits.RotateLeft64(f.c, flea64Rot2)
	f.b = f.c + f.d
	f.c = f.d + e
	f.d = e + f.a

	i++
	if i < flea64Rounds {
		goto loop
	}
	return f
}

// Next returns a next random number.
func (f *Flea64) Next() uint64 {
	e := f.a - bits.RotateLeft64(f.b, flea64Rot1)
	f.a = f.b ^ bits.RotateLeft64(f.c, flea64Rot2)
	f.b = f.c + f.d
	f.c = f.d + e
	f.d = e + f.a
	return f.d
}
