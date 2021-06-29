package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	"github.com/valyala/fastjson"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)

	data := bufio.NewScanner(r)
	for i := 0; data.Scan(); i++ {
		line := data.Bytes()

		email := fastjson.GetString(line, "Email")

		if !strings.Contains(email, domain) {
			continue
		}

		result[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
	}

	return result, nil
}
