package config

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	e := FromMap(map[string]string{"A": "604800", "B": "7d", "C": "36h", "D": "soon", "E": ""})
	for k, want := range map[string]time.Duration{"A": 7 * 24 * time.Hour, "B": 7 * 24 * time.Hour, "C": 36 * time.Hour, "E": 0, "F": 0} {
		if got := e.Duration(k); got != want {
			t.Errorf("%s: %v, want %v", k, got, want)
		}
	}
	if e.Duration("D") != 0 || len(e.Problems) != 1 {
		t.Fatalf("a bad value: %v", e.Problems)
	}
}

func TestFloat(t *testing.T) {
	e := FromMap(map[string]string{"A": "25", "B": "29.97", "C": "-1", "D": "NaN", "E": "fast"})
	if a, b, f := e.Float("A", 0), e.Float("B", 0), e.Float("F", 7); a != 25 || b != 29.97 || f != 7 || len(e.Problems) != 0 {
		t.Fatalf("%v %v %v %v", a, b, f, e.Problems)
	}
	for _, k := range []string{"C", "D", "E"} {
		if got := e.Float(k, 0); got != 0 {
			t.Errorf("%s: %v", k, got)
		}
	}
	if len(e.Problems) != 3 {
		t.Fatalf("three bad values: %v", e.Problems)
	}
}

func TestLocation(t *testing.T) {
	e := FromMap(map[string]string{"A": "America/Sao_Paulo", "B": "Mars/Olympus", "C": ""})
	if l := e.Location("A"); l == nil || l.String() != "America/Sao_Paulo" {
		t.Fatalf("A: %v", l)
	}
	if e.Location("C") != nil || e.Location("D") != nil || len(e.Problems) != 0 {
		t.Fatalf("unset: %v", e.Problems)
	}
	if e.Location("B") != nil || len(e.Problems) != 1 {
		t.Fatalf("a zone that does not exist: %v", e.Problems)
	}
}
