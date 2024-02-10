package soda

import (
	"sync"
)

type DocumentIterator struct {
	response *queryResponse
	lastItem int
	mu       sync.Mutex
	err      error
}

func newDocumentIterator(q *queryResponse) *DocumentIterator {
	return &DocumentIterator{
		response: q,
		lastItem: -1,
	}
}

func (O *DocumentIterator) Next() (*DocumentSnapshot, error) {
	O.mu.Lock()
	defer O.mu.Unlock()

	if O.err != nil {
		return nil, O.err
	}

	O.lastItem++

	if O.lastItem >= len(O.response.Items) {
		O.err = ErrIteratorDone
		return nil, O.err
	}

	var docSnap = DocumentSnapshot{
		item: O.response.Items[O.lastItem],
	}

	return &docSnap, nil
}
