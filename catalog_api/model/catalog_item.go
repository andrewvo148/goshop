package model

type CatalogItem struct {
	ID int
	Name string
	Description string
	Price float64
	PictureURL string
	CatalogTypeID int
	CatalogType CatalogType
	CatalogBrandId int
	CatalogBrand CatalogBrand
}