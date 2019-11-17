package prime

import (
	"context"
	"errors"
	"fmt"
	"math"
)

var maxIndex = int(math.Pow(10, 10))

func nthprimePrimitive(n int) (int, error) {
	if n < 1 || n > maxIndex {
		return 0, fmt.Errorf("valid indexes are [1..%d]", maxIndex)
	}

	count := 0
	index := 2
	for {
		if isPrimePrimitive(index) {
			count++
			if count == n {
				return index, nil
			}
		}
		index++
	}
}

func nthprimeSqrMethod(n int) (int, error) {
	if n < 1 || n > maxIndex {
		return 0, fmt.Errorf("valid indexes are [1..%d]", maxIndex)
	}

	count := 0
	index := 2
	for {
		if isPrimeSqrMethod(index) {
			count++
			if count == n {
				return index, nil
			}
		}
		index++
	}
}

// NthprimeEratosthenes returns the nth prime number
func NthprimeEratosthenes(ctx context.Context, n int) (int, error) {
	if n < 1 || n > maxIndex {
		return 0, fmt.Errorf("valid indexes are [1..%d]", maxIndex)
	}
	if n < 6 {
		return nthprimeSqrMethod(n)
	}

	return nthPrimesSieveOfEratosthenes(ctx, n)
}

func isPrimePrimitive(number int) bool {
	if number <= 1 {
		return false
	}

	if number > 1 && number <= 3 {
		return true
	}

	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func isPrimeSqrMethod(number int) bool {
	if number <= 1 {
		return false
	}

	if number > 1 && number <= 3 {
		return true
	}

	limit := int(math.Sqrt(float64(number)))
	for i := 2; i <= limit; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func nthPrimesSieveOfEratosthenes(ctx context.Context, n int) (int, error) {
	if n < 1 {
		panic("given index must be >=1")
	}

	realN := float64(n)
	higherLimit := int(realN * math.Log(realN*math.Log(realN)))
	f := make([]bool, higherLimit+1)
	for i := 2; i <= higherLimit; i++ {
		if f[i] == false {
			for j := i * i; j < higherLimit; j += i {
				select {
				case <-ctx.Done():
					fmt.Println("Received context done, giving up.")
					return 0, context.Canceled
				default:
					f[j] = true
				}
			}
		}
	}

	count := 0
	for i := 2; i < len(f); i++ {
		if f[i] == false {
			count++
			if count == n {
				return i, nil
			}
		}
	}

	return 0, errors.New("failed to find the prime number")
}
