package main

func main() {
	stock :=NewStock()
	buyStock := NewBuyStock(stock)
	sellStock := NewSellStock(stock)

	broker:=NewBroker()
	broker.takeOrder(&buyStock)
	broker.takeOrder(&sellStock)
	broker.placeOrder()
}