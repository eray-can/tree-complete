# Tree-Complete

___

Tree-Complete aims to create a general autocomplete by supporting various types of trees. It is designed as a system that can understand data input across a wide range, encompassing different tree structures.
By providing users with a more flexible and comprehensive autocomplete experience, Tree-Complete can effectively operate on various datasets.
This approach is versatile, adapting to different scenarios, and can meet the diverse needs of users.

___

Supported Trees

| **Type**          | **Structure**                | **tree-complete** | 
|:------------------|:-----------------------------|:-----------------:| 
| [Trees](#trees)   |
|                   | RedBlackTree                 |        no         |
|                   | AVLTree                      |        yes        |
|                   | BTree                        |        no         |
|                   | BinaryHeap                   |        no         |

___

```shell
go get github.com/eray-can/tree-complete
```
[Example GetTreeCompleteResponse](./example/treecomplete/handler.go)

___

# Usage

The supported tree types are transformed into the tree-complete type for processing

___
```go

func example() {
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
}

```
___

[NewCompleteTree](./tree_complete.go) To instantiate TreeComplete, we need to create a struct.

___
```go

// MyTree: Converted tree;
// language: Language-based word Lowercase
// Prefix: search expressıon
// AsCii: Convert words used in some languages
type NewTreeCompleteSchema struct {
        MyTree *mytree.MyTree
        Lang   unicode.SpecialCase
        AsCii  []string
        Prefix string
}

//new schema
treecomplete.NewCompleteTree(schema.NewTreeCompleteSchema{
    MyTree: convertTree,
    Lang:   unicode.TurkishCase,
    AsCii:  treecomplete.TrAscii,
    Prefix: "ery",
})

func NewCompleteTree(NewSchema schema.NewTreeCompleteSchema) *ITreeComplete {
    return &ITreeComplete{Tree: NewSchema.MyTree, language: NewSchema.Lang, AsciiReplace: NewSchema.AsCii, Prefix: NewSchema.Prefix}
}


```
___

It is completed by calling the "TreeComplete" method to create an auto-complete structure.

```
ResponseBuilder: Generic bir response oluşturmak için kullanılır
Filter: istediğiniz parametreleri sorgulamak için kullanılır
```
___


```go

func example() {
    //MyTree: converted tree;
    //ResponseBuilder: Generic response;
    //Filter: filter conditions;
    type TreeCompleteSchema struct {
        MyTree          *mytree.MyTree
        Filter          func(interface{}, string) bool
        ResponseBuilder func(interface{}) interface{}
    
}


//MyTree: converted tree;
//ResponseBuilder: Generic response;
//Filter: filter conditions;
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


```