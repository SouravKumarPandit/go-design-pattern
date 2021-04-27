package shaper

type FactoryProducer struct {
}

func (p FactoryProducer) GetFactory(rounded bool) ( AbstractFactory) {
	if rounded {
		return RoundShapeFactory{}
	} else {
		return ShapeFactory{}
	}
}
