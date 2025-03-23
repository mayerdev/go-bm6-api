package bm6

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type BILLmanager struct {
	url             string
	username        string
	password        string
	client          *http.Client
	logging_enabled bool
}

func New(url, username, password string, logging_enabled bool) *BILLmanager {
	return &BILLmanager{
		url:             url,
		username:        username,
		password:        password,
		client:          &http.Client{},
		logging_enabled: logging_enabled,
	}
}

func (ctx *BILLmanager) buildGetParams(data map[string]string) string {
	params := url.Values{}

	for key, value := range data {
		params.Add(key, value)
	}

	return params.Encode()
}

func (ctx *BILLmanager) Request(data map[string]string, authorized bool) ([]byte, error) {
	if authorized {
		data["authinfo"] = fmt.Sprintf("%s:%s", ctx.username, ctx.password)
	}

	if _, exists := data["out"]; !exists {
		data["out"] = "json"
	}

	// Некоторые CDN/ддос-защиты кэшируют API, это нужно для обхода кэширования
	data["time"] = time.Now().String()

	u := fmt.Sprintf("%s/?%s", ctx.url, ctx.buildGetParams(data))

	if ctx.logging_enabled {
		fmt.Println(u)
	}

	resp, err := ctx.client.Get(u)
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
