package logger

import (
	"fmt"
	"io"

	"github.com/spf13/viper"
)

// Logger implements a logging system
type Logger struct {
	Writer io.Writer
}

func (l *Logger) shouldPrint(level int) bool {
	v := viper.GetInt("verbose")
	return v >= level
}

// Print formats using the default formats for its operands and writes to Logger.Writer.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (l *Logger) Print(level int, text string) (int, error) {
	if l.shouldPrint(level) == false {
		return 0, nil
	}

	return fmt.Fprint(l.Writer, text)
}

// Printf formats according to a format specifier and writes to Logger.Writer.
// It returns the number of bytes written and any write error encountered.
func (l *Logger) Printf(level int, text string, a ...interface{}) (int, error) {
	if l.shouldPrint(level) == false {
		return 0, nil
	}

	return fmt.Fprintf(l.Writer, text, a...)
}

// Println formats using the default formats for its operands and writes to Logger.Writer.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (l *Logger) Println(level int, text string) (int, error) {
	if l.shouldPrint(level) == false {
		return 0, nil
	}

	return fmt.Fprintln(l.Writer, text)
}
