package rt

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type RoundTripper interface {
	http.RoundTripper
	Wrap(http.RoundTripper) RoundTripper
	Unwrap() http.RoundTripper
}

// RoundTrippers will collect many RoundTripper interfaces and return one that
// contains them all.
func RoundTrippers(base http.RoundTripper, parents ...RoundTripper) http.RoundTripper {
	prev := base
	for _, p := range parents {
		p.Wrap(prev)
		prev = p
	}
	return prev
}

// Debugger is a RoundTripper that prints debug info about http requests and
// responses.
type Debugger struct {
	http.RoundTripper
	SkipBody bool
	count    int
}

func (d *Debugger) Wrap(rt http.RoundTripper) RoundTripper {
	d.RoundTripper = rt
	return d
}

func (d *Debugger) Unwrap() http.RoundTripper { return d.RoundTripper }

func (d *Debugger) RoundTrip(req *http.Request) (*http.Response, error) {
	var (
		reqbody bytes.Buffer
		resbody bytes.Buffer
	)
	if !d.SkipBody && req.Body != nil {
		_, err := io.Copy(&reqbody, req.Body)
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewReader(reqbody.Bytes()))
	}
	res, err := d.RoundTripper.RoundTrip(req)
	if err != nil {
		return res, err
	}
	if !d.SkipBody && res.Body != nil {
		_, err := io.Copy(&resbody, res.Body)
		if err != nil {
			return res, err
		}
		if err = res.Body.Close(); err != nil {
			return res, err
		}
		res.Body = io.NopCloser(bytes.NewReader(resbody.Bytes()))
	}
	r := res.Request
	id := slog.Int("id", d.count)
	slog.Debug("START")
	slog.Info("Send", id, slog.String("method", r.Method), slog.String("url", r.URL.String()))
	for k, v := range r.Header {
		slog.Debug("request header", slog.String(k, strings.Join(v, ", ")))
	}
	if !d.SkipBody && reqbody.Len() > 0 {
		slog.Debug("request body", slog.String("body", reqbody.String()))
	}
	slog.Info("Receive", id, slog.String("status", res.Status))
	for k, v := range res.Header {
		slog.Debug("response header", slog.String(k, strings.Join(v, ", ")))
	}
	if !d.SkipBody && resbody.Len() > 0 {
		slog.Debug("response body", slog.String("body", resbody.String()))
	}
	slog.Debug("END")
	d.count += 1
	return res, nil
}

func NewUserAgent(ua string) *UserAgent {
	return &UserAgent{
		UserAgent: ua,
		parent:    nil,
	}
}

type UserAgent struct {
	UserAgent string
	parent    http.RoundTripper
}

func (uat *UserAgent) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(req.Header.Get("User-Agent")) == 0 {
		req.Header.Set("User-Agent", uat.UserAgent)
	}
	return uat.parent.RoundTrip(req)
}

func (uat *UserAgent) Wrap(rt http.RoundTripper) RoundTripper {
	uat.parent = rt
	return uat
}

func (uat *UserAgent) Unwrap() http.RoundTripper { return uat.parent }

type HttpMock struct {
	base http.RoundTripper
}

func (m *HttpMock) Wrap(rt http.RoundTripper) RoundTripper {
	switch r := rt.(type) {
	case *http.Transport:
	case RoundTripper:
		m.base = r
	}
	return m
}

func (m *HttpMock) Unwrap() http.RoundTripper { return m.base }

func (m *HttpMock) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func (m *HttpMock) Inject(client *http.Client) {
	client.Transport = m.Wrap(client.Transport)
}

func removeHttpTransport(transport http.RoundTripper) http.RoundTripper {
	var inner http.RoundTripper = transport
	for {
		switch rt := inner.(type) {
		case *http.Transport:
		case RoundTripper:
			inner = rt.Unwrap()
		default:
			return transport
		}
		// if rt, ok := inner.(RoundTripper); ok {
		// 	inner = rt.Unwrap()
		// } else {
		// 	return transport
		// }
	}
}
