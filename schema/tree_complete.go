package schema

import (
	"tree-complete/trees/mytree"
	"unicode"
)

// MyTree: converted tree;
// ResponseBuilder: Generic response;
// Filter: filter conditions;
type TreeCompleteSchema struct {
	MyTree          *mytree.MyTree
	Filter          func(interface{}, string) bool
	ResponseBuilder func(interface{}) interface{}
}

type NewTreeCompleteSchema struct {
	MyTree *mytree.MyTree
	Lang   unicode.SpecialCase
	AsCii  []string
	Prefix string
}
