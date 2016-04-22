package supersplit // import "fknsrs.biz/p/supersplit"

import (
	"strings"
)

type state int

const (
	stateBase = iota
	stateEscape
)

func Escaped(s, sep, esc string) []string {
	var escaping bool

	var (
		bits []string
		bit  string
	)

	for _, c := range s {
		switch escaping {
		case true:
			switch string(c) {
			case esc:
				bit += esc
			case sep:
				bit += sep
			default:
				bit += esc + string(c)
			}

			escaping = false
		case false:
			switch string(c) {
			case esc:
				escaping = true
			case sep:
				bits = append(bits, bit)
				bit = ""
			default:
				bit += string(c)
			}
		}
	}

	if escaping {
		bit += esc
	}

	if len(bit) > 0 {
		bits = append(bits, bit)
	}

	return bits
}

func Join(a []string, sep, esc string) string {
	r := make([]string, len(a))

	for i, v := range a {
		r[i] = strings.Replace(strings.Replace(v, esc, esc+esc, -1), sep, esc+sep, -1)
	}

	return strings.Join(r, sep)
}
