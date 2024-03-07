package db

import (
	types "github.com/Rhisiart/Merchandise/types"
)

type DesignOperation struct {
	design *types.Design
}

func (design *DesignOperation) Create() error {
	return nil
}
