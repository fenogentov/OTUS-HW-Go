package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var tmp strings.Builder
	runeS := []rune(s)

	re := regexp.MustCompile(`^[0-9]|[\t\v\f\r]| |[@/_]|\\[A-Za-z]`)
	if re.MatchString(s) {
		return "", ErrInvalidString
	}

	if len(runeS) == 0 { // test ""
		return "", nil
	}

	i := 0
	for i < (len(runeS)) {
		if string(runeS[i]) == `\` && len(runeS) > i+1 {
			if unicode.IsDigit(runeS[i+1]) || runeS[i+1] == '\\' { // \45 \\5
				if len(runeS) > i+2 && unicode.IsDigit(runeS[i+2]) {
					d, _ := strconv.Atoi(string(runeS[i+2]))
					r := strings.Repeat(string(runeS[i+1]), d)
					tmp.WriteString(r)
					i += 3
					continue
				}
				tmp.WriteString(string(runeS[i+1]))
				i += 2
				continue
			}
			return "", ErrInvalidString
		}

		if len(runeS) > i+2 && unicode.IsDigit(runeS[i+1]) && unicode.IsDigit(runeS[i+2]) { // число, а не цифра
			return "", ErrInvalidString
		} else if len(runeS) > i+1 && unicode.IsDigit(runeS[i+1]) { // цифра после символа
			d, _ := strconv.Atoi(string(runeS[i+1]))
			r := strings.Repeat(string(runeS[i]), d)
			tmp.WriteString(r)
			i += 2
			continue
		}

		tmp.WriteString(string(runeS[i]))
		i++
	}

	out := tmp.String()
	return out, nil
}
