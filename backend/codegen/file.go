package codegen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type GoFile struct {
	FilePath string
	PkgPath  string
	Types    GoTypeImpls
}

func (f *GoFile) PkgName() string {
	return path.Base(f.PkgPath)
}

func (f *GoFile) Build() error {
	if err := os.MkdirAll(path.Dir(f.FilePath), 0755); err != nil {
		return err
	}
	body, err := f.Generate()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f.FilePath, body, 0644)
}

func (f *GoFile) Generate() ([]byte, error) {
	// generate types
	var typeDefBuf bytes.Buffer
	fctx := &goFileContext{}
	for _, t := range f.Types {
		if err := t.WriteDef(fctx, &typeDefBuf); err != nil {
			return nil, fmt.Errorf("generating type %#v (%w)", t, err)
		}
	}
	// generate go file header
	var headerBuf bytes.Buffer
	fmt.Fprintf(&headerBuf, "package %s\n\n", f.PkgName())
	if len(fctx.imported) > 0 {
		fmt.Fprintf(&headerBuf, "import(\n")
		for _, imp := range fctx.imported {
			fmt.Fprintf(&headerBuf, "\t%s %q\n", imp.Alias, imp.PkgPath)
		}
		fmt.Fprintf(&headerBuf, ")\n\n")
	}
	return append(headerBuf.Bytes(), typeDefBuf.Bytes()...), nil
}

type goFileContext struct {
	imported []*GoFileImport
}

func (fctx *goFileContext) RequireImport(pkgPath string) *GoFileImport {
	if imp := fctx.lookup(pkgPath); imp != nil {
		return imp
	} else {
		imp = &GoFileImport{PkgPath: pkgPath, Alias: fmt.Sprintf("pd%d", len(fctx.imported)+1)}
		fctx.imported = append(fctx.imported, imp)
		return imp
	}
}

func (fctx *goFileContext) lookup(pkgPath string) *GoFileImport {
	for _, imp := range fctx.imported {
		if imp.PkgPath == pkgPath {
			return imp
		}
	}
	return nil
}
