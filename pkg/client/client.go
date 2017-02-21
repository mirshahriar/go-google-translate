package client

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	Endpoint = "http://translate.google.com/"
	Uri      = "m"
)

type Config struct {
	Source string
	Target string
}

type Client struct {
	config     *Config
	pathPrefix string
}

type Translator struct {
	client *http.Client

	endpoint string
	uri      string

	err error
	req *http.Request
}

type Response struct {
	Err          error
	Status       string
	StatusCode   int
	ResponseBody []byte
}

func NewClient(config *Config) *Client {
	if config == nil {
		return &Client{}
	}
	pathPrefix := ""
	if config.Source != "" {
		pathPrefix = fmt.Sprintf("sl=%v", config.Source)
	}
	if config.Target != "" {
		if pathPrefix == "" {
			pathPrefix = fmt.Sprintf("tl=%v", config.Target)
		} else {
			pathPrefix = fmt.Sprintf("%v&tl=%v", pathPrefix, config.Target)
		}
	}
	c := &Client{
		config:     config,
		pathPrefix: pathPrefix,
	}
	return c
}

func (c *Client) Translate(text string) *Translator {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	escapesText := url.QueryEscape(text)

	uri := fmt.Sprintf("client=m&oe=UTF-8&ie=UTF-8")
	uri = fmt.Sprintf("%v&%v", uri, fmt.Sprintf(`text="%v"`, escapesText))
	if c.pathPrefix != "" {
		uri = fmt.Sprintf("%v&%v", uri, c.pathPrefix)
	}
	return &Translator{
		endpoint: Endpoint + Uri,
		uri:      uri,
		client:   client,
	}
}

func (t *Translator) From(lang string) *Translator {
	t.uri = fmt.Sprintf("%v&%v", t.uri, fmt.Sprintf("sl=%v", lang))
	return t
}

func (t *Translator) To(lang string) *Translator {
	t.uri = fmt.Sprintf("%v&%v", t.uri, fmt.Sprintf("to=%v", lang))
	return t
}

func (t *Translator) Get() *Translator {
	urlStr := fmt.Sprintf("%v?%v", t.endpoint, t.uri)
	t.req, t.err = http.NewRequest("GET", urlStr, nil)
	return t
}

func (t *Translator) Do() *Response {
	t.req.Header.Set("Accept", "application/json")
	resp, err := t.client.Do(t.req)
	if err != nil {
		return &Response{
			Err: err,
		}
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Err: err,
		}
	}
	return &Response{
		StatusCode:   resp.StatusCode,
		Status:       resp.Status,
		ResponseBody: responseBody,
	}
}
