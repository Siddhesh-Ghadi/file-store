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

func TestGetWordCount(t *testing.T) {
	tests := [] struct {
			input string
			want int
		}{
			{input: "./testdata/example.txt", want: 4}, 
			{input: "./testdata/sample.txt", want: 0}, 
			{input: "./testdata/missing.txt", want: 0}, // on error
		}
	
	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			got, _ := GetWordCount(tc.input)
			if  got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestGetWords(t *testing.T) {
	tests := [] struct {
			input string
			want []string
		}{
			{input: "./testdata/example.txt", want: []string{"hello", "world", "hello,", "world"}}, 
			{input: "./testdata/sample.txt", want: []string{}}, 
			{input: "./testdata/missing.txt", want: nil}, // on error
		}
	
	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			got, _ := GetWords(tc.input)
			if  !reflect.DeepEqual(got,tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}