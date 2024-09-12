package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	Func "ascii-art/Functioin"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name     string
		banner   string
		text     string
		expected string
	}{
		// TODO: Add test cases.
		{
			name:     "sample text",
			banner:   "standard",
			text:     "hello",
			expected: "test-files/test1.txt",
		},
		{
			name:     "sample text",
			banner:   "shadow",
			text:     "hello world",
			expected: "test-files/test2.txt",
		},
		{
			name:     "sample text",
			banner:   "thinkertoy",
			text:     "nice 2 meet you",
			expected: "test-files/test3.txt",
		},
		{
			name:     "sample text",
			banner:   "thinkertoy",
			text:     "\"#$%&/()*+,-./",
			expected: "test-files/test4.txt",
		},
		{
			name:     "sample text",
			banner:   "standard",
			text:     "Hello\nThere",
			expected: "test-files/test5.txt",
		},
		{
			name:     "sample text",
			banner:   "standard",
			text:     "\n\n\n",
			expected: "test-files/test6.txt",
		},
	}

	for _, v := range tests {
		textsep := strings.Split(v.text, "\n")
		if strings.ReplaceAll(v.text, "\n", "") == "" && len(v.text) > 1 {
			textsep = textsep[1:]
		}
		contentfile := Func.Readfile(v.banner)
		matrix := Func.Split((string(contentfile)), v.banner)
		t.Run(v.name, func(tseting *testing.T) {
			read, _ := os.ReadFile(v.expected)
			resultefin := string(read)

			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			Func.Print(textsep, matrix)
			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)
			result := buf.String()
			if result != resultefin {
				t.Errorf("Print() output =\n%v\nWant:\n%v", result, resultefin)
			}
		})
	}
}
