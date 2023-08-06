package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 12345
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if got.Port != wantPort {
		t.Errorf("want: %v, got: %v", wantPort, got.Port)
	}

	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want: %v, got: %v", wantEnv, got.Env)
	}
}
