package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Generate struct {
	file *ast.File
}

func New(filename string) (*Generate, error) {
	g := Generate{}
	f, err := g.parse(filename)
	if err != nil {
		return nil, err
	}

	g.file = f
	return &g, nil
}

func (g *Generate) Convert(strcts map[string]struct{}) error {
	specs := g.filterStructs(strcts)
	if len(specs) == 0 {
		return nil
	}

	return nil
}

func (g *Generate) convertStruct(spec *ast.TypeSpec) (string, error) {
	s := "interface " + spec.Name.Name + " {\n"
	for _, f := range spec.TypeParams.List {
		s += f.Names
	}

	s += "}"
	spec.TypeParams
}

func (g *Generate) parse(filename string) (*ast.File, error) {
	tfs := token.NewFileSet()
	return parser.ParseFile(tfs, filename, nil, 0)
}

func (g *Generate) filterStructs(strcts map[string]struct{}) []*ast.TypeSpec {
	tss := make([]*ast.TypeSpec, 0)
	for _, dec := range g.file.Decls {
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
