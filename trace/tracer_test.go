package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	tracer.Trace("Hello trace package.")
	if buf.String() != "Hello trace package.\n" {
		t.Errorf("Trace should not write '%s'.", buf.String())
	}
}
