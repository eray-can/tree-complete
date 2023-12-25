package treecomplete

import (
	"encoding/json"
	"fmt"
	"github.com/emirpasic/gods/trees/avltree"
	treecomplete "github.com/eray-can/tree-complete"
	"github.com/eray-can/tree-complete/mock"
	"github.com/eray-can/tree-complete/schema"
	myTree "github.com/eray-can/tree-complete/trees/avltree"
	"net/http"

	"unicode"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	prefix := r.URL.Query().Get("q")
	if prefix == "" {
		http.Error(w, "q url param empty or not found", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(GetTreeCompleteResponse(prefix)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("success")
}

func GetTreeCompleteResponse(prefix string) []interface{} {

	//example mock data for example
	MockDataList := mock.GetTreeMockData()

	//new avl tree for example
	newAvlTree := avltree.NewWithIntComparator()

	//set tree mock data for example
	for idx := range MockDataList {
		newAvlTree.Put(idx, MockDataList[idx])
	}

	//convert tree-complete
	convertTree := myTree.ConvertAVLToMyTree(newAvlTree)

	//new tree complete
	newCompleteTree := treecomplete.NewCompleteTree(schema.NewTreeCompleteSchema{
		MyTree: convertTree,
		Lang:   unicode.TurkishCase,
		AsCii:  treecomplete.TrAscii,
		Prefix: prefix,
	})

	//MyTree: converted tree, ResponseBuilder: Generic response, Filter: filter conditions
	newCompleteTree.TreeComplete(schema.TreeCompleteSchema{
		MyTree: convertTree,
		Filter: func(value interface{}, prefix string) bool {
			response := value.(*mock.TestTree)
			return newCompleteTree.ContainsSubstringIgnoreCase(response.Name) || newCompleteTree.ContainsSubstringIgnoreCase(response.Surname)
		},
		ResponseBuilder: func(value interface{}) interface{} {
			response := value.(*mock.TestTree)
			return &myResponse{
				Name:    response.Name,
				Surname: response.Surname,
			}
		},
	})

	//get tree response
	return newCompleteTree.GetResponse()
}

// my custom response
type myResponse struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
