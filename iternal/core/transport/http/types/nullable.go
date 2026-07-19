package core_http_types

import (
	"encoding/json"

	"github.com/mihail2771/todogo/iternal/core/domain"
)

type Nullebel[T any] struct {
	domain.Nullebel[T]
}

func (n *Nullebel[T]) UnmarshalJSON(b []byte) error {
	n.Set = true
	if string(b) == "null" {
		n.Value = nil

		return nil
	}

	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	n.Value = &value
	return nil
}

func (n *Nullebel[T]) ToDomain() domain.Nullebel[T] {
	return domain.Nullebel[T]{
		Value: n.Value,
		Set:   n.Set,
	}
}
