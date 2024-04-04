package collection

import (
	"reflect"
	"testing"
)

type K int // Define the type K

func TestSortSlice(t *testing.T) {
	type args struct {
		slice []K
	}
	tests := []struct {
		name string
		args args
		want []K
	}{
		{
			name: "Ensure that the slice is sorted correctly",
			args: args{
				slice: []K{4, 2, 6, 1, 3, 5},
			},
			want: []K{1, 2, 3, 4, 5, 6},
		},
		// Add more test cases if needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortSlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

type T int // Define the type T

func TestSliceIsEqual(t *testing.T) {
	type args struct {
		aa []T
		bb []T
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Ensure both slices are equal",
			args: args{
				aa: []T{1, 2, 3},
				bb: []T{1, 2, 3},
			},
			want: true,
		},
		{
			name: "Ensure both slices are not equal",
			args: args{
				aa: []T{1, 2, 3},
				bb: []T{1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceIsEqual(tt.args.aa, tt.args.bb); got != tt.want {
				t.Errorf("SliceIsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}