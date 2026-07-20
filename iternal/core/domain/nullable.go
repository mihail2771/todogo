package domain

type Nullabel[T any] struct {
	Value *T
	Set   bool
}
