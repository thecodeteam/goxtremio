package goxtremio

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type XMS struct {
	endpoint   string
	insecure   bool
	username   string
	password   string
	httpClient *http.Client
}

func New(endpoint string, insecure bool, username string, password string) *XMS {
	var client *http.Client
	if insecure {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecure,
				},
			},
		}
	} else {
		client = &http.Client{}
	}

	return &XMS{endpoint, insecure, username, password, client}
}

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}

func (xms *XMS) query(path string, id string, params map[string]string, resp interface{}) error {

	endpoint := fmt.Sprintf("%s/%s", xms.endpoint, path)
	if id != "" {
		endpoint = fmt.Sprintf("%s/%s", endpoint, id)
	}

	encodedParams := multimap(params).Encode()
	if encodedParams != "" {
		endpoint = fmt.Sprintf("%s?%s", endpoint, encodedParams)
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(xms.username, xms.password)

	r, err := xms.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if debug {
		dump, _ := httputil.DumpResponse(r, true)
		log.Printf("response:\n")
		log.Printf("%v\n}\n", string(dump))
	}

	// if r.StatusCode != 200 {
	// 	return buildError(r)
	// }

	err = json.NewDecoder(r.Body).Decode(resp)
	return err
}
