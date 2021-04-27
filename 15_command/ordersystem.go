package main

import "fmt"

type Order interface {
	execute()
}

type Stock struct {
	name     string
	quantity int
}

func NewStock() Stock {
	return Stock{name: "ABC", quantity: 10}
}
func (stock *Stock) buy() {
	fmt.Printf("Stock [ Name : %v , Quantitiy: %v ] bought \n", stock.name, stock.quantity)
}
func (stock *Stock) sell() {
	fmt.Printf("Stock [ Name : %v , Quantitiy: %v ] sold  \n", stock.name, stock.quantity)

}

type BuyStock struct {
	abcStock Stock
}

func NewBuyStock(stock Stock) BuyStock {
	return BuyStock{stock}
}
func (bs *BuyStock) execute() {
	bs.abcStock.buy()
}

type SellStock struct {
	abcStock Stock
}

func NewSellStock(stock Stock) SellStock {
	return SellStock{stock}
}
func (bs *SellStock) execute() {
	bs.abcStock.sell()
}

type Broker struct {
	orderList []Order
}

func NewBroker() Broker {
	return Broker{}
}

func (b *Broker) takeOrder(order Order) {
	b.orderList = append(b.orderList, order)

}
func (b Broker) placeOrder() {
	for _, order := range b.orderList {
		order.execute()
	}
	b.orderList = []Order{}
}
