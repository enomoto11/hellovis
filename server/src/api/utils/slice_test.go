package utils

import (
	errors2 "errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSlice(t *testing.T) {
	type args[T, O any] struct {
		arg    []T
		mapper func(t T) O
	}
	type testCase[T, O any] struct {
		name     string
		args     args[T, O]
		expected []O
	}
	tests := []testCase[string, string]{
		{
			name: "文字列 -> 文字列",
			args: args[string, string]{
				arg: []string{"foo", "bar"},
				mapper: func(s string) string {
					return fmt.Sprintf("prefix_%s", s)
				},
			},
			expected: []string{"prefix_foo", "prefix_bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, MapSlice(tt.args.arg, tt.args.mapper), tt.expected)
		})
	}
	tests2 := []testCase[int, string]{
		{
			name: "数値 -> 文字列",
			args: args[int, string]{
				arg: []int{0, 1, 2, 3, 4},
				mapper: func(i int) string {
					return strconv.Itoa(i)
				},
			},
			expected: []string{"0", "1", "2", "3", "4"},
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, MapSlice(tt.args.arg, tt.args.mapper), tt.expected)
		})
	}
}

func TestMapSliceWithError(t *testing.T) {
	type args[T, O any] struct {
		arg    []T
		mapper func(t T) (O, error)
	}
	type expects[T any] struct {
		expected []T
		err      error
	}
	type testCase[T, O any] struct {
		name     string
		args     args[T, O]
		expected expects[O]
	}
	tests := []testCase[string, string]{
		{
			name: "文字列 -> 文字列",
			args: args[string, string]{
				arg: []string{"foo", "bar"},
				mapper: func(s string) (string, error) {
					return fmt.Sprintf("prefix_%s", s), nil
				},
			},
			expected: expects[string]{
				expected: []string{"prefix_foo", "prefix_bar"},
				err:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := MapSliceWithError(tt.args.arg, tt.args.mapper)
			assert.Equal(t, actual, tt.expected.expected)
			assert.NoError(t, err)
		})
	}
	tests2 := []testCase[int, string]{
		{
			name: "数値 -> 文字列",
			args: args[int, string]{
				arg: []int{0, 1, 2, 3, 4},
				mapper: func(i int) (string, error) {
					return strconv.Itoa(i), nil
				},
			},
			expected: expects[string]{
				expected: []string{"0", "1", "2", "3", "4"},
				err:      nil,
			},
		},
		{
			name: "エラー",
			args: args[int, string]{
				arg: []int{0, 1, 2, 3, 4},
				mapper: func(i int) (string, error) {
					return strconv.Itoa(i), errors2.New("error")
				},
			},
			expected: expects[string]{
				expected: nil,
				err:      errors2.New("error"),
			},
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := MapSliceWithError(tt.args.arg, tt.args.mapper)
			assert.Equal(t, actual, tt.expected.expected)
			assert.Equal(t, err, tt.expected.err)
		})
	}
}
