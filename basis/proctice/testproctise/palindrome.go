package testproctise

import (
	"fmt"
	"strings"
)

func checkPalindrome(s string) bool {
	lenS := len(s)
    res := make([]string, lenS)
    for i :=0; i < lenS; i++ {
		res = append(res, string(s[lenS - i - 1]))
	}
	out := strings.Join(res, "")
	if out == s{
		return true
	}
	return false
}