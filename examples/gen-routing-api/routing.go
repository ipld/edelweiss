package main

import (
	"os"
	"path"

	log "github.com/ipfs/go-log"
	"github.com/ipld/edelweiss/compile"
	"github.com/ipld/edelweiss/defs"
)

var proto = defs.Defs{

	// delegated routing service definition
	defs.Named{
		Name: "DelegatedRouting",
		Type: defs.Service{
			Methods: defs.Methods{
				defs.Method{
					Name: "GetP2PProvide",
					Type: defs.Fn{
						Arg:    defs.Ref{Name: "GetP2PProvideRequest"},
						Return: defs.Ref{Name: "GetP2PProvideResponse"},
					},
				},
			},
		},
	},

	// request type
	defs.Named{
		Name: "GetP2PProvideRequest",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{
					Name: "Keys",
					Type: defs.List{Element: defs.Ref{Name: "Multihash"}},
				},
			},
		},
	},

	// response type
	defs.Named{
		Name: "GetP2PProvideResponse",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{
					Name: "ProvidersByKey",
					Type: defs.List{Element: defs.Ref{Name: "ProvidersByKey"}},
				},
			},
		},
	},
	defs.Named{
		Name: "ProvidersByKey",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{
					Name: "Key",
					Type: defs.Ref{Name: "Multihash"},
				},
				defs.Field{
					Name: "Provider",
					Type: defs.Ref{Name: "Provider"},
				},
			},
		},
	},

	// general routing types
	defs.Named{
		Name: "Multihash",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{
					Name: "Bytes",
					Type: defs.Bytes{},
				},
			},
		},
	},

	defs.Named{
		Name: "Provider",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{
					Name: "Nodes",
					Type: defs.List{Element: defs.Ref{Name: "Node"}},
				},
				defs.Field{
					Name: "Proto",
					Type: defs.List{Element: defs.Ref{Name: "TransferProto"}},
				},
			},
		},
	},

	defs.Named{
		Name: "Node",
		Type: defs.Inductive{
			Cases: defs.Cases{
				defs.Case{Name: "Peer", Type: defs.Ref{Name: "Peer"}},
			},
		},
	},

	defs.Named{
		Name: "Peer",
		Type: defs.Structure{
			Fields: defs.Fields{
				defs.Field{Name: "ID", Type: defs.Bytes{}},
				defs.Field{Name: "Multiaddresses", Type: defs.List{Element: defs.Bytes{}}},
			},
		},
	},

	defs.Named{
		Name: "TransferProto",
		Type: defs.Inductive{
			Cases: defs.Cases{
				defs.Case{Name: "Bitswap", Type: defs.Ref{Name: "BitswapTransfer"}},
			},
		},
	},

	defs.Named{
		Name: "BitswapTransfer",
		Type: defs.Structure{},
	},
}

var logger = log.Logger("proto generator")

func main() {
	wd, err := os.Getwd()
	if err != nil {
		logger.Errorf("working dir (%v)\n", err)
		os.Exit(-1)
	}
	dir := path.Join(wd, "proto")
	x := &compile.GoPkgCodegen{
		GoPkgDirPath: dir,
		GoPkgPath:    "github.com/ipld/edelweiss/examples/routing-api/proto",
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
