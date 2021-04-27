package main

type Memento struct {
	state string
}

func NewMemento(state string) Memento {
	return Memento{state: state}
}
func (m Memento) getState() string {
	return m.state
}

type Originator struct {
	state string
}

func NewOriginator() Originator{
	return Originator{}
}

func (o *Originator) setState(state string) {
	o.state=state
}

func (o *Originator) getState() string {
	return o.state
}
func (o *Originator) saveStateToMemento() Memento {
	return Memento{
		state: o.getState(),
	}
}


func (o *Originator) getStateFromMemento(memento Memento)  {
	o.state=memento.state
}

type CareTaker struct {
	mementoList []Memento
}

func NewCareTaker() CareTaker {
	return CareTaker{}
}

func (c *CareTaker) add(memento Memento) {
	c.mementoList=append(c.mementoList, memento)
}
func (c *CareTaker) get(index int) Memento {
	return c.mementoList[index]
}