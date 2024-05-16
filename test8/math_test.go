package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type testData struct {
		a    float64
		b    float64
		want float64
	}
	tests := []testData{
		{a: 4, b: 7, want: 11},
		{a: 55, b: 45, want: 100},
	}

	for index, v := range tests {
		v := v
		t.Run(fmt.Sprintf("Test %v", index), func(subT *testing.T) {
			subT.Parallel()
			assert.Equal(subT, Add(v.a, v.b), v.want)
		})
	}
}

func TestSubtract(t *testing.T) {
	assert.Equal(t, Subtract(8, 4), float64(4))
}

func ErrorString(a, b, got, want int) string {
	return fmt.Sprintf("Larger(%d,%d)=%d, want %d", a, b, got, want)
}
