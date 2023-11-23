package sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSelection(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name:     "unsorted array",
			args:     args{array: []int{3, 2, 1}},
			expected: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Selection(tt.args.array)
			require.Equal(t, tt.args.array, tt.expected)
		})
	}
}
