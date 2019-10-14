package bench

import (
	"bytes"
	"testing"
)

func BenchmarkFilterInplace1(b *testing.B) {
	str := bytes.Repeat([]byte("abcx"), 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(str); {
			if str[i] == 'x' {
				str = append(str[:i], str[i+1:]...)
			} else {
				i++
			}
		}
	}
}

func BenchmarkFilterInplace2(b *testing.B) {
	str := bytes.Repeat([]byte("abcx"), 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str2 := str[:0]
		for _, r := range str {
			if r == 'x' {
				continue
			}
			str2 = append(str2, r)
		}
		str = str2
	}
}
