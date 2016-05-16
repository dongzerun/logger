package logger

import (
	. "common"
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewDefaultLogger()
	logger.Info(fmt.Sprintf("xxxxxxxxxxx"))
}
