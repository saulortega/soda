package soda

import (
	"encoding/json"
	"time"
)

type queryResponse struct {
	Items []*queryResponseItem `json:"items"`
}

type queryResponseItem struct {
	ID           string          `json:"id"`
	ETag         string          `json:"etag"`
	LastModified time.Time       `json:"lastModified"`
	Created      time.Time       `json:"created"`
	Value        json.RawMessage `json:"value"`
	Links        []*ItemLink     `json:"links"`
}

type ItemLink struct {
	Rel  string `json:"rel"`
	HRef string `json:"href"`
}
