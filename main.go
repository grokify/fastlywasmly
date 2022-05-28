package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/grokify/fastlywasmly/tarutil"
)

func main() {
	tomlfile := flag.String("t", "", "binary directory")
	wasmfile := flag.String("w", "", "WASM filepath")
	bindir := flag.String("b", "", "binary directory")

	flag.Parse()

	outfile, err := tarutil.BuildEdgePackage(*tomlfile, *wasmfile, *bindir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("WROTE [%s]\n", outfile)
}
