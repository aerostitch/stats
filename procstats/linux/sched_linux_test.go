package linux

import (
	"os"
	"testing"
)

func TestGetProcSched(t *testing.T) {
	if sched, err := GetProcSched(os.Getpid()); err != nil {
		t.Error("GetProcSched:", err)
	}
}
