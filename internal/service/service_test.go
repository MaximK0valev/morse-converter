package service

import "testing"

func TestAutoConvert_EmptyInput(t *testing.T) {
	_, err := AutoConvert("")
	if err == nil {
		t.Fatalf("expected error for empty input, got nil")
	}

	_, err = AutoConvert("   \n\t  ")
	if err == nil {
		t.Fatalf("expected error for whitespace-only input, got nil")
	}
}

func TestAutoConvert_TextToMorse(t *testing.T) {
	// Not Morse because it contains letters.
	got, err := AutoConvert("Привет")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == "" {
		t.Fatalf("expected non-empty morse output")
	}

	// Output should contain only morse alphabet symbols and separators.
	for _, r := range got {
		switch r {
		case '.', '-', ' ', '/':
			// ok
		default:
			t.Fatalf("unexpected rune in morse output: %q (%U), output=%q", r, r, got)
		}
	}
}

func TestAutoConvert_MorseToText(t *testing.T) {
	// This is Morse-like input.
	got, err := AutoConvert("... --- ...")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == "" {
		t.Fatalf("expected non-empty text output")
	}

	// Should not contain morse-only characters (dot/dash) in normal case.
	// (If unknown sequences exist, your handler might drop/replace them; here it's a known one.)
	for _, r := range got {
		if r == '.' || r == '-' {
			t.Fatalf("unexpected morse character in text output: %q, output=%q", r, got)
		}
	}
}
