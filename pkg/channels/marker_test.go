// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT
//
// Copyright (c) 2026 PicoClaw contributors

package channels

import (
	"testing"
)

func TestSplitByMarker_Basic(t *testing.T) {
	content := "Hello <|[SPLIT]|>World"
	chunks := SplitByMarker(content)

	if len(chunks) != 2 {
		t.Fatalf("Expected 2 chunks, got %d: %q", len(chunks), chunks)
	}
	if chunks[0] != "Hello" {
		t.Errorf("Expected first chunk 'Hello', got %q", chunks[0])
	}
	if chunks[1] != "World" {
		t.Errorf("Expected second chunk 'World', got %q", chunks[1])
	}
}

func TestSplitByMarker_NoMarker(t *testing.T) {
	content := "Hello World"
	chunks := SplitByMarker(content)

	if len(chunks) != 1 {
		t.Fatalf("Expected 1 chunk, got %d: %q", len(chunks), chunks)
	}
	if chunks[0] != "Hello World" {
		t.Errorf("Expected chunk 'Hello World', got %q", chunks[0])
	}
}

func TestSplitByMarker_MultipleMarkers(t *testing.T) {
	content := "Part1 <|[SPLIT]|> Part2 <|[SPLIT]|> Part3"
	chunks := SplitByMarker(content)

	if len(chunks) != 3 {
		t.Fatalf("Expected 3 chunks, got %d: %q", len(chunks), chunks)
	}
	if chunks[0] != "Part1" || chunks[1] != "Part2" || chunks[2] != "Part3" {
		t.Errorf("Unexpected chunks: %q", chunks)
	}
}

func TestSplitByMarker_EmptyParts(t *testing.T) {
	// Test consecutive markers and leading/trailing markers
	content := "<|[SPLIT]|>Hello <|[SPLIT]|><|[SPLIT]|>World<|[SPLIT]|>"
	chunks := SplitByMarker(content)

	if len(chunks) != 2 {
		t.Fatalf("Expected 2 chunks, got %d: %q", len(chunks), chunks)
	}
	if chunks[0] != "Hello" || chunks[1] != "World" {
		t.Errorf("Unexpected chunks: %q", chunks)
	}
}

func TestSplitByMarker_WhitespaceTrimmed(t *testing.T) {
	content := "  Hello   <|[SPLIT]|>   World  "
	chunks := SplitByMarker(content)

	if len(chunks) != 2 {
		t.Fatalf("Expected 2 chunks, got %d: %q", len(chunks), chunks)
	}
	if chunks[0] != "Hello" || chunks[1] != "World" {
		t.Errorf("Whitespace should be trimmed: %q", chunks)
	}
}

func TestSplitByMarker_EmptyInput(t *testing.T) {
	chunks := SplitByMarker("")
	if len(chunks) != 0 {
		t.Errorf("Expected empty slice for empty input, got %d chunks", len(chunks))
	}
}
