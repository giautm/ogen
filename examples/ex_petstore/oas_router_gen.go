// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := []byte(r.URL.Path)
	if len(p) == 0 {
		s.notFound(w, r)
		return
	}

	var (
		idx  int               // index of next slash
		elem []byte            // current element, without slashes
		args map[string]string // lazily initialized
	)

	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		// Root edge.
		if len(p) > 1 && p[0] == '/' {
			p = p[1:]
		}
		if idx = bytes.IndexByte(p[:], '/'); idx < 0 { // looking for next element
			elem, p = p, p[:0] // slash not found, using full path
		} else {
			elem = p[:idx] // slash found, element is path until slash
			p = p[idx:]
		}
		switch string(elem) {
		case "pets": // -> 1
			// Edge: 1, path: "pets".
			if len(p) > 1 && p[0] == '/' {
				p = p[1:]
			}
			if idx = bytes.IndexByte(p[:], '/'); idx < 0 { // looking for next element
				elem, p = p, p[:0] // slash not found, using full path
			} else {
				elem = p[:idx] // slash found, element is path until slash
				p = p[idx:]
			}
			switch string(elem) {
			default:
				if args == nil {
					args = make(map[string]string)
				}
				args["petId"] = string(elem)
				// GET /pets/{petId}.
				s.handleShowPetByIdRequest(args, w, r)
				return
			}
		default:
			s.notFound(w, r)
			return
		}
	case "POST":
		// Root edge.
		if len(p) > 1 && p[0] == '/' {
			p = p[1:]
		}
		if idx = bytes.IndexByte(p[:], '/'); idx < 0 { // looking for next element
			elem, p = p, p[:0] // slash not found, using full path
		} else {
			elem = p[:idx] // slash found, element is path until slash
			p = p[idx:]
		}
		switch string(elem) {
		case "pets": // -> 1
			// POST /pets.
			s.handleCreatePetsRequest(args, w, r)
			return
		default:
			s.notFound(w, r)
			return
		}
	default:
		s.notFound(w, r)
		return
	}
}
