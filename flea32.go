package flearand

import "math/bits"

const (
	fleaSeed32   = uint32(0xf1ea5eed)
	flea32Rot1   = 27
	flea32Rot2   = 17
	flea32Rounds = 3
)

// Flea32 is a small 32-bit noncryptographic PRNG.
// See details: http://burtleburtle.net/bob/rand/smallprng.html
//
type Flea32 struct {
	a, b, c, d uint32
}

// New32 instantiates a new 32-bit FLEA generator with a given seed.
func New32(seed uint32) *Flea32 {
	f := &Flea32{
		a: fleaSeed32,
		b: seed,
		c: seed,
		d: seed,
	}

	// Functions with for-loops aren't inlined.
	// See https://github.com/golang/go/issues/14768
	i := 0
loop:
	e := f.a - bits.RotateLeft32(f.b, flea32Rot1)
	f.a = f.b ^ bits.RotateLeft32(f.c, flea32Rot2)
	f.b = f.c + f.d
	f.c = f.d + e
	f.d = e + f.a

	i++
	if i < flea32Rounds {
		goto loop
	}
	return f
}

// Next returns a next random number.
func (f *Flea32) Next() uint32 {
	e := f.a - bits.RotateLeft32(f.b, flea32Rot1)
	f.a = f.b ^ bits.RotateLeft32(f.c, flea32Rot2)
	f.b = f.c + f.d
	f.c = f.d + e
	f.d = e + f.a
	return f.d
}
