package easy

import (
	"testing"
)

func repeatedSubstringPattern(s string) bool {
	var sSize int = len(s)
	switch {
	case sSize == 1:
		return false
	case sSize == 2 && s[0] != s[1]:
		return false
	case sSize == 2 && s[0] == s[1]:
		return true
	default:
		var v string
		for j := 2; j <= sSize; j++ {
			sub_str := s[:j]
			var pivot int
			for i := j; i < sSize; i += j {
				pivot = min(i+j, sSize)
				next_str := s[i:pivot]
				if sub_str != next_str {
					break
				}
				sub_str = next_str
			}
			v = sub_str
			if pivot == sSize {
				break
			}
		}

		var result string
		if sSize%len(v) > 0 {
			result += string(v[0])
		}
		for range sSize / len(v) {
			result += v
		}

		return result == s
	}
}

func TestCase1(t *testing.T) {
	s := "abab"
	expected := true
	result := repeatedSubstringPattern(s)
	if result != expected {
		t.Errorf("\ntest fail: s=%v; result: %v; expected: %v", s, result, expected)
	}
}

func TestCase2(t *testing.T) {
	s := "aba"
	expected := false
	result := repeatedSubstringPattern(s)
	if result != expected {
		t.Errorf("\ntest fail: s=%v; result: %v; expected: %v", s, result, expected)
	}
}

func TestCase3(t *testing.T) {
	s := "abcabcabcabc"
	expected := true
	result := repeatedSubstringPattern(s)
	if result != expected {
		t.Errorf("\ntest fail: s=%v; result: %v; expected: %v", s, result, expected)
	}
}

func TestCase4(t *testing.T) {
	s := "a"
	expected := false
	result := repeatedSubstringPattern(s)
	if result != expected {
		t.Errorf("\ntest fail: s=%v; result: %v; expected: %v", s, result, expected)
	}
}

func TestCase5(t *testing.T) {
	s := "ababba"
	expected := false
	result := repeatedSubstringPattern(s)
	if result != expected {
		t.Errorf("\ntest fail: s=%v; result: %v; expected: %v", s, result, expected)
	}
}

func TestCase6(t *testing.T) {
	s := "abaababaab"
	expected := true
	result := repeatedSubstringPattern(s)
	if result != expected {
		t.Errorf("\ntest fail: s=%v; result: %v; expected: %v", s, result, expected)
	}
}

func TestCase7(t *testing.T) {
	s := "aaaaaaaaaaaaaaaaaaaaa"
	expected := true
	result := repeatedSubstringPattern(s)
	if result != expected {
		t.Errorf("\ntest fail: s=%v; result: %v; expected: %v", s, result, expected)
	}
}

func TestCase8(t *testing.T) {
	s := "ab"
	expected := false
	result := repeatedSubstringPattern(s)
	if result != expected {
		t.Errorf("\ntest fail: s=%v; result: %v; expected: %v", s, result, expected)
	}
}

func TestCase9(t *testing.T) {
	s := "bb"
	expected := true
	result := repeatedSubstringPattern(s)
	if result != expected {
		t.Errorf("\ntest fail: s=%v; result: %v; expected: %v", s, result, expected)
	}
}
