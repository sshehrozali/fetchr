package main

import "net/http"

type DefaultHttpClient struct{}

// implmenets HttpClient interface Get() method
func (c *DefaultHttpClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
