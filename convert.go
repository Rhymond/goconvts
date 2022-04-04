package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type (
	Generate struct {
		records []record
	}

	field struct {
		name string
		t    string
	}

	record struct {
		name   string
		fields []field
	}
)

func New() *Generate {
	g := Generate{}
	return &g
}

func newRecord(spec *ast.TypeSpec) *record {
	r := new(record)
	r.name = spec.Name.Name
	r.fields = make([]field, 0)

	st := spec.Type.(*ast.StructType)
	for _, f := range st.Fields.List {
		switch f.Type.(type) {
		case *ast.Ident:
			stype := f.Type.(*ast.Ident).Name
			// tag = ""
			// if f.Tag != nil {
			// 	tag = f.Tag.Value
			// }
			name := f.Names[0].Name
			r.fields = append(r.fields, field{
				name: name,
				t:    stype,
			})
		}
	}

	return r
}

func (g *Generate) Convert(filename string, strcts map[string]struct{}) error {
	f, err := g.parse(filename)
	if err != nil {
		return err
	}

	specs := g.filter(f, strcts)
	if len(specs) == 0 {
		return nil
	}

	g.records = make([]record, len(specs))
	return nil
}

func (g *Generate) parse(filename string) (*ast.File, error) {
	tfs := token.NewFileSet()
	return parser.ParseFile(tfs, filename, nil, 0)
}

func (g *Generate) filter(file *ast.File, strcts map[string]struct{}) []*ast.TypeSpec {
	tss := make([]*ast.TypeSpec, 0)
	for _, dec := range file.Decls {
		gd, ok := dec.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}

		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			switch ts.Type.(type) {
			case *ast.StructType:
				if _, ok := strcts[ts.Name.Name]; ok {
					tss = append(tss, ts)
				}
			}
		}
	}

	return tss
}

// func (g *Generate) Convert(name string, v interface{}) string {
// 	el := reflect.ValueOf(v).Elem()
// 	s := "interface " + name + " {\n"
// 	for j := 0; j < el.NumField(); j++ {
// 		f := el.Field(j)
// 		n := el.Type().Field(j).Name
// 		t := f.Type().String()
// 		s += "    " + n + " " + t + ",\n"
// 	}
// 	s += "}"
//
// 	return s
// }
