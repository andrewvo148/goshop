package model


type BasketItem struct {
	ID string
	ProductID int
	ProductName string
	UnitPrice float64
	OldUnitPrice float64
	Quantity int
	PictureURL string
}