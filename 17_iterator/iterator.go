package main

type Iterator interface {
	hasNext() bool
	next() interface{}
}

type Container interface {
	getIterator() Iterator
}

type NameRepository struct {
	NewIterator
}

func NewNameRepository(names []string) NameRepository {
	return NameRepository{NewIterator{names: names, index: 0}}
}
func (itr NameRepository) getIterator() NewIterator {
	iterator := itr.NewIterator

	return iterator
}

type NewIterator struct {
	names []string
	index int
}

func (itr *NewIterator) hasNext() bool {
	if itr.index < len(itr.names) {
		return true
	}
	return false
}

func (itr *NewIterator) next() interface{} {
	if itr.hasNext() {

		s := itr.names[itr.index]
		itr.index++
		return s
	}
	return nil
}
