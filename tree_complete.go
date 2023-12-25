package tree_complete

import (
	"strings"
	"tree-complete/schema"
	"tree-complete/trees/mytree"
	"tree-complete/utils"
	"unicode"
)

type TreeCompleteContract interface {
	TreeComplete(TreeCompleteSchema schema.TreeCompleteSchema)
	GetResponse() []interface{}
}

// Tree: Converted tree;
// Response: Your created jeneric response
// language: Language-based word Lowercase
// Prefix: search expressıon
// AsciiReplace: Convert words used in some languages
type ITreeComplete struct {
	Tree         *mytree.MyTree
	Response     []interface{}
	language     unicode.SpecialCase
	Prefix       string
	AsciiReplace []string
}

// lang ascii
var (
	TrAscii = []string{
		"ö", "o",
		"ç", "c",
		"ğ", "g",
		"ü", "u",
		"ş", "s",
		"i", "ı",
	}
)

// It checks all the data within the tree and returns the ones that meet the criteria in the built struct.
func (t *ITreeComplete) TreeComplete(TreeCompleteSchema schema.TreeCompleteSchema) {

	it := TreeCompleteSchema.MyTree.Iterator()
	for it.Next() {
		value := it.Value()
		if value != nil && TreeCompleteSchema.Filter(value, t.Prefix) {
			item := TreeCompleteSchema.ResponseBuilder(value)
			t.Response = append(t.Response, item)
		}
	}

}

// Get struct response
func (t *ITreeComplete) GetResponse() []interface{} {
	return t.Response
}

//TODO More advanced algorithms will come in the future.

// This function checks for the presence of a specific substring within a main string, considering case sensitivity.
// If the substring is found, it returns true; otherwise, it returns false. During this check,
// ASCII characters may be replaced according to certain rules.
func (t *ITreeComplete) ContainsSubstringIgnoreCase(s string) bool {
	var prefix string
	if len(t.AsciiReplace) > 0 {
		prefix = t.replaceAscii(t.Prefix)
		s = t.replaceAscii(s)
	}

	s, prefix = utils.ToLowerByLang(s, t.language), utils.ToLowerByLang(prefix, t.language)
	index := 0

	for _, r := range s {
		if index < len(prefix) && byte(r) == prefix[index] {
			index++
		}
	}
	return index == len(prefix)
}

// replace by language
func (t *ITreeComplete) replaceAscii(prefix string) string {
	replacer := strings.NewReplacer(t.AsciiReplace...)
	return replacer.Replace(prefix)
}

// New Tree struct
func NewCompleteTree(NewSchema schema.NewTreeCompleteSchema) *ITreeComplete {
	return &ITreeComplete{Tree: NewSchema.MyTree, language: NewSchema.Lang, AsciiReplace: NewSchema.AsCii, Prefix: NewSchema.Prefix}
}
