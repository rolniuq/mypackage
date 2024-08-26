package httpbuilder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type httpMethod string

const (
	GET    httpMethod = "GET"
	POST   httpMethod = "POST"
	PUT    httpMethod = "PUT"
	PATCH  httpMethod = "PATCH"
	DELETE httpMethod = "DELETE"
)

type Ops struct {
	url     *url.URL
	headers []map[string]string
	method  httpMethod
}

func (o *Ops) WithMethod(method httpMethod) *Ops {
	o.method = method

	return o
}

func (o *Ops) WithUrl(url *url.URL) *Ops {
	o.url = url

	return o
}

func (o *Ops) WithHeaders(headers ...map[string]string) *Ops {
	o.headers = headers

	return o
}

type HttpRequestBuilder struct {
	options []Ops
}

func NewHttpRequestBuilder(ops ...Ops) *HttpRequestBuilder {
	return &HttpRequestBuilder{ops}
}

func (h *HttpRequestBuilder) Build() (*http.Request, error) {
	if len(h.options) == 0 {
		return nil, fmt.Errorf("no options provided")
	}

	opt := h.options[0]

	if opt.method == "" {
		return nil, fmt.Errorf("missing method")
	}

	req, err := http.NewRequest(string(opt.method), opt.url.String(), nil)
	if err != nil {
		return nil, err
	}

	for _, header := range opt.headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	return req, nil
}

func Send[T any](req *http.Request) (*T, error) {
	var result T

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	u, _ := url.Parse("https://jsonplaceholder.typicode.com/todos/1")

	ops := &Ops{}
	reqBuilder := NewHttpRequestBuilder(
		*ops.WithUrl(u),
		*ops.WithMethod(GET),
		*ops.WithHeaders(map[string]string{"Content-Type": "application/json"}),
	)

	req, err := reqBuilder.Build()
	if err != nil {
		fmt.Println("Error building request:", err)
		return
	}
	if req == nil {
		fmt.Println("request is nil")
		return
	}

	type Todo struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	response, err := Send[Todo](req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	fmt.Println("Id: ", response.ID)
	fmt.Println("UserId: ", response.UserID)
}
