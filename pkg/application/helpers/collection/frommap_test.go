package collection

import (
	"testing"
)

func TestKeysFromMap(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Ensure that the keys are extracted from the map correctly",
			args: args{
				m: map[string]int{
					"key1": 1,
					"key2": 2,
					"key3": 3,
				},
			},
			want: []string{"key1", "key2", "key3"},
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeysFromMap(tt.args.m); !SliceIsEqual(got, tt.want) {
				t.Errorf("KeysFromMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesFromMap(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Ensure that the values are extracted from the map correctly",
			args: args{
				m: map[string]int{
					"key1": 1,
					"key2": 2,
					"key3": 3,
				},
			},
			want: []int{1, 2, 3},
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValuesFromMap(tt.args.m); !SliceIsEqual(got, tt.want) {
				t.Errorf("ValuesFromMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapToSlice(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want []KeyValuePair[string, int]
	}{
		{
			name: "Ensure that the map is converted to a slice correctly",
			args: args{
				m: map[string]int{
					"key1": 1,
					"key2": 2,
					"key3": 3,
				},
			},
			want: []KeyValuePair[string, int]{
				{Key: "key1", Value: 1},
				{Key: "key2", Value: 2},
				{Key: "key3", Value: 3},
			},
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToSlice(tt.args.m); !SliceIsEqual(got, tt.want) {
				t.Errorf("MapToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}