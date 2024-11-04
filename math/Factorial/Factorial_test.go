package factorial_test

import (
	"fmt"
	"testing"

	factorial "github.com/rishavmehra/golangPractice/math/Factorial"
)

func TestFactorial(t *testing.T) {
	t.Run("Factorial of 5", func(t *testing.T) {
		result := factorial.Factorial(5)
		want := 121
		if want != result {
			fmt.Errorf("Result is not corrects")
		}
	})
}
