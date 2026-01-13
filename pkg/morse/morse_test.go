package morse

import "testing"

func TestToMorseToText_RoundTrip_Cyrillic(t *testing.T) {
	in := "Привет мир"
	m := ToMorse(in)
	if m == "" {
		t.Fatalf("expected non-empty morse output")
	}

	out := ToText(m)
	if out == "" {
		t.Fatalf("expected non-empty text output")
	}

	// DefaultConverter has WithLowercaseHandling(true) (uppercasing),
	// so we expect an uppercase result for Cyrillic.
	want := "ПРИВЕТ МИР"
	if out != want {
		t.Fatalf("round-trip mismatch: want=%q got=%q morse=%q", want, out, m)
	}
}

func TestToMorseToText_RoundTrip_Digits(t *testing.T) {
	in := "12345 0"
	m := ToMorse(in)
	out := ToText(m)

	// digits should be preserved
	if out != in {
		t.Fatalf("digits round-trip mismatch: want=%q got=%q morse=%q", in, out, m)
	}
}

func TestRuneToMorseAndBack_KnownRune(t *testing.T) {
	r := 'А'
	code := RuneToMorse(r)
	if code == "" {
		t.Fatalf("expected non-empty morse code for %q", r)
	}

	back := MorseToRune(code)
	if back != r {
		t.Fatalf("expected %q back, got %q (code=%q)", r, back, code)
	}
}

func TestDefaultMorse_UniqueValues(t *testing.T) {
	seen := make(map[string]rune, len(DefaultMorse))

	for r, code := range DefaultMorse {
		if prev, ok := seen[code]; ok {
			if (prev == 'Ъ' && r == 'Ь') || (prev == 'Ь' && r == 'Ъ') {
				continue
			}
			t.Fatalf("duplicate morse code %q for runes %q and %q", code, prev, r)
		}
		seen[code] = r
	}
}

func TestNewConverter_NilMapPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for nil EncodingMap")
		}
	}()

	_ = NewConverter(nil)
}
