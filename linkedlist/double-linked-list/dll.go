package doublelinkedlist

type Element struct {
	Val  int
	Prev *Element
	Next *Element
}

type List struct {
	Head   *Element
	Length int
}

func New() *List {
	myList := List{nil, 0}
	return &myList
}