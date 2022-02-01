package codegen

type GoFileContext interface {
	RequireImport(pkgPath string) *GoFileImport
}

type GoFileImport struct {
	PkgPath string
	Alias   string
}
