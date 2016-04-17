package main

import (
	"fmt"
)

var _Primes []uint64 = []uint64{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
}
var _N int

func _CalcPrimes() {
	N := len(_Primes)
	i := 0
	for n := uint64(101); n < 10000; n += 2 {
		for i = 1; i < N; i++ {
			if n%_Primes[i] == 0 {
				break
			}
		}
		if i == N {
			_Primes = append(_Primes, n)
		}
	}
	N = len(_Primes)

	for n := uint64(10001); n < 100000; n += 2 {
		for i = 1; i < N; i++ {
			if n%_Primes[i] == 0 {
				break
			}
		}
		if i == N {
			_Primes = append(_Primes, n)
		}
	}

	N = len(_Primes)
	_N = N
}
func main() {
	_CalcPrimes()
	fmt.Println(_Primes)
	fmt.Println(100000, "以内的素数个数为", _N)
}
