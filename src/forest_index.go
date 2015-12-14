package lsh

type TreeNode struct {
	// Hash key for this intermediate node. nil/empty for root nodes.
	hashKey TableKey
	// A list of indices to the source dataset.
	indices []int
	// Child nodes.
	children []TreeNode
}

type Tree struct {
	// Number of distinct elements in the tree.
	count int
	// Pointer to the root node.
	root TreeNode
}

type ForestIndex struct {
	// Embedded type
	*LshSettings
	// Number of leaves in the tree.
	count int
	// Trees.
	trees []Tree
}

func NewLshForest(dim, l, m int, w float64) *ForestIndex {
	trees := make([]Tree, l)
	return &ForestIndex{
		LshSettings: NewLshSettings(m, l, dim, w),
		count:       0,
		trees:       trees,
	}
}

// Inserts a point into the index.
func (index *ForestIndex) Insert(point Point) {
	// Apply hash functions.
	//	hvs := index.Hash(point)

}
