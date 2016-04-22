package supersplit

import (
	"testing"
)

func BenchmarkBareSplitChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped("a,b", ",", "\\")
	}
}

func BenchmarkEscapedSplitChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped("a\\,b", ",", "\\")
	}
}

func BenchmarkMixedSplitChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped("a\\,b,c", ",", "\\")
	}
}

func BenchmarkEndingSplitChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped("a,", ",", "\\")
	}
}

func BenchmarkEndingEscapeChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped("a,\\", ",", "\\")
	}
}

func BenchmarkBeginningSplitChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped(",a,", ",", "\\")
	}
}

func BenchmarkBeginningEscapeChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped("\\,a,", ",", "\\")
	}
}

func BenchmarkAdjacentSplitChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped("a,,b", ",", "\\")
	}
}

func BenchmarkAdjacentEscapeChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped("a\\\\,b", ",", "\\")
	}
}

func BenchmarkFalseEscapedChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped("a\\T,b\\F,c", ",", "\\")
	}
}

func BenchmarkFalseEscapedCharsAndEscapedSplitChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escaped("a\\T,b\\F\\,c", ",", "\\")
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"a", "b", "c"}, ",", "\\")
	}
}

func BenchmarkJoinEscapingSeparator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"a,", "b,", "c,"}, ",", "\\")
	}
}

func BenchmarkJoinEscapingEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"a\\", "b\\", "c\\"}, ",", "\\")
	}
}

func BenchmarkJoinEscapingBoth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"a\\,", "b\\,", "c\\,"}, ",", "\\")
	}
}
