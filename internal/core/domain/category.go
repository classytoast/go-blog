package domain

import (
	"fmt"

	core_errors "github.com/classytoast/go-blog/internal/core/errors"
)

type Category struct {
	ID   int64
	Name string
}

func NewCategory(
	id int64,
	name string,
) Category {
	return Category{
		ID:   id,
		Name: name,
	}
}

func NewCategoryUninitialized(name string) Category {
	return Category{
		ID:   int64(UninitializedID),
		Name: name,
	}
}

func (c *Category) Validate() error {
	nameLen := len([]rune(c.Name))
	if nameLen < 3 || nameLen > 200 {
		return fmt.Errorf(
			"invalid `name` len: %d: %w",
			nameLen,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
