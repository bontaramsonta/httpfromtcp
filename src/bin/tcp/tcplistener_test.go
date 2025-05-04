package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

// collectLines is a helper to read all lines from the channel returned by getLinesChannel.
func collectLines(f io.ReadCloser) []string {
	ch := getLinesChannel(f)
	var lines []string
	for line := range ch {
		lines = append(lines, line)
	}
	return lines
}

func TestGetLinesChannelBasic(t *testing.T) {
	data := "hello\nworld\nfoo\n"
	f := io.NopCloser(strings.NewReader(data))
	got := collectLines(f)
	want := []string{"hello", "world", "foo"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetLinesChannelEmpty(t *testing.T) {
	data := ""
	f := io.NopCloser(strings.NewReader(data))
	got := collectLines(f)
	var want []string
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetLinesChannelNoTrailingNewline(t *testing.T) {
	// Current behavior: only lines terminated with '\n' are sent
	data := "line1\nline2"
	f := io.NopCloser(strings.NewReader(data))
	got := collectLines(f)
	want := []string{"line1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetLinesChannelSingleLineNoNewline(t *testing.T) {
	// No '\n' at all => no lines emitted
	data := "solo"
	f := io.NopCloser(strings.NewReader(data))
	got := collectLines(f)
	var want []string
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
