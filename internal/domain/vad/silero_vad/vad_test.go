package silero_vad

import (
	"runtime"
	"testing"
)

func TestResolveNumSessionsDefaultsToCPUCount(t *testing.T) {
	got := resolveNumSessions(withDefaults(map[string]interface{}{}))
	want := runtime.NumCPU()
	if got != want {
		t.Fatalf("resolveNumSessions() = %d, want %d", got, want)
	}
}

func TestResolveNumSessionsUsesPoolSizeFallback(t *testing.T) {
	got := resolveNumSessions(withDefaults(map[string]interface{}{
		"pool_size": 4,
	}))
	if got != 4 {
		t.Fatalf("resolveNumSessions() = %d, want 4", got)
	}
}

func TestResolveNumSessionsPrefersNumSessions(t *testing.T) {
	got := resolveNumSessions(withDefaults(map[string]interface{}{
		"pool_size":    4,
		"num_sessions": 2,
	}))
	if got != 2 {
		t.Fatalf("resolveNumSessions() = %d, want 2", got)
	}
}

func TestResolveNumSessionsFallsBackToCPUCountForInvalidValue(t *testing.T) {
	got := resolveNumSessions(withDefaults(map[string]interface{}{
		"pool_size": 0,
	}))
	want := runtime.NumCPU()
	if got != want {
		t.Fatalf("resolveNumSessions() = %d, want %d", got, want)
	}
}
