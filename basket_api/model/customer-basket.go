package model

type CustomerBasket struct {
	BuyerID string
	Items []BasketItem
}

func NewCustomerBasket(customerID string) *CustomerBasket {
	return &CustomerBasket{
		BuyerID: customerID,
		Items: make([]BasketItem, 0),
	}
}