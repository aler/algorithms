package data

// order of grow
// average: put ~ lg N, get ~ lg N
// worst:   put ~ N,    get ~ N

// BST is a symbol table implemented with binary search tree
type BST struct {
	root *bstNode
}

type bstNode struct {
	key, val    int
	left, right *bstNode
	sz          int
}

// NewBST creates a new BST data type
func NewBST() *BST {
	return &BST{root: nil}
}

// Put adds a new key value pair if key is not already there
// or update existing key with provided val value
func (s *BST) Put(key, val int) {
	s.root = s.root.put(key, val)
}

func (node *bstNode) put(key, val int) *bstNode {
	if node == nil {
		return &bstNode{
			key:   key,
			val:   val,
			left:  nil,
			right: nil,
			sz:    1,
		}
	}
	if key < node.key {
		node.left = node.left.put(key, val)
	} else if key > node.key {
		node.right = node.right.put(key, val)
	} else {
		node.val = val
	}
	node.sz = node.left.size() + node.right.size() + 1
	return node
}

func (node *bstNode) size() int {
	if node == nil {
		return 0
	}
	return node.sz
}

// Get returns value for a key if key is found
// If key is not found it return false ok flag
func (s *BST) Get(key int) (int, bool) {
	node := s.root
	for node != nil {
		if key < node.key {
			node = node.left
		} else if key > node.key {
			node = node.right
		} else {
			return node.val, true
		}
	}
	return 0, false
}

// Size return how many nodes are in the tree including root
func (s *BST) Size() int {
	if s.root == nil {
		return 0
	}
	return s.root.sz
}

// Min returns minimum key in the symbol table
func (s *BST) Min() int {
	if s.root == nil {
		panic("no max for empty tree")
	}
	return s.root.min().key
}

func (node *bstNode) min() *bstNode {
	if node.left == nil {
		return node
	}
	return node.left.min()
}

// Max returns maximum key in the symbol table
func (s *BST) Max() int {
	if s.root == nil {
		panic("no max for empty tree")
	}
	return s.root.max().key
}

func (node *bstNode) max() *bstNode {
	if node.right == nil {
		return node
	}
	return node.right.max()
}

// Floor returns key for greates element
// that is less or equal then key
func (s *BST) Floor(key int) (int, bool) {
	node := s.root.floor(key)
	if node == nil {
		return 0, false
	}
	return node.key, true
}

// to ceil replace < with > and left with right
func (node *bstNode) floor(key int) *bstNode {
	if node == nil {
		return nil
	}
	if key == node.key {
		return node
	}
	if key < node.key {
		return node.left.floor(key)
	}
	n := node.right.floor(key)
	// if right subtree has node less or equal to key return it
	if n != nil {
		return n
	}
	// otherwise return current
	return node
}

// Select returns key that is n's smallest in the symbol table
func (s *BST) Select(n int) int {
	if s.root == nil {
		panic("cannot select in an empty BSD")
	}
	return s.root.sel(n).key
}

func (node *bstNode) sel(n int) *bstNode {
	if node == nil {
        return nil
    }
	lsz := node.left.size()
	if n < lsz {
		return node.left.sel(n)
	} else if n > lsz {
		return node.right.sel(n - lsz - 1)
	} else {
		return node
	}
}

// Rank returns number of keys that a less then key
func (s *BST) Rank(key int) int {
    if s.root == nil {
        panic("tree is empty")
    }
    return s.root.rank(key)
}

func (node *bstNode) rank(key int) int {
    if node == nil {
        return 0
    }
    if key < node.key {
        return node.left.rank(key)
    } else if key > node.key {
        return 1 + node.left.size() + node.right.rank(key)
    } else {
        return node.left.size()
    }
}

// Keys returns an array of all keys in order
func (s *BST) Keys() []int {
    var keys []int
    return s.root.inorder(keys)
}

func (node *bstNode) inorder(a []int) []int {
    if node == nil {
        return a
    }
    l := node.left.inorder(a)
    m := append(l, node.key)
    r := node.right.inorder(m)
    return r
}

// todo: delete
