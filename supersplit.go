package supersplit // import "fknsrs.biz/p/supersplit"

import (
	"strings"
)

func Escaped(s, sep, esc string) []string {
	var r []string

	o := 0
	var bits []string
	for {
		if p := strings.Index(s[o:], esc); p != -1 {
			if p != len(s[o:])-1 {
				bits = append(bits, s[o:o+p], s[o+p+1:o+p+2])

				o += p + 2

				continue
			}
		}

		if p := strings.Index(s[o:], sep); p != -1 {
			r = append(r, strings.Join(append(bits, s[o:o+p]), ""))

			bits = nil

			o += p + 1

			continue
		}

		break
	}

	if o < len(s) {
		bits = append(bits, s[o:])
	}

	if len(bits) > 0 {
		r = append(r, strings.Join(bits, ""))
	}

	return r
}
