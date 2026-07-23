package core_postgres_pool

import "errors"

var (
	ErrNoRows            = errors.New("no rows")
	ErrViolateForeignKey = errors.New("violate foreign key")
	ErrUnknown           = errors.New("unkmown")
)
