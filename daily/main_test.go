package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func plusstring(n int, s string) string {
	tmp := ""
	for i := 0; i < n; i++ {
		tmp += s
	}
	return tmp
}
func fmtsprintf(n int, s string) string {
	tmp := ""
	for i := 0; i < n; i++ {
		tmp = fmt.Sprintf("%s%s", tmp, s)
	}
	return tmp
}

func stringbuilder(n int, s string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(s)
	}
	return builder.String()
}

func stringbuilderalloc(n int, s string) string {
	var builder strings.Builder
	builder.Grow(n * len(s))
	for i := 0; i < n; i++ {
		builder.WriteString(s)
	}
	return builder.String()
}

func bytesbuffer(n int, s string) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

func bytesbufferalloc(n int, s string) string {
	var buf bytes.Buffer
	buf.Grow(n * len(s))
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

func bytearray(n int, s string) string {
	tmp := []byte{}
	for i := 0; i < n; i++ {
		tmp = append(tmp, []byte(s)...)
	}
	return string(tmp)
}

func bytearrayprealloc(n int, s string) string {
	tmp := make([]byte, 0, n*len([]byte(s)))
	for i := 0; i < n; i++ {
		tmp = append(tmp, []byte(s)...)
	}
	return string(tmp)
}

func BenchmarkPlusstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusstring(10, "ha")
	}
}

func BenchmarkFmtsprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmtsprintf(10, "ha")
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringbuilder(10, "ha")
	}
}
func BenchmarkStringBuilderalloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringbuilderalloc(10, "ha")
	}
}

func BenchmarkBytesbuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesbuffer(10, "ha")
	}
}
func BenchmarkBytesbufferalloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesbufferalloc(10, "ha")
	}
}
func BenchmarkBytearray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytearray(10, "ha")
	}
}
func BenchmarkBytearrayprealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytearrayprealloc(10, "ha")
	}
}
