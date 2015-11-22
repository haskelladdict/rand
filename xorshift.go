// xorshift64* and xorshift1024*.
//
// The implementation was taken from the paper
//
// An experimental exploration of Marsaglia’s xorshift generators, scrambled
// Sebastiano Vigna, Universit`a degli Studi di Milano, Italy
//
// (C) Markus Dittrich 2015

package rand

const (
	mul64   = 2685821657736338717
	mul1024 = 1181783497276652981
)

// rand64 creates random numbers according to the xorshift64* RNG
type rand64 struct {
	s uint64
}

// Gen generates a new random number from the underlying xorshift64* RNG
// in the interval [0, 2^64-1]
func (r *rand64) Int64() uint64 {
	r.s ^= r.s >> 12 // a
	r.s ^= r.s << 25 // b
	r.s ^= r.s >> 27 // c
	return r.s * mul64
}

// rand1024 creates random numbers according to the xorshift1024* RNG
type rand1024 struct {
	s [16]uint64
	p uint64
}

// initXR1024 is a helper function for initializing rand1024 using the rand64 RNG
func initXR1024(r *rand64) *rand1024 {
	var s [16]uint64
	for i := 0; i < len(s); i++ {
		s[i] = r.Int64()
	}
	return &rand1024{s: s}
}

// Int64 creates a new random number from the underlying xorshift1024* RNG
// in the interval [0, 2^64-1]
func (r *rand1024) Int64() uint64 {
	s0 := r.s[r.p]
	r.p = (r.p + 1) & 15
	s1 := r.s[r.p]
	s1 ^= s1 << 31
	s1 ^= s1 >> 11
	s0 ^= s0 >> 30
	r.s[r.p] = s0 ^ s1
	return r.s[r.p] * mul1024
}

// Specialized member functions for maximum performance. This avoids going
// through the Rander interface which slows things down by a factor of two
// for xorshift1024*
// Float64 creates a new random number in the close interval [0,1]
func (r *rand1024) Float64c() float64 {
	return float64(r.Int64()>>11) * (1.0 / 9007199254740991.0)
}

// Float64 creates a new random number in the open interval [0,1)
func (r *rand1024) Float64() float64 {
	return float64(r.Int64()>>11) * (1.0 / 9007199254740992.0)
}
