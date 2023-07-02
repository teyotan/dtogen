package main

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/dave/jennifer/jen"
)

type generator struct{}

func (g *generator) generate(file *ast.File) *jen.File {
	f := jen.NewFile("dto")
	packageMap := map[string]string{}
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ImportSpec:
			pth := x.Path.Value[1 : len(x.Path.Value)-1]
			name := pth[strings.LastIndex(pth, "/")+1:]
			if x.Name != nil {
				name = x.Name.Name
			}
			packageMap[name] = pth
		case *ast.TypeSpec:
			fields := []field{}
			if str, ok := x.Type.(*ast.StructType); ok {
				for _, l := range str.Fields.List {
					for _, fieldName := range l.Names {
						field := processAstExp(fieldName.Name, l.Type)
						fields = append(fields, field)
					}
				}

				g.generateGetRequest(f, x.Name.Name, fields, packageMap)
			}

		}
		return true
	})

	return f
}

func (g *generator) generateGetRequest(f *jen.File, structName string, fields []field, packageMap map[string]string) {
	f.Type().Id(fmt.Sprintf("Get%sListRequest", structName)).Struct(
		jen.Id("Limit").Uint(),
		jen.Id("Offset").Uint(),
	)

	jenFields := []jen.Code{}
	for _, f := range fields {
		jenFields = append(jenFields, f.toJenCode(packageMap))
	}
	f.Type().Id(fmt.Sprintf("Get%sListResponse", structName)).Struct(
		jen.Id("Data").Index().Struct(
			jenFields...,
		),
	)
}
