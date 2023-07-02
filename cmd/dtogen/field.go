package main

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/dave/jennifer/jen"
)

type field struct {
	Name      string
	Package   string
	Type      string
	IsPointer bool
}

func processAstExp(name string, exp ast.Expr) field {
	f := process(exp)
	f.Name = name

	return f
}

func process(exp ast.Expr) field {
	switch v := exp.(type) {
	case *ast.ArrayType:
		fmt.Println("ArrayType")
	case *ast.StructType:
		fmt.Println("StructType")
	case *ast.FuncType:
		fmt.Println("FuncType")
	case *ast.InterfaceType:
		fmt.Println("InterfaceType")
	case *ast.MapType:
		fmt.Println("MapType")
	case *ast.ChanType:
		fmt.Println("ChanType")
	case *ast.ParenExpr:
		fmt.Println("ParenExpr")
	case *ast.SelectorExpr:
		f := process(v.X)
		return field{
			Package: f.Type,
			Type:    v.Sel.Name,
		}
	case *ast.IndexExpr:
		fmt.Println("IndexExpr")
	case *ast.IndexListExpr:
		fmt.Println("IndexListExpr")
	case *ast.SliceExpr:
		fmt.Println("SliceExpr")
	case *ast.BadExpr:
		fmt.Println("BadExpr")
	case *ast.Ident:
		return field{
			Type: v.Name,
		}
	case *ast.Ellipsis:
		fmt.Println("Ellipsis")
	case *ast.BasicLit:
		fmt.Println("BasicLit")
	case *ast.FuncLit:
		fmt.Println("FuncLit")
	case *ast.CompositeLit:
		fmt.Println("CompositeLit")
	case *ast.TypeAssertExpr:
		fmt.Println("TypeAssertExpr")
	case *ast.CallExpr:
		fmt.Println("CallExpr")
	case *ast.StarExpr:
		f := process(v.X)
		return field{
			Package:   f.Package,
			Type:      f.Type,
			IsPointer: true,
		}
	case *ast.UnaryExpr:
		fmt.Println("UnaryExpr")
	case *ast.BinaryExpr:
		fmt.Println("BinaryExpr")
	case *ast.KeyValueExpr:
		fmt.Println("KeyValueExpr")
	}
	return field{}
}

func (f field) toJenCode(packageMap map[string]string) jen.Code {
	jenField := jen.Id(f.Name)
	if strings.HasPrefix(f.Type, "*") {
		f.Type = f.Type[1:]
		jenField = jenField.Op("*")
	}

	jenField = jenField.Qual(packageMap[f.Package], f.Type)

	return jenField
}
