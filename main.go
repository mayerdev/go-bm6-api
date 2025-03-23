package bm6

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type BILLmanager struct {
	url      string
	username string
	password string
	client   *http.Client
}

func New(url, username, password string) *BILLmanager {
	return &BILLmanager{
		url:      url,
		username: username,
		password: password,
		client:   &http.Client{},
	}
}

func (b *BILLmanager) buildGetParams(data map[string]string) string {
	params := url.Values{}

	for key, value := range data {
		params.Add(key, value)
	}

	return params.Encode()
}

func (b *BILLmanager) Request(data map[string]string, authorized bool) ([]byte, error) {
	if authorized {
		data["authinfo"] = fmt.Sprintf("%s:%s", b.username, b.password)
	}

	if _, exists := data["out"]; !exists {
		data["out"] = "json"
	}

	// Некоторые CDN/ддос-защиты кэшируют API, это нужно для обхода кэширования
	data["time"] = time.Now().String()

	u := fmt.Sprintf("%s/?%s", b.url, b.buildGetParams(data))
	fmt.Println(u)

	resp, err := b.client.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
