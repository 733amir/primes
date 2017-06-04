package primes

// Generate gets number of primes to generate and return then in a slice.
func Generate(primesCount int) []int {
	primes := make([]int, 0, primesCount)
	for num, pcounter := 2, 0; pcounter < primesCount; num++ {
		isPrime := true
		for i := 0; i < pcounter; i++ {
			if num%primes[i] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			pcounter++
			primes = append(primes, num)
		}
	}
	return primes
}
