package logserver

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestLogToFile(t *testing.T) {
	logToFile("test log")

	file, _ := os.Open(logFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lastLine string
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	if !strings.Contains(lastLine, "test log") {
		t.Errorf("Log file does not contain the expected log")
	}
}
