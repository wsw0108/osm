package filter

import (
	"github.com/Knetic/govaluate"
	"github.com/paulmach/osm"
)

type Filter interface {
	Apply(e osm.Element) bool
}

type PassFilter struct{}

func NewPassFilter() *PassFilter {
	return &PassFilter{}
}

func (f *PassFilter) Apply(e osm.Element) bool {
	return true
}

type ExprFilter struct {
	raw  string
	expr *govaluate.EvaluableExpression
}

func NewExprFilter(raw string) (Filter, error) {
	f := &ExprFilter{raw: raw}
	if err := f.parse(); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *ExprFilter) parse() error {
	var err error
	f.expr, err = govaluate.NewEvaluableExpression(f.raw)
	return err
}

type parameters map[string]string

func (p parameters) Get(name string) (interface{}, error) {
	value, found := p[name]
	if !found {
		// support expr like: "Variable != nil" / "Variable == nil"
		return nil, nil
	}
	return value, nil
}

func (f *ExprFilter) Apply(e osm.Element) bool {
	tags := e.TagMap()
	params := parameters(tags)
	result, err := f.expr.Eval(params)
	if err != nil {
		return false
	}
	if b, ok := result.(bool); ok {
		return b
	}
	return false
}
