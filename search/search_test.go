package search

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		array  []int
		target int
		want   int
	}{
		{name: "find_last_item", array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 10, want: 9},
		{name: "find_first_item", array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 1, want: 0},
		{name: "find_middle_item", array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 5, want: 4},
		{name: "cannot_find_item", array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 11, want: -1},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Binary(tt.array, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
