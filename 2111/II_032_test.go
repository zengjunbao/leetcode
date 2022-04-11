package _111

import (
	"fmt"
	"testing"
)

// 给定两个字符串 s 和 t ，编写一个函数来判断它们是不是一组变位词（字母异位词）。
//
// 注意：若 s 和 t 中每个字符出现的次数都相同且字符顺序不完全相同，则称 s 和 t 互为变位词（字母异位词）。
// 输入: s = "anagram", t = "nagaram"
// 输出: true
func isAnagram(s string, t string) bool {
	if s == t {
		return false
	}

	return true
}

func Test032(t *testing.T) {
	s1 := "anagram"
	t1 := "nagaram"
	// value:=  isAnagram(s1,t1)
	fmt.Println("11:", isAnagram22(s1, t1))

	s2 := "a"
	t2 := "a"
	// value:=  isAnagram(s1,t1)
	fmt.Println("22:", isAnagram22(s2, t2))
}


func isAnagram22(s, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}
