package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4
	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

// TestCountLines test the count function set to count lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3")

	exp := 3
	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

// TestCountBytes test the count function to count bytes
func TestCountBytes(t *testing.T) {
	// test_byte := []byte("falcon")
	b := bytes.NewBuffer([]byte("žincica mňam"))

	// b := []byte("falcon")

	exp := 14
	res := count(b, false, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}

}
