package main

import "fmt"

type DependentObject1 struct {
	data string
}


type DependentObject2 struct {
	data string
}

type CoarseGainedObject struct {
	do1 DependentObject1
	do2 DependentObject2
}

func (o CoarseGainedObject) getData() []string {
	return []string{o.do1.data,o.do2.data}
}

func (o *CoarseGainedObject) setData(data1 string, data2 string) {
	o.do1.data=data1
	o.do2.data=data2
}

type CompositeEntity struct {
	cgo *CoarseGainedObject
}

func (o *CompositeEntity) setData(data1 , data2 string)  {
	o.cgo.setData(data1,data2)

}
func (o CompositeEntity) getData() []string {
	return o.cgo.getData()
}

type Client struct {
	composeEntity *CompositeEntity
}

func NewClient() Client {
	return Client{composeEntity: &CompositeEntity{cgo: &CoarseGainedObject{}}}
}
func (c *Client) printData() {
	for i,_ := range c.composeEntity.cgo.getData() {
		fmt.Printf("Data : %v\n",c.composeEntity.getData()[i] )
	}
}

func (c *Client) setData(data1 ,data2 string) {
	c.composeEntity.setData(data1,data2)
}