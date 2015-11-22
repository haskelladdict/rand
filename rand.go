// package rand implements several different efficient and well behaved RNGs
//
// NOTE: While the RNGs are of pretty decent quality none of them is
// cryptographically strong
//
// (C) Markus Dittrich 2015

package rand

type Rander interface {
	Int64() uint64
}

// NewXR64 constructs a xorshift64* RNG
func NewXR64(seed uint64) *rand64 {
	return &rand64{s: seed}
}

// NewXR1024 constructs a xorshift1024* RNG
// NOTE: We use xorshift64* to initialize the generator's initial state
func NewXR1024(seed uint64) *rand1024 {
	r := NewXR64(seed)
	return initXR1024(r)
}

// NewMers64 constructs a Mersenne-Twister RNG initialized with a single seed
func NewMers64(seed uint64) *mt64 {
	return initMers64(seed)
}

// NewMers64A constructs a new Mersenne-Twister RNG initialized by a key array
func NewMers64Arr(keys []uint64) *mt64 {
	r := NewMers64(19650218)
	return initMers64A(r, keys)
}

// Int64 generates a random number on [0, 2^63-1]
func Int63(r Rander) int64 {
	return (int64)(r.Int64() >> 1)
}

// Float64c generates a random number in the closed interval [0,1]
func Float64c(r Rander) float64 {
	return float64(r.Int64()>>11) * (1.0 / 9007199254740991.0)
}

// Float64 creates a random number in the open interval [0,1)
func Float64(r Rander) float64 {
	return float64(r.Int64()>>11) * (1.0 / 9007199254740992.0)
}
