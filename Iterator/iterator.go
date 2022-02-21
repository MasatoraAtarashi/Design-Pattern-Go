package main

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookShelfIterator struct {
	BookShelf BookShelf
	Index     int
}

func NewBookShelfIterator(shelf BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		BookShelf: shelf,
		Index:     0,
	}
}

func (b *BookShelfIterator) HasNext() bool {
	return b.Index < len(b.BookShelf.Books)
}

func (b *BookShelfIterator) Next() interface{} {
	book := b.BookShelf.Books[b.Index]
	b.Index++
	return book
}
