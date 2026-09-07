package cmd

import (
	"os"
	"runtime/debug"
)

// defaultGCPercent is the garbage collector target used when GOGC is not set.
// Formatting allocates a great deal of short-lived memory and keeps little of
// it, so with the runtime default of 100 the collector runs often and takes
// about half of the process's CPU time. At 200 the CPU spent formatting the
// afmt benchmark corpus drops by about a third; peak memory rises only when
// many files are formatted at once.
const defaultGCPercent = 200

// gcPercent returns the garbage collector target to apply, or -1 to leave the
// runtime's setting alone because GOGC is set in the environment.
func gcPercent(gogc string) int {
	if gogc != "" {
		return -1
	}
	return defaultGCPercent
}

func configureGC() {
	if percent := gcPercent(os.Getenv("GOGC")); percent > 0 {
		debug.SetGCPercent(percent)
	}
}
