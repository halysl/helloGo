package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}


func TestSplitGroup(t *testing.T) {
	type test struct {
		input string
		sep string
		want []string
	}
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
		}
	}
}

func TestSplitWithChind(t *testing.T) {
	type test struct {
		input string
		sep string
		want []string
	}
	tests := map[string]test{
		"simple": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep": {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}


func BenchmarkSplit(b *testing.B) {
	for i :=0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}
