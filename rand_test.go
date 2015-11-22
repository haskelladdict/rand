package rand

import "testing"

var rng uint64

func BenchmarkRand64(b *testing.B) {
	r := NewXR64(1234)
	for i := 0; i < 100000000; i++ {
		r.Gen()
	}
}

func BenchmarkRand1024(b *testing.B) {
	r := NewXR1024(1234)
	for i := 0; i < 100000000; i++ {
		r.Gen()
	}
}
