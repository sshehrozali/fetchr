package network

import "net/http"

type HttpClient struct{}

func (h *HttpClient) HttpGet(url string) (*http.Response, error) {
	return http.Get(url)
}