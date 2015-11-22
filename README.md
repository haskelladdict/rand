# rand
rand provides implementations of pseudorandom number generators written in Go

Available Generators
--------------------

* xorshift64*
* xorshift1024*
* 64bit Mersenne-Twister (MT19937-64)

All three are fast and high-quality pseudo random number generators.
If in doubt, choose xorshift1024* since it is the fastest of the lot and provides
the highest quality PRNG of the three. See [this link](http://xorshift.di.unimi.it/)
for an in depth comparison.

Based on benchmarks run on an Intel Broadwell 3.1Ghz i7, xorshift1024* is
about twice as fast as either xorshift64* and MT19937-64. In addition, xorshift1024*
is about three times as fast as Go's built-in default PRGN source.

**NOTE:** For highest performance use the costom Float64() and Float64c() members
of xorshift1024* instead of the interface based functions. Using
the latter will slow things down by a factor of two or more.