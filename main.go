package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi/pkg/v2/codegen/go/provider"
	"golang.org/x/tools/go/packages"
)

func main() {
	name, rootPackagePath, pattern, outdir := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedTypes | packages.NeedCompiledGoFiles,
	}, pattern)
	if err != nil {
		log.Fatalf("failed to load %v: %v", pattern, err)
	}

	files, diags := provider.Generate(name, rootPackagePath, pkgs...)
	if len(diags) != 0 {
		writer, err := provider.NewDiagnosticWriter(os.Stderr, 0, true, pkgs...)
		if err != nil {
			log.Fatalf("internal error: failed to create diagnostic writer: %w", err)
		}
		writer.WriteDiagnostics(diags)
		if diags.HasErrors() {
			return
		}
	}

	for name, contents := range files {
		path := filepath.Join(outdir, name)
		if err = os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			log.Fatalf("failed to create output directory: %v", err)
		}
		if err = ioutil.WriteFile(path, contents, 0600); err != nil {
			log.Fatalf("failed to write file %v: %v", path, err)
		}
	}
}
