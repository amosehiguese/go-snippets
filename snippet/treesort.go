package snippet

type Tree struct {
	value       int
	left, right *Tree
}

func TreeSort(values []int) {
	var root *Tree
	for _, v := range values {
		root = AddTree(root, v)
	}
	AppendValues(values[:0], root)
}

func AppendValues(values []int, t *Tree) []int {
	if t != nil {
		values = AppendValues(values, t.left)
		values = append(values, t.value)
		values = AppendValues(values, t.right)
	}
	return values
}

func AddTree(t *Tree, value int) *Tree {
	if t == nil {
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = AddTree(t.left, value)
	} else {
		t.right = AddTree(t.right, value)
	}
	return t
}