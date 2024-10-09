package expressgo

import (
	"io"
	"net/http"
	"net/url"

	"github.com/ContentGamer/expressgo/utils"
)

type Request struct {
	data *http.Request
	resp *Response

	Method  string
	URL     string
	Headers map[string]string
	Body    *Body
	Params  map[string]string
	Query   url.Values
}

func (req *Request) setupVariables() {
	req.Method = req.data.Method
	req.URL = req.data.URL.String()

	req.Headers = make(map[string]string)
	for n, v := range req.data.Header {
		req.Headers[n] = v[0]
	}

	body, err := io.ReadAll(req.data.Body)
	utils.HandleError(err)

	req.Body = &Body{byt: body}
	req.Query = req.data.URL.Query()
}

// Get a cookie from the browser.
func (req *Request) GetCookie(name string) string {
	cookie, err := req.data.Cookie(name)
	if err != nil {
		return ""
	}

	return cookie.Value
}
