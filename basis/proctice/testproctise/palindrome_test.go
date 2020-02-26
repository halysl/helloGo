package testproctise

import "testing"

func TestCheckPalindrome(t *testing.T) {
	type test struct {
		input string
		want bool
	}
	tests := []test{
		{input: "test", want: false},
		{input: "测啊", want: false},
		{input: "A0A", want: true},
	}

	for _, ts := range tests {
		got := checkPalindrome(ts.input)
		if ts.want != got {
			t.Errorf("excepted:%v, got:%v", ts.want, got)
		}
	}
}

