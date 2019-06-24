package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type IRetriever interface {
	Get(string) string
}

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

type IPost interface {
	Post(url string, form url.Values) string
}

type RetriverPoster interface {
	IRetriever
	IPost
}

func (me *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	return string(result)
}

func (me *Retriever) Post(strurl string, form url.Values) string {
	resp, err := http.PostForm(strurl, form)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	return string(result)
}
