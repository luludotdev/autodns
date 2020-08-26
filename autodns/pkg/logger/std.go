package logger

import (
	"os"
)

var (
	// Stdout is a default logger configured to write to os.Stdout
	Stdout = &Logger{Writer: os.Stdout}

	// Stderr is a default logger configured to write to os.Stderr
	Stderr = &Logger{Writer: os.Stderr}
)
