package slice

import (
	"testing"
)

func TestFindIndexInt(t *testing.T) {
	testCases := []struct {
		name   string
		slice  []int
		target int
		want   int
	}{
		{
			name:   "target in slice",
			slice:  []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "target not in slice",
			slice:  []int{1, 2, 3, 4, 5},
			target: 7,
			want:   -1,
		},
		{
			name:   "empty slice",
			slice:  []int{},
			target: 3,
			want:   -1,
		},
		// Add more test cases as needed.
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindIndexInt(tc.slice, tc.target)
			if got != tc.want {
				t.Errorf("findIndexInt(%v, %d) = %d; want %d", tc.slice, tc.target, got, tc.want)
			}
		})
	}
}
