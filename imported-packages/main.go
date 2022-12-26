package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

/*
	特定ディレクトリ配下の go ファイルが import しているパッケージの一覧を表示する。
*/

var pkgdir string

func main() {
	flag.StringVar(&pkgdir, "d", "./sample", "directory to be parsed")
	flag.Parse()

	fset := token.NewFileSet()
	f, err := parser.ParseDir(fset, pkgdir, nil, parser.ImportsOnly)
	if err != nil {
		log.Fatal(err)
	}

	for p, a := range f {
		fmt.Printf("package: %s\n\n", p)
		for fn, af := range a.Files {
			fmt.Printf("%s imports\n", fn)
			for _, imp := range af.Imports {
				fmt.Printf("\t%s\n", imp.Path.Value)
			}
		}
	}
}
