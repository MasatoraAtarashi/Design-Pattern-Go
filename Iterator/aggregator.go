package main

type Aggregator interface {
	Iterator() Iterator
}

type BookShelf struct {
	Books []Book
}

func (b BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(b)
}
