package core_http_types

import (
	"encoding/json"

	"github.com/mihail2771/todogo/iternal/core/domain"
)

type Nullabel[T any] struct {
	domain.Nullabel[T]
}

func (n *Nullabel[T]) UnmarshalJSON(b []byte) error {
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

func (n *Nullabel[T]) ToDomain() domain.Nullabel[T] {
	return domain.Nullabel[T]{
		Value: n.Value,
		Set:   n.Set,
	}
}
