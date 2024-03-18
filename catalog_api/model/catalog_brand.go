package model

type CatalogBrand struct {
	ID int
	Brand string
}

func NewCatalogBrand(id int, brand string) *CatalogBrand {
	return &CatalogBrand{
		ID: id,
		Brand: brand,
	}
}