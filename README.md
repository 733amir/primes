# Golang Tools

I'm newcomer to go programming language. I'm try to learn go and write useful
packages.

## PrimeNumbers

`GeneratePrimes` can generate `N` first prime numbers.

Example:

```go
package main

import (
    "fmt"

    "github.com/733amir/primes"
)

func main() {
    fmt.Println(primes.Generate(25))
}
```
