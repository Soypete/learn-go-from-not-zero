package add

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
}

func TestAddTable(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{10, -5, 5},
	}
	tt := tests
	for _, tc := range tt {
		t.Run(fmt.Sprintf("Add(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
