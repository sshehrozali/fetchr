package network

import "net/http"

type HttpClient struct{}

func (h *HttpClient) httpGet(url string) (*http.Response, error) {
	return http.Get(url)
}