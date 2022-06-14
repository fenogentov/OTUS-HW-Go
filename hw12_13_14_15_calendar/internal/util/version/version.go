package version

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	release   = "0.0"
	buildDate = "2022/06/07"
	gitHash   = "1cdfgv"
)

func Print() {
	if err := json.NewEncoder(os.Stdout).Encode(struct {
		Release   string
		BuildDate string
		GitHash   string
	}{
		Release:   release,
		BuildDate: buildDate,
		GitHash:   gitHash,
	}); err != nil {
		fmt.Printf("error while decode version info: %v\n", err)
	}
}
