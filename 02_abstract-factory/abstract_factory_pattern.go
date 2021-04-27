package main

import "design.pattern.go/02_abstract-factory/shaper"

func main() {
	sf := shaper.FactoryProducer{}
	shape := sf.GetFactory(false)
	// lazy work group enum based on shape
	rec := shape.GetShape(shaper.RECTANGLE)
	rec.Draw()
	sq := shape.GetShape(shaper.SQUARE)
	sq.Draw()

}
