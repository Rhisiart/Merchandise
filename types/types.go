package types

type IOperation interface {
	Create() error
	Read(id int64) error
	Update() error
	Delete() error
}

type Design struct {
	DesignId    int64
	Name        string
	Description string
}

type ClothingType struct {
	ClothingTypeId int64
	Name           string
	Description    string
}

type Clothing struct {
	ClothingId int64
	DesignId   int64
	Name       string
	Price      int
}
