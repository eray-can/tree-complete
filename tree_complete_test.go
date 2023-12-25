package tree_complete

import (
	"encoding/json"
	"github.com/emirpasic/gods/trees/avltree"
	"github.com/eray-can/tree-complete/mock"
	"github.com/eray-can/tree-complete/schema"
	avltree2 "github.com/eray-can/tree-complete/trees/avltree"
	"testing"
	"unicode"
)

func TestTreeComplete_TreeComplete(t *testing.T) {
	MockDataList := mock.GetTreeMockData()

	newAvlTree := avltree.NewWithIntComparator()
	for idx := range MockDataList {
		newAvlTree.Put(idx, MockDataList[idx])
	}

	convertTree := avltree2.ConvertAVLToMyTree(newAvlTree)
	newTreeComplete := schema.NewTreeCompleteSchema{
		MyTree: convertTree,
		Lang:   unicode.TurkishCase,
		AsCii:  TrAscii,
		Prefix: "ery",
	}

	newCompleteTree := NewCompleteTree(newTreeComplete)

	treeCompleteSchema := schema.TreeCompleteSchema{
		MyTree: convertTree,
		ResponseBuilder: func(value interface{}) interface{} {
			response := value.(*mock.TestTree)
			return &myTestResponse{
				Name:    response.Name,
				Surname: response.Surname,
			}
		},
		Filter: func(value interface{}, prefix string) bool {
			response := value.(*mock.TestTree)
			return newCompleteTree.ContainsSubstringIgnoreCase(response.Name) || newCompleteTree.ContainsSubstringIgnoreCase(response.Surname)
		},
	}

	newCompleteTree.TreeComplete(treeCompleteSchema)
	responseJSON, err := json.Marshal(newCompleteTree.GetResponse())
	if err != nil {
		t.Log("Marshal Err: ", err.Error())
	}

	t.Log("Response: ", string(responseJSON))
}

type myTestResponse struct {
	Name    string
	Surname string
}
