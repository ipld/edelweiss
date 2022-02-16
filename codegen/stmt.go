package codegen

import (
	"fmt"
	"io"
)

type BoolLiteral bool

func (x BoolLiteral) Write(ctx GoFileContext, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%v", x)
	return err
}

type IntLiteral int

func (x IntLiteral) Write(ctx GoFileContext, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%d", x)
	return err
}

type FloatLiteral float64

func (x FloatLiteral) Write(ctx GoFileContext, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%v", x)
	return err
}

type StringLiteral string

func (x StringLiteral) Write(ctx GoFileContext, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%q", x)
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
	Returns    Blueprints
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
	if _, err := fmt.Fprint(w, ") {\n"); err != nil {
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

type Block Blueprints

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

type StructDef struct {
	Name   string
	Fields []VarDef
}

func (x StructDef) Write(ctx GoFileContext, w io.Writer) error {
	if _, err := fmt.Fprintf(w, "type %s struct {\n", x.Name); err != nil {
		return err
	}
	b := make(Blueprints, len(x.Fields))
	for i := range x.Fields {
		b[i] = x.Fields[i]
	}
	if err := Block(b).Write(ctx, Indent(w)); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, "}"); err != nil {
		return err
	}
	return nil
}
