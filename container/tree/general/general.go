package general

type (
	Node[T any] struct {
		Data     T
		Next     *Node[T]
		Children Children[T]
	}

	Children[T any] struct {
		Head, Tail *Node[T]
	}

	Iterator[T any] struct {
		Root *Node[T]
		Node *Node[T]
	}
)

func (node *Node[T]) PushTail(children *Children[T]) {
	tail := children.Tail

	if tail != nil {
		tail.Next = node
	} else {
		children.Head = node
	}

	children.Tail = node
}

func (node *Node[T]) PopHead(children *Children[T]) bool {
	node = children.Head

	if node == nil {
		return false
	}

	children.Head = node.Next
	node.Next = nil

	return true
}

func (node *Node[T]) Remove(children *Children[T]) {
	head := children.Head
	last := (*Node[T])(nil)

	if head == nil {
		goto clear
	}

	if head == node {
		children.Head = head.Next

		goto clear
	}

check:
	if head.Next == nil {
		goto clear
	}

	last = head
	head = head.Next

	if head == node {
		last.Next = head.Next

		if head == children.Tail {
			children.Tail = last
		}

		goto clear
	}

	goto check

clear:
	node.Next = nil
}

func (node *Node[T]) Iterator() Iterator[T] {
	return Iterator[T]{
		Root: node,
	}
}

func (сhildren *Children[T]) PushTail(data T) *Node[T] {
	tail := сhildren.Tail
	node := &Node[T]{
		Data: data,
	}

	if tail != nil {
		tail.Next = node
	} else {
		сhildren.Head = node
	}

	сhildren.Tail = node

	return node
}

func (сhildren *Children[T]) PopHead() (data T, exists bool) {
	head := сhildren.Head

	if head == nil {
		return
	}

	сhildren.Head = head.Next

	return head.Data, true
}

func (сhildren *Children[T]) Each(action func(data T) T) {
	node := сhildren.Head

	for node != nil {
		node.Data = action(node.Data)
		node = node.Next
	}
}

func (сhildren *Children[T]) Find(equal func(data, clue T) bool, clue T) (node *Node[T], found bool) {
	node = сhildren.Head

	for node != nil {
		found = equal(node.Data, clue)

		if found == false {
			node = node.Next
		} else {
			return
		}
	}

	return
}

func (children *Children[T]) Iterator() Iterator[T] {
	return Iterator[T]{
		Root: children.Head,
	}
}

func (iterator *Iterator[T]) Next() bool {
	if iterator.Node != nil {
		iterator.Node = iterator.Node.Next
	} else {
		iterator.Node = iterator.Root
	}

	if iterator.Node != nil {
		return true
	}

	return false
}
