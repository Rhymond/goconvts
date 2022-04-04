package main

import (
	"testing"
)

func TestGenerate_filter(t *testing.T) {
	g, err := New("testdata/filtertest.go")
	NoError(t, err)
	specs := g.filterStructs(map[string]struct{}{
		"Foo": {},
		"Bar": {},
	})

	if len(specs) != 1 {
		t.Errorf("expected 1 struct returned after filter, but got %d", len(specs))
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected no error, but got: %s", err)
	}
}
