package main

import "testing"

func TestParseCSV(t *testing.T) {
	testData := []struct {
		input  string
		output []int
	}{
		{"1,2,3,4,5", []int{1, 2, 3, 4, 5}},
		{"1, 2, 3, 4, 5", []int{1, 2, 3, 4, 5}},
		{"1,2, 3 ,4,5,", []int{1, 2, 3, 4, 5}},
	}

	for _, test := range testData {
		testParseCSV(t, test.input, test.output)
	}
}

var input string = "1,2,3,4,5"
var expected []int = []int{1, 2, 3, 4, 5}

func BenchmarkParseCSV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testParseCSV(b, input, expected)
	}
}

func BenchmarkParseCSVParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testParseCSV(b, input, expected)
		}
	})
}

func testParseCSV(t testing.TB, input string, expected []int) {
	output, err := parseCSV(input)
	if err != nil {
		t.Fatalf("input: %q\n\t%v", input, err)
	}
	if len(expected) != len(output) {
		t.Fatalf("input: %q\n\texp: %v\n\tact: %v", input, expected, output)
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != output[i] {
			t.Fatalf("input: %q\n\texp[%d]: %v\n\tact[%d]: %v", input, i, expected[i], i, output[i])
		}
	}
}
