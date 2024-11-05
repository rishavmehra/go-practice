package factorial_test

import (
	"fmt"
	"testing"

	factorial "github.com/rishavmehra/golangPractice/math/Factorial"
)

func TestFactorialRecursion(t *testing.T) {
	t.Run("Factorial of 5", func(t *testing.T) {
		result := factorial.FactorialRecusion(5)
		want := 120
		if want != result {
			fmt.Errorf("Result is not corrects")
		}
	})
}
