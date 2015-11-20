package supersplit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBareSplitChar(t *testing.T) {
	a := assert.New(t)

	bits := Escaped("a,b", ",", "\\")

	a.Equal([]string{"a", "b"}, bits)
}

func TestEscapedSplitChar(t *testing.T) {
	a := assert.New(t)

	bits := Escaped("a\\,b", ",", "\\")

	a.Equal([]string{"a,b"}, bits)
}

func TestMixedSplitChar(t *testing.T) {
	a := assert.New(t)

	bits := Escaped("a\\,b,c", ",", "\\")

	a.Equal([]string{"a,b", "c"}, bits)
}

func TestEndingSplitChar(t *testing.T) {
	a := assert.New(t)

	bits := Escaped("a,", ",", "\\")

	a.Equal([]string{"a"}, bits)
}

func TestEndingEscapeChar(t *testing.T) {
	a := assert.New(t)

	bits := Escaped("a,\\", ",", "\\")

	a.Equal([]string{"a", "\\"}, bits)
}

func TestBeginningSplitChar(t *testing.T) {
	a := assert.New(t)

	bits := Escaped(",a,", ",", "\\")

	a.Equal([]string{"", "a"}, bits)
}

func TestBeginningEscapeChar(t *testing.T) {
	a := assert.New(t)

	bits := Escaped("\\,a,", ",", "\\")

	a.Equal([]string{",a"}, bits)
}

func TestAdjacentSplitChars(t *testing.T) {
	a := assert.New(t)

	bits := Escaped("a,,b", ",", "\\")

	a.Equal([]string{"a", "", "b"}, bits)
}

func TestAdjacentEscapeChars(t *testing.T) {
	a := assert.New(t)

	bits := Escaped("a\\\\,b", ",", "\\")

	a.Equal([]string{"a\\", "b"}, bits)
}

func TestJoin(t *testing.T) {
	a := assert.New(t)

	s := Join([]string{"a", "b", "c"}, ",", "\\")

	a.Equal("a,b,c", s)
}

func TestJoinEscapingSeparator(t *testing.T) {
	a := assert.New(t)

	s := Join([]string{"a,", "b,", "c,"}, ",", "\\")

	a.Equal("a\\,,b\\,,c\\,", s)
}

func TestJoinEscapingEscape(t *testing.T) {
	a := assert.New(t)

	s := Join([]string{"a\\", "b\\", "c\\"}, ",", "\\")

	a.Equal("a\\\\,b\\\\,c\\\\", s)
}

func TestJoinEscapingBoth(t *testing.T) {
	a := assert.New(t)

	s := Join([]string{"a\\,", "b\\,", "c\\,"}, ",", "\\")

	a.Equal("a\\\\\\,,b\\\\\\,,c\\\\\\,", s)
}
