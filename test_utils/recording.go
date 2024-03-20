package test_utils

import (
	"io"
	"os"
	"strings"
)

type Recording struct {
	reader         *os.File
	writer         *os.File
	originalStdout *os.File
}

func StartRecording() Recording {
	// Save the real stdout for later
	originalStdout := os.Stdout

	// Create a new os pipe
	reader, writer, _ := os.Pipe()

	// Set os stdout to be the writer of this pipe
	os.Stdout = writer

	return Recording{reader, writer, originalStdout}
}

func EndRecording(rec Recording) string {
	// Close the pipe
	rec.writer.Close()

	// Read from the reader
	out, _ := io.ReadAll(rec.reader)

	// Convert to string and trim whitespace
	result := string(out)
	result = strings.TrimSpace(result)

	// Restore Stdout
	os.Stdout = rec.originalStdout

	return result
}
