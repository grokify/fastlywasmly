package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/grokify/fastlywasmly/tarutil"
)

const (
	wasmfileArchivePath = "package/bin/main.wasm"
	tomlArchivePath     = "package/fastly.toml"
	outfileDefault      = "package.tar.gz"
)

func main() {
	wasmfile := flag.String("w", "", "WASM filepath")
	// bindir := flag.String("b", "", "binary directory")
	tomlfile := flag.String("t", "", "binary directory")
	outfile := flag.String("o", "", "output .tar.gz file")
	flag.Parse()

	files := map[string]string{}

	if len(*wasmfile) > 0 {
		files[*wasmfile] = wasmfileArchivePath
	}
	if len(*tomlfile) > 0 {
		files[*tomlfile] = tomlArchivePath
	}
	if len(*outfile) == 0 {
		*outfile = outfileDefault
	}
	err := tarutil.CreateArchiveGzipFile(*outfile, files)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("WROTE [%s]\n", *outfile)
}
