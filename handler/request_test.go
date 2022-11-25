package handler

import (
	"testing"
)

func TestUnserializeURL(t *testing.T) {
	if want, have := "/", unserializeURL("").String(); want != have {
		t.Fatalf("unexpected unserialized URL, want %q, have %q", want, have)
	}

	if want, have := "https://mytest.org/path?key=v1", unserializeURL("https://mytest.org/path?key=v1").String(); want != have {
		t.Fatalf("unexpected unserialized URL, want %q, have %q", want, have)
	}
}
