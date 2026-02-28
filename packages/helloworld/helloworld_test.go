package main

import (
	"io"
	"os"
	"testing"
)

func TestMainOutput(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	out, _ := io.ReadAll(r)
	got := string(out)
	want := "Hello World!\n"
	if got != want {
		t.Errorf("main() output = %q, want %q", got, want)
	}
}
