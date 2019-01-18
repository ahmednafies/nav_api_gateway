package request

import (
	"bytes"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/errorhandler"
	"io/ioutil"
	"net/http"
	"net/url"
)

func clientRequest(req *http.Request, method string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	if method == "POST" {
		if resp.StatusCode != http.StatusCreated {
			return nil, errorhandler.Handle(resp.StatusCode, respBody)
		}
	} else {
		if resp.StatusCode != http.StatusOK {
			return nil, errorhandler.Handle(resp.StatusCode, respBody)

		}
	}
	return respBody, nil
}

func headers(uri string, method string, body []byte) *http.Request {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	u.RawQuery = u.Query().Encode()
	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(body))
	req.SetBasicAuth(config.Username, config.Passwd)
	req.Header.Add("If-Match", "*")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json; odata.metadata=minimal")
	return req

}

func request(uri string, method string, body []byte) ([]byte, error) {
	req := headers(uri, method, body)
	return clientRequest(req, method)
}