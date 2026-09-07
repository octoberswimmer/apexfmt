package cmd

import "testing"

func Test_gc_percent_defaults_to_200_when_GOGC_is_unset(t *testing.T) {
	if got := gcPercent(""); got != 200 {
		t.Errorf("got %d, want 200", got)
	}
}

func Test_gc_percent_leaves_the_runtime_setting_when_GOGC_is_set(t *testing.T) {
	for _, gogc := range []string{"100", "off", "400"} {
		if got := gcPercent(gogc); got != -1 {
			t.Errorf("GOGC=%s: got %d, want -1", gogc, got)
		}
	}
}
