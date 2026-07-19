package domain

type Nullebel[T any] struct {
	Value *T
	Set   bool
}
