package filter

import (
	"github.com/paulmach/osm"
)

type Filter interface {
	Apply(e osm.Element) bool
}

type FilterFunc func(e osm.Element) bool

func (f FilterFunc) Apply(e osm.Element) bool {
	return f(e)
}

func All(osm.Element) bool {
	return true
}

func AllFilter() Filter {
	return FilterFunc(All)
}
