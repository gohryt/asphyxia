package double

type (
	List[T any] struct {
		Head, Tail *Node[T]
	}

	Node[T any] struct {
		Previous, Next *Node[T]
		Data           T
	}

	Iterator[T any] struct {
		Root *Node[T]
		Node *Node[T]
	}
)

func (list *List[T]) PushHead(data T) *Node[T] {
	head := list.Head
	node := &Node[T]{
		Next: head,
		Data: data,
	}

	if head != nil {
		head.Previous = node
	} else {
		list.Tail = node
	}

	list.Head = node

	return node
}

func (list *List[T]) PushTail(data T) *Node[T] {
	tail := list.Tail
	node := &Node[T]{
		Previous: tail,
		Data:     data,
	}

	if tail != nil {
		tail.Next = node
	} else {
		list.Head = node
	}

	list.Tail = node

	return node
}

func (list *List[T]) PopHead() (data T, exists bool) {
	head := list.Head

	if head == nil {
		return
	}

	list.Head = head.Next

	return head.Data, true
}

func (list *List[T]) PopTail() (data T, exists bool) {
	tail := list.Tail

	if tail == nil {
		return
	}

	list.Tail = tail.Previous

	return tail.Data, true
}

func (list *List[T]) Each(action func(data T) T) {
	node := list.Head

	for node != nil {
		node.Data = action(node.Data)
		node = node.Next
	}
}

func (list *List[T]) Find(equal func(data, clue T) bool, clue T) (node *Node[T], found bool) {
	node = list.Head

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

func (list *List[T]) Iterator() Iterator[T] {
	return Iterator[T]{
		Root: list.Head,
	}
}

func (node *Node[T]) PushHead(list *List[T]) {
	head := list.Head

	if head != nil {
		head.Previous = node
	} else {
		list.Tail = node
	}

	node.Next = head
	list.Head = node
}

func (node *Node[T]) PushTail(list *List[T]) {
	tail := list.Tail

	if tail != nil {
		tail.Next = node
	} else {
		list.Head = node
	}

	node.Previous = tail
	list.Tail = node
}

func (node *Node[T]) PopHead(list *List[T]) bool {
	node = list.Head

	if node == nil {
		return false
	}

	list.Head = node.Next
	node.Next = nil

	return true
}

func (node *Node[T]) PopTail(list *List[T]) bool {
	node = list.Tail

	if node == nil {
		return false
	}

	list.Tail = node.Previous
	node.Previous = nil

	return true
}

func (node *Node[T]) Remove(list *List[T]) {
	head := list.Head

	if head == nil {
		goto clear
	}

	if head == node {
		list.Head = head.Next

		goto clear
	}

check:
	if head.Next == nil {
		goto clear
	}

	head = head.Next

	if head == node {
		head.Previous.Next = head.Next

		if head == list.Tail {
			list.Tail = head.Previous
		} else {
			head.Next.Previous = head.Previous
		}

		goto clear
	}

	goto check

clear:
	node.Previous = nil
	node.Next = nil
}

func (node *Node[T]) Iterator() Iterator[T] {
	return Iterator[T]{
		Root: node,
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
