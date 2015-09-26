package api

import (
	"encoding/json"
	l4g "github.com/macedo/category_service-go/Godeps/_workspace/src/code.google.com/p/log4go"
	"github.com/macedo/category_service-go/model"
	"net"
	"net/http"
)

type AppError struct {
	Message       string `json:"message"`
	DetailedError string `json:"detailed_error"`
	RequestId     string `json:"request_id"`
	StatusCode    int    `json:"status_code"`
	Where         string `json:"-"`
}

func (ae *AppError) ToJson() string {
	b, err := json.Marshal(ae)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

type Controller struct {
	RequestId string
	IpAddress string
	Path      string
	Protocol  string
	Err       *AppError
}

func (c *Controller) LogError(err *AppError) {
	l4g.Error("%v:%v code=%v rid=%v ip=%v %v [details: %v]", c.Path, err.Where, err.StatusCode, c.RequestId, c.IpAddress, err.Message, err.DetailedError)
}

func ApiHandler(h func(*Controller, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{h}
}

type handler struct {
	handleFunc func(*Controller, http.ResponseWriter, *http.Request)
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l4g.Debug("%v", r.URL.Path)

	c := &Controller{
		RequestId: model.NewId(),
		IpAddress: GetIpAddress(r),
		Protocol:  GetProtocol(r),
	}

	w.Header().Set("X-Request-ID", c.RequestId)
	w.Header().Set("Content-Type", "application/json")

	if c.Err == nil {
		h.handleFunc(c, w, r)
	}

	if c.Err != nil {
		c.Err.RequestId = c.RequestId
		c.LogError(c.Err)
		c.Err.Where = r.URL.Path

		w.WriteHeader(c.Err.StatusCode)
		w.Write([]byte(c.Err.ToJson()))
	}
}

func GetIpAddress(r *http.Request) string {
	address := r.Header.Get("X-Forwarded-For")

	if len(address) == 0 {
		address = r.Header.Get("X-Real-IP")
	}

	if len(address) == 0 {
		address, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	return address
}

func GetProtocol(r *http.Request) string {
	if r.Header.Get("X-Forwarded-Proto") == "https" {
		return "https"
	} else {
		return "http"
	}
}
