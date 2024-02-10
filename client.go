package soda

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	url        string
	username   string
	password   string
	httpClient *http.Client
}

func NewWithFullURL(url string, username string, password string) *Client {
	return &Client{
		url:      url,
		username: username,
		password: password,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (O *Client) formatURLForQuery(collection string) string {
	return fmt.Sprintf(`%s/%s?action=query`, O.url, collection)
}

func (O *Client) request(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	var authPairPlain = fmt.Sprintf("%s:%s", O.username, O.password)
	var authPair = base64.StdEncoding.EncodeToString([]byte(authPairPlain))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", authPair))

	return req, nil
}

func (O *Client) Query(collection string, filter map[string]interface{}) (*DocumentIterator, error) {
	var bodyFilter io.Reader

	if filter != nil {
		resJSON, err := json.Marshal(filter)
		if err != nil {
			return nil, err
		}

		bodyFilter = bytes.NewReader(resJSON)
	}

	req, err := O.request("POST", O.formatURLForQuery(collection), bodyFilter)
	if err != nil {
		return nil, err
	}

	res, err := O.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var errTxt = res.Status
		if len(body) > 0 {
			errTxt = string(body)
		}

		return nil, errors.New(errTxt)
	}

	var R queryResponse
	err = json.Unmarshal(body, &R)
	if err != nil {
		return nil, err
	}

	return newDocumentIterator(&R), nil
}

func (O *Client) GetByID(collection string, id string) (doc *DocumentSnapshot, err error) {
	iter, err := O.Query(collection, map[string]interface{}{"_id": id})
	if err != nil {
		return nil, err
	}

	for {
		doc, err = iter.Next()
		if err == ErrIteratorDone {
			return nil, ErrNotFound
		}
		if err != nil {
			return nil, err
		}

		break
	}

	return doc, nil
}
