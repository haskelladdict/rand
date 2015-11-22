// This file implements a 64 bit version of the Mersenne Twister RNG
//
// The code follows closely the C implementation MT19937-64 by
// by Takuji Nishimura and Makoto Matsumoto.
//

package rand

const (
	NN       = 312
	MM       = 156
	MATRIX_A = 0xB5026F5AA96619E9
	UM       = 0xFFFFFFFF80000000 /* Most significant 33 bits */
	LM       = 0x7FFFFFFF         /* Least significant 31 bits */
	I1       = 6364136223846793005
	I2       = 3935559000370003845
	I3       = 2862933555777941757
)

// mt64 defines the state of a Mersenne-Twister RNG
type mt64 struct {
	mt    [NN]uint64
	mti   uint64
	mag01 [2]uint64
}

// initMers64 is a helper function for initializing mt64 based on a single seed
func initMers64(seed uint64) *mt64 {
	var mti uint64
	r := mt64{}
	r.mt[0] = seed
	for mti = 1; mti < NN; mti++ {
		r.mt[mti] = (I1*(r.mt[mti-1]^(r.mt[mti-1]>>62)) + mti)
	}
	r.mti = mti
	return &r
}

// initMers64A is a helper function for initializing mt64 via a key array
func initMers64A(r *mt64, keys []uint64) *mt64 {
	keyLen := uint64(len(keys))
	var i uint64 = 1
	var j uint64
	var k uint64 = keyLen
	if NN > keyLen {
		k = NN
	}

	for ; k != 0; k-- {
		r.mt[i] = (r.mt[i] ^ ((r.mt[i-1] ^ (r.mt[i-1] >> 62)) * I2)) + keys[j] + j // non linear
		i++
		j++
		if i >= NN {
			r.mt[0] = r.mt[NN-1]
			i = 1
		}
		if j >= keyLen {
			j = 0
		}
	}

	for k = NN - 1; k != 0; k-- {
		r.mt[i] = (r.mt[i] ^ ((r.mt[i-1] ^ (r.mt[i-1] >> 62)) * I3)) - i // non linear
		i++
		if i >= NN {
			r.mt[0] = r.mt[NN-1]
			i = 1
		}
	}

	r.mt[0] = 1 << 63 // MSB is 1; assuring non-zero initial array
	return r
}

// Gen() generates new random numbers number from the mt64 RNG
// in the interval [0, 2^64-1]
func (r *mt64) Int64() uint64 {

	var x uint64

	if r.mti >= NN { // generate NN words at one time
		var i int
		for i = 0; i < NN-MM; i++ {
			x = (r.mt[i] & UM) | (r.mt[i+1] & LM)
			r.mt[i] = r.mt[i+MM] ^ (x >> 1) ^ r.mag01[(int)(x&1)]
		}
		for ; i < NN-1; i++ {
			x = (r.mt[i] & UM) | (r.mt[i+1] & LM)
			r.mt[i] = r.mt[i+(MM-NN)] ^ (x >> 1) ^ r.mag01[(int)(x&1)]
		}
		x = (r.mt[NN-1] & UM) | (r.mt[0] & LM)
		r.mt[NN-1] = r.mt[MM-1] ^ (x >> 1) ^ r.mag01[(int)(x&1)]
		r.mti = 0
	}

	x = r.mt[r.mti]
	r.mti++

	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)

	return x
}
