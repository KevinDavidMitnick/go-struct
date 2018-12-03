package list

type List interface {
	IsEmpty() bool
	Length() int
	Size() int
	Equal(list List) bool
	IsFull() bool
	GetElem(i int) (interface{}, error)
	LocateElem(e interface{}) int
	Insert(i int, e interface{}) error
	Delete(i int) (interface{}, error)
	Clone() List
	Diff(list List) List
	Union(list List) List
	ClearList()
	Display()
}
