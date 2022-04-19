package main

import (
	"os"
	"path"

	log "github.com/ipfs/go-log"
	"github.com/ipld/edelweiss/compile"
	"github.com/ipld/edelweiss/defs"
)

var proto = defs.Defs{

	// hello service definition
	defs.Named{
		Name: "GreetingService",
		Type: defs.Service{
			Methods: defs.Methods{
				defs.Method{
					Name: "Hello",
					Type: defs.Fn{
						Arg:    defs.Ref{Name: "HelloRequest"},
						Return: defs.Ref{Name: "HelloResponse"},
					},
				},
			},
		},
	},

	// request type
	defs.Named{
		Name: "HelloRequest",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{
					Name: "Name",
					Type: defs.String{},
				},
				defs.Field{
					Name: "Address",
					Type: defs.Ref{Name: "Address"},
				},
			},
		},
	},

	defs.Named{
		Name: "Address",
		Type: defs.Inductive{
			Cases: defs.Cases{
				defs.Case{Name: "US", GoName: "US", Type: defs.Ref{Name: "USAddress"}},
				defs.Case{Name: "SouthKorea", GoName: "SK", Type: defs.Ref{Name: "SKAddress"}},
			},
			Default: defs.DefaultCase{
				GoKeyName:   "OtherCountry",
				GoValueName: "OtherAddress",
				Type: defs.Named{ // type AddressLines is defined and named inline
					Name: "AddressLines",
					Type: defs.List{Element: defs.String{}},
				},
			},
		},
	},

	defs.Named{
		Name: "USAddress",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{Name: "street", GoName: "Street", Type: defs.String{}},
				defs.Field{Name: "city", GoName: "City", Type: defs.String{}},
				defs.Field{Name: "state", GoName: "State", Type: defs.Ref{Name: "State"}},
				defs.Field{Name: "zip", GoName: "ZIP", Type: defs.Int{}},
			},
		},
	},

	defs.Named{
		Name: "State",
		Type: defs.Union{
			Cases: defs.Cases{
				defs.Case{
					Name: "ca", GoName: "CA",
					Type: defs.Named{ // inline definition and naming of type StateCA
						Name: "StateCA",
						Type: defs.SingletonString{String: "CA"},
					},
				},
				defs.Case{
					Name: "ny", GoName: "NY",
					Type: defs.Named{ // inline definition and naming of type StateNY
						Name: "StateNY",
						Type: defs.SingletonString{String: "NY"},
					},
				},
				defs.Case{
					Name: "other", GoName: "Other",
					Type: defs.String{},
				},
			},
		},
	},

	defs.Named{
		Name: "SKAddress",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{Name: "street", GoName: "Street", Type: defs.String{}},
				defs.Field{Name: "city", GoName: "City", Type: defs.String{}},
				defs.Field{Name: "province", GoName: "Province", Type: defs.String{}},
				defs.Field{Name: "postal_code", GoName: "PostalCode", Type: defs.Int{}},
			},
		},
	},

	// response type
	defs.Named{
		Name: "HelloResponse",
		Type: defs.Union{
			Cases: defs.Cases{
				defs.Case{Name: "english", GoName: "English", Type: defs.String{}},
				defs.Case{Name: "korean", GoName: "Korean", Type: defs.String{}},
			},
		},
	},
}

var logger = log.Logger("api generator")

func main() {
	wd, err := os.Getwd()
	if err != nil {
		logger.Errorf("working dir (%v)\n", err)
		os.Exit(-1)
	}
	dir := path.Join(wd, "proto")
	x := &compile.GoPkgCodegen{
		GoPkgDirPath: dir,
		GoPkgPath:    "github.com/ipld/edelweiss/examples/greeting-service/api/proto",
		Defs:         proto,
	}
	goFile, err := x.Compile()
	if err != nil {
		logger.Errorf("compilation (%v)\n", err)
		os.Exit(-1)
	}
	if err = os.Mkdir(dir, 0755); err != nil {
		logger.Errorf("making pkg dir (%v)\n", err)
		os.Exit(-1)
	}
	if err = goFile.Build(); err != nil {
		logger.Errorf("build (%v)\n", err)
		os.Exit(-1)
	}
}
