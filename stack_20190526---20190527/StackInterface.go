package stack_20190526___20190527

type Stack interface {
	Push(v interface{})
	Pop(v interface{})
	IsEmpty(v interface{})
	Top(v interface{})
	Flush()
}
