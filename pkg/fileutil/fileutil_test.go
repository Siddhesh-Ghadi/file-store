package fileutil

import (
	"testing"
	"reflect"
	"fmt"
)

func TestGetAllFileNames(t *testing.T) {
	tests := [] struct {
			input string
			want []string
		}{
			{input: "./testdata", want: []string{"example.txt", "sample.txt"}}, 
			{input: "./testdata/empty", want: []string{}}, 
			{input: "./testdat", want: nil}, // on error
		}
	
	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			got, _ := GetAllFileNames(tc.input)
			if  !reflect.DeepEqual(got,tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}