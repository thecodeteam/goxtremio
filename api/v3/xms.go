package apiv3

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
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

func New(endpoint string, insecure bool, username string, password string) (*XMS, error) {
	if endpoint == "" || username == "" || password == "" {
		return nil, errors.New("Missing endpoint, username, or password")
	}

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

	return &XMS{endpoint, insecure, username, password, client}, nil
}

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		if v != "" {
			q[k] = []string{v}
		}
	}
	return q
}

type Error struct {
	StatusCode int
	Message    string `json:"message"`
	Code       int    `json:"error_code"`
}

func (err *Error) Error() string {
	if err.Code == 0 {
		return err.Message
	}

	return fmt.Sprintf("%s", err.Message)
}

func buildError(r *http.Response) error {
	jsonError := Error{}
	json.NewDecoder(r.Body).Decode(&jsonError)

	jsonError.StatusCode = r.StatusCode
	if jsonError.Message == "" {
		jsonError.Message = r.Status
	}

	return &jsonError
}

func (xms *XMS) query(method string, path string, id string, params map[string]string, body interface{}, resp interface{}) error {

	endpoint := fmt.Sprintf("%s/%s", xms.endpoint, path)
	if id != "" {
		endpoint = fmt.Sprintf("%s/%s", endpoint, id)
	}

	encodedParams := multimap(params).Encode()
	if encodedParams != "" {
		endpoint = fmt.Sprintf("%s?%s", endpoint, encodedParams)
	}

	bodyBytes, _ := json.Marshal(body)

	if debug {
		log.Printf("endpoint: %v\n", endpoint)
	}

	if debug && body != nil {
		log.Printf("body:\n")
		log.Printf("%v\n", string(bodyBytes))
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(bodyBytes))
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

	switch {
	case r.StatusCode == 400:
		return buildError(r)
	case resp == nil:
		return nil
	case r.StatusCode == 200 || r.StatusCode == 201:
		err = json.NewDecoder(r.Body).Decode(resp)
		return err
	default:
		return buildError(r)
	}

}
