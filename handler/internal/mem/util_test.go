package mem

import (
	"log"
	"reflect"
	"testing"
)

func TestGetNULTerminated(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []string
	}{
		{
			name: "none",
		},
		{
			name:     "empty",
			input:    []byte{0},
			expected: []string{""},
		},
		{
			name:     "single",
			input:    []byte{'a', 0},
			expected: []string{"a"},
		},
		{
			name:     "double",
			input:    []byte{'a', 0, 'b', 0},
			expected: []string{"a", "b"},
		},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			if want, have := tc.expected, GetNULTerminated(tc.input); !reflect.DeepEqual(want, have) {
				log.Panicf("unexpected strings, want: %q, have: %q", want, have)
			}
		})
	}
}
