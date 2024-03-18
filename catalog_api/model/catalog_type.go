package model

type CatalogType struct {
	ID int
	Type string
}

func NewCatalogType(id int, t string) *CatalogType {
	return &CatalogType{
		ID: id,
		Type: t,
	}
}