package soda

import (
	"encoding/json"
	"time"
)

type DocumentSnapshot struct {
	item *queryResponseItem
}

func (O *DocumentSnapshot) ID() string {
	return O.item.ID
}

func (O *DocumentSnapshot) ETag() string {
	return O.item.ETag
}

func (O *DocumentSnapshot) Created() time.Time {
	return O.item.Created
}

func (O *DocumentSnapshot) LastModified() time.Time {
	return O.item.LastModified
}

func (O *DocumentSnapshot) Links() []*ItemLink {
	return O.item.Links
}

func (O *DocumentSnapshot) Data() []byte {
	return O.item.Value
}

func (O *DocumentSnapshot) JSONDataTo(v interface{}) error {
	return json.Unmarshal(O.Data(), v)
}
