package user

import (
	"context"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	const in = "258"
	resp, err := Get(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	if got := resp.Label.Value; !strings.Contains(got, in) {
		t.Errorf("Get(%q) = %q, expected to contain %q", in, got, in)
	}
}
