package codegen

import (
	"bytes"
	"fmt"
	"io"
	"text/template"

	"github.com/ipld/edelweiss/util/indent"
)

type Blueprint interface {
	Write(GoFileContext, io.Writer) error
}

type BlueSlice []Blueprint

func (x BlueSlice) Write(ctx GoFileContext, w io.Writer) error {
	for _, b := range x {
		if err := b.Write(ctx, w); err != nil {
			return nil
		}
	}
	return nil
}

func Indent(w io.Writer) io.Writer {
	return indent.NewWriter(w, "\t")
}

// V stands for verbatim.
type V string

func (x V) Write(_ GoFileContext, w io.Writer) error {
	_, err := fmt.Fprint(w, x)
	return err
}

func Printf(f string, a ...interface{}) Blueprint {
	return V(fmt.Sprintf(f, a...))
}

// T stands for template.
type T struct {
	Src  string
	Data BlueMap
}

type BlueMap map[string]Blueprint

func (x BlueMap) Write(ctx GoFileContext, w io.Writer) error {
	for _, b := range x {
		if err := b.Write(ctx, w); err != nil {
			return nil
		}
	}
	return nil
}

func MergeBlueMaps(x, y BlueMap) BlueMap {
	xy := BlueMap{}
	for k, v := range x {
		xy[k] = v
	}
	for k, v := range y {
		xy[k] = v
	}
	return xy
}

var cachedTemplates = map[string]*template.Template{}

func compileTemplate(src string) *template.Template {
	if t, ok := cachedTemplates[src]; ok {
		return t
	} else {
		t = template.Must(template.New("").Parse(src))
		cachedTemplates[src] = t
		return t
	}
}

func (x T) Write(ctx GoFileContext, w io.Writer) error {
	data := map[string]interface{}{}
	for k, v := range x.Data {
		var buf bytes.Buffer
		if err := v.Write(ctx, &buf); err != nil {
			return err
		}
		data[k] = buf.String()
	}
	compileTemplate(x.Src).Execute(w, data)
	return nil
}
