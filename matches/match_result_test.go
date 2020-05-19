package matches

import (
	"testing"
	"time"
)

func TestNewMatchResult(t *testing.T) {
	matchResultMatch := NewMatchResult().SetMatch(true)
	matchResultNotMatch := NewMatchResult().SetMatch(false)

	if matchResultMatch != nil && matchResultMatch.IsMatched {
		t.Logf("NewMatchResult(\"true\") PASSED got %v.",
			matchResultMatch)
	} else {
		t.Errorf("NewMatchResult(\"true\") FAILED got %v.",
			matchResultMatch)
	}

	if matchResultNotMatch != nil && !matchResultNotMatch.IsMatched {
		t.Logf("NewMatchResult(\"false\") PASSED got %v.",
			matchResultNotMatch)
	} else {
		t.Errorf("NewMatchResult(\"false\") FAILED got %v.",
			matchResultNotMatch)
	}
}

func TestTime(t *testing.T) {
	matchResult := NewMatchResult()
	startTime := time.Now()

	time.Sleep(1 * time.Second)

	matchResult.SetMatch(true)

	if matchResult.Elapsed >= 1000 {
		t.Logf("Time(\"%v\") PASSED got %v.",
			startTime, matchResult.Elapsed)
	} else {
		t.Errorf("Time(\"%v\") FAILED got %v.",
			startTime, matchResult.Elapsed)
	}
}

func TestSetMatch(t *testing.T) {
	matchResultNotMatch := NewMatchResult().SetMatch(false)

	if matchResultNotMatch != nil && !matchResultNotMatch.IsMatched {
		t.Logf("NewMatchResult(\"false\") PASSED got %v.",
			matchResultNotMatch)
	} else {
		t.Errorf("NewMatchResult(\"false\") FAILED got %v.",
			matchResultNotMatch)
	}

	matchResultMatch := matchResultNotMatch.SetMatch(true)

	if matchResultMatch != nil && matchResultMatch.IsMatched {
		t.Logf("SetMatch(\"true\") PASSED got %v.",
			matchResultMatch)
	} else {
		t.Errorf("SetMatch(\"true\") FAILED got %v.",
			matchResultMatch)
	}
}
