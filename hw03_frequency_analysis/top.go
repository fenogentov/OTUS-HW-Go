package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(txt string) []string {
	N := 10

	re := regexp.MustCompile("[\".,?!:`;]|- ")
	txt = re.ReplaceAllString(txt, " ")
	re = regexp.MustCompile(" +|[\"\n\t\v\f\r]+")
	txt = re.ReplaceAllString(txt, " ")

	s := strings.Split(txt, " ")
	t := make(map[string]int)
	for _, w := range s {
		w = strings.ToLower(w)
		t[w]++
	}

	words := make(map[int][]string)
	for word, c := range t {
		words[c] = append(words[c], word)
	}

	keys := []int{}
	for key, sl := range words {
		keys = append(keys, key)
		sort.Strings(sl)
	}
	sort.Ints(keys)

	var out []string
	k := 1
	for j := len(keys); j > 0; j-- {
		for _, p := range words[j] {
			out = append(out, p)
			if k == N {
				return out
			}
			k++
		}
	}

	return nil
}
