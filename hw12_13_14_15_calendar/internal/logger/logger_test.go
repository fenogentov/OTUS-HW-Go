package logger

import (
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	fmt.Println("point 1")
	logg := New("tst/tst.log", "DEBUG")
	logg.Debug("Debug logging")
	logg.Info("Info logging")
	logg.Warn("Warn logging")
	logg.Error("Error logging")
}
