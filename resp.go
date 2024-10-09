package expressgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ContentGamer/expressgo/utils"
)

type CookieData struct {
	MaxAge   int
	HttpOnly bool
	Secure   bool
	SameSite http.SameSite
}

type JSONData map[string]interface{}
type Response struct {
	writer http.ResponseWriter

	headers map[string]string
	body    strings.Builder
	status  int
}

// Send HTML Data.
func (resp *Response) SendHTML(code string) {
	resp.SetHeader("Content-Type", "text/html")
	resp.Text(code)
}

// Redirect to another URL.
func (resp *Response) Redirect(path string) {
	resp.Status(http.StatusMovedPermanently)
	resp.SetHeader("Location", path)
}

// Send a file. (WARNING: uses os.ReadFile)
func (resp *Response) SendFile(filePath string) {
	data, err := os.ReadFile(filePath)
	utils.HandleError(err)

	if strings.HasSuffix(filePath, ".json") {
		resp.SetHeader("Content-Type", "application/json")
	}
	if strings.HasPrefix(filePath, ".html") {
		resp.SetHeader("Content-Type", "text/html")
	}

	resp.Text(string(data))
}

// Set a new header, for example: 'Content-Type'.
func (resp *Response) SetHeader(header string, data string) {
	resp.headers[header] = data
}

// Send JSON Data.
func (resp *Response) Json(data JSONData) {
	resp.SetHeader("Content-Type", "application/json")
	jdata, err := json.Marshal(data)
	utils.HandleError(err)

	resp.body.WriteString(string(jdata))
}

// Send Text Data.
func (resp *Response) Text(data interface{}) {
	str := fmt.Sprintf("%v", data)
	resp.body.WriteString(str)
}

// Send a status (recommended to use http.StatusOK for example).
func (resp *Response) Status(st int) *Response {
	resp.status = st
	return resp
}

// Send a text version of the status
// (Will also send the status at the same time).
func (resp *Response) SendStatus(st int) {
	status := http.StatusText(st)

	resp.Status(st)
	resp.Text(fmt.Sprintf("%d %s", st, status))
}

// Create a new cookie or override one.
func (resp *Response) SetCookie(name string, value string, data *CookieData) {
	cookie := &http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
	}

	if data != nil {
		cookie.MaxAge = data.MaxAge
		cookie.HttpOnly = data.HttpOnly
		cookie.Secure = data.Secure
		cookie.SameSite = data.SameSite
	}

	http.SetCookie(resp.writer, cookie)
}
