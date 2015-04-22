package multierror

import (
	"fmt"
	"testing"
)

func BenchmarkError1(b *testing.B) {
	err := MultiError{}

	err.Add(fmt.Errorf("testing"))

	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}

func BenchmarkError10(b *testing.B) {
	err := MultiError{}

	for i := 0; i < 10; i++ {
		err.Add(fmt.Errorf("testing %d", i))
	}

	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}

func BenchmarkError100(b *testing.B) {
	err := MultiError{}

	for i := 0; i < 100; i++ {
		err.Add(fmt.Errorf("testing %d", i))
	}

	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}

func BenchmarkError1000(b *testing.B) {
	err := MultiError{}

	for i := 0; i < 1000; i++ {
		err.Add(fmt.Errorf("testing %d", i))
	}

	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}
