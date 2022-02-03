package codegen

import (
	"fmt"
	"io"

	"github.com/ipld/edelweiss/util/indent"
)

type Blueprint interface {
	Write(GoFileContext, io.Writer) error
}

func Indent(w io.Writer) io.Writer {
	return indent.NewWriter(w, "\t")
}

// V stands for verbatim.
type V struct {
	String string
}

func (x V) Write(_ GoFileContext, w io.Writer) error {
	_, err := fmt.Fprint(w, x.String)
	return err
}

type VarDef struct {
	Var  Blueprint
	Type Blueprint
}

func (x VarDef) Write(ctx GoFileContext, w io.Writer) error {
	if err := x.Var.Write(ctx, w); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, " "); err != nil {
		return err
	}
	if err := x.Type.Write(ctx, w); err != nil {
		return err
	}
	return nil
}

type Return struct {
	Values []Blueprint
}

func (x Return) Write(ctx GoFileContext, w io.Writer) error {
	if _, err := fmt.Fprint(w, "return "); err != nil {
		return err
	}
	for _, v := range x.Values {
		if err := v.Write(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

type MethodDef struct {
	Receiver   VarDef
	MethodName string
	Args       []VarDef
	Returns    []Blueprint
	Body       Blueprint
}

func (x MethodDef) Write(ctx GoFileContext, w io.Writer) error {
	if _, err := fmt.Fprint(w, "func ("); err != nil {
		return err
	}
	if err := x.Receiver.Write(ctx, w); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, ") %s(", x.MethodName); err != nil {
		return err
	}
	for _, v := range x.Args {
		if err := v.Write(ctx, w); err != nil {
			return err
		}
		if _, err := fmt.Fprint(w, ", "); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprint(w, ") {\n", x.MethodName); err != nil {
		return err
	}
	if err := x.Body.Write(ctx, Indent(w)); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, "}"); err != nil {
		return err
	}
	return nil
}

type Block []Blueprint

func (x Block) Write(ctx GoFileContext, w io.Writer) error {
	for _, b := range x {
		if err := b.Write(ctx, w); err != nil {
			return err
		}
		if _, err := fmt.Fprintln(w, ""); err != nil {
			return err
		}
	}
	return nil
}
