package vendorapi

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func httpRequest(method, url string, getParams map[string]string, postParams io.Reader) ([]byte, error) {

	client := http.Client{
		Timeout: time.Second,
	}

	request, err := http.NewRequest(method, url, postParams)

	q := request.URL.Query()
	for k, v := range getParams {
		q.Add(k, v)
	}

	request.URL.RawQuery = q.Encode()

	if err != nil {
		return nil, err
	}
	resp, err := client.Do(request)
	if resp != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err := request.Body.Close(); err != nil {
		return nil, err
	}

	return body, err
}
