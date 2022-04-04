package main

import (
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestGenerate_filter(t *testing.T) {
	g := New()
	file := `
	package foo
	
	type Foo struct {
		Bar string
	}
	
	type FooInterface interface {}
	`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", file, 0)
	NoError(t, err)

	specs := g.filter(f, map[string]struct{}{
		"Foo": {},
		"Bar": {},
	})

	if len(specs) != 1 {
		t.Errorf("expected 1 struct returned after filter, but got %d", len(specs))
	}
}

func TestGenerate_newRecord(t *testing.T) {
	g := New()
	f, err := g.parse("testdata/simple.go")
	NoError(t, err)
	specs := g.filter(f, map[string]struct{}{
		"Simple": {},
	})

	if len(specs) != 1 {
		t.Fatalf("expected 1 struct returned after filter, but got %d", len(specs))
	}

	r := newRecord(specs[0])
	Equal(t, "record name", "Simple", r.name)
	Equal(t, "record field count", 2, len(r.fields))
	Equal(t, "record field[0] name", "Bar", r.fields[0].name)
	Equal(t, "record field[0] type", "string", r.fields[0].t)
}

func Equal(t *testing.T, name string, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %s to be equal %v, but got %v", name, expected, actual)
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected no error, but got: %s", err)
	}
}
