package prime

import (
	"context"
	"fmt"
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var TestNumber = int(math.Pow(10, 6))
var TestIndex = int(math.Pow(10, 4))

func TestIsPrime(t *testing.T) {
	Convey("Given some numbers if they are prime", t, func() {
		var number int

		Convey("When the number is negative", func() {
			number = -1

			Convey("Primitive Method", func() {
				So(isPrimePrimitive(number), ShouldEqual, false)
			})
			Convey("Sqrt Method", func() {
				So(isPrimeSqrMethod(number), ShouldEqual, false)
			})
		})
		Convey("When the number is zero", func() {
			number = 0

			Convey("Primitive Method", func() {
				So(isPrimePrimitive(number), ShouldEqual, false)
			})
			Convey("Sqrt Method", func() {
				So(isPrimeSqrMethod(number), ShouldEqual, false)
			})
		})
		Convey("When the number is non prime positive", func() {
			numbers := []int{1, 4, 6, 8, 9, 10, 20, 50, 100, 1000}
			for _, number := range numbers {
				Convey(fmt.Sprintf("%d", number), func() {
					Convey("Primitive Method", func() {
						So(isPrimePrimitive(number), ShouldEqual, false)
					})
					Convey("Sqrt Method", func() {
						So(isPrimeSqrMethod(number), ShouldEqual, false)
					})
				})
			}
		})
		Convey("When the number is prime", func() {
			numbers := []int{2, 3, 5, 7, 11, 23, 31, 41, 53, 97, 997, 1997}
			for _, number := range numbers {
				Convey(fmt.Sprintf("%d", number), func() {
					Convey("Primitive Method", func() {
						So(isPrimePrimitive(number), ShouldEqual, true)
					})
					Convey("Sqrt Method", func() {
						So(isPrimeSqrMethod(number), ShouldEqual, true)
					})
				})
			}
		})
	})
}

func TestNthPrime(t *testing.T) {
	Convey("Given an index for the Nth prime", t, func() {
		Convey("Given an invalid index", func() {
			Convey("Given a negative index", func() {
				Convey("Primitive Method", func() {
					_, e := nthprimePrimitive(-1)
					So(e, ShouldNotBeNil)
				})
				Convey("Sqrt Method", func() {
					_, e := nthprimeSqrMethod(-1)
					So(e, ShouldNotBeNil)
				})
				Convey("Eratosthenes Method", func() {
					_, e := NthprimeEratosthenes(context.TODO(), -1)
					So(e, ShouldNotBeNil)
				})
			})
			Convey("Given a zero index", func() {
				Convey("Primitive Method", func() {
					_, e := nthprimePrimitive(0)
					So(e, ShouldNotBeNil)
				})
				Convey("Sqrt Method", func() {
					_, e := nthprimeSqrMethod(0)
					So(e, ShouldNotBeNil)
				})
				Convey("Eratosthenes Method", func() {
					_, e := NthprimeEratosthenes(context.TODO(), 0)
					So(e, ShouldNotBeNil)
				})
			})
			Convey("Given a too big index", func() {
				Convey("Primitive Method", func() {
					_, e := nthprimePrimitive(math.MaxInt64)
					So(e, ShouldNotBeNil)
				})
				Convey("Sqrt Method", func() {
					_, e := nthprimeSqrMethod(math.MaxInt64)
					So(e, ShouldNotBeNil)
				})
				Convey("Eratosthenes Method", func() {
					_, e := NthprimeEratosthenes(context.TODO(), math.MaxInt64)
					So(e, ShouldNotBeNil)
				})
			})
		})
		Convey("Given a valid index", func() {
			primes := map[int]int{1: 2, 2: 3, 3: 5, 10: 29, 32767: 386083}
			for k, v := range primes {
				Convey(fmt.Sprintf("%dth prime should be %d", k, v), func() {
					Convey("Primitive Method", func() {
						r, e := nthprimePrimitive(k)
						So(e, ShouldBeNil)
						So(v, ShouldEqual, r)
					})
					Convey("Sqrt Method", func() {
						r, e := nthprimeSqrMethod(k)
						So(e, ShouldBeNil)
						So(v, ShouldEqual, r)
					})
					Convey("Eratosthenes Method", func() {
						r, e := NthprimeEratosthenes(context.TODO(), k)
						So(e, ShouldBeNil)
						So(v, ShouldEqual, r)
					})
				})
			}
		})
	})
}

func BenchmarkIsPrimePrimitive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrimePrimitive(TestNumber)
	}
}

func BenchmarkIsPrimeSqrMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrimeSqrMethod(TestNumber)
	}
}

func BenchmarkNthPrimePrimitive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nthprimePrimitive(TestIndex)
	}
}

func BenchmarkNthPrimeSqrtMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nthprimeSqrMethod(TestIndex)
	}
}

func BenchmarkNthPrimeEratosthenes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NthprimeEratosthenes(context.TODO(), TestIndex)
	}
}
