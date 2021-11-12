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

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// GetBook implements getBook operation.
	//
	// GET /api/gallery/{book_id}
	GetBook(ctx context.Context, params GetBookParams) (GetBookRes, error)
	// GetPageCoverImage implements getPageCoverImage operation.
	//
	// GET /galleries/{media_id}/cover.{format}
	GetPageCoverImage(ctx context.Context, params GetPageCoverImageParams) (GetPageCoverImageRes, error)
	// GetPageImage implements getPageImage operation.
	//
	// GET /galleries/{media_id}/{page}.{format}
	GetPageImage(ctx context.Context, params GetPageImageParams) (GetPageImageRes, error)
	// GetPageThumbnailImage implements getPageThumbnailImage operation.
	//
	// GET /galleries/{media_id}/{page}t.{format}
	GetPageThumbnailImage(ctx context.Context, params GetPageThumbnailImageParams) (GetPageThumbnailImageRes, error)
	// Search implements search operation.
	//
	// GET /api/galleries/search
	Search(ctx context.Context, params SearchParams) (SearchRes, error)
	// SearchByTagID implements searchByTagID operation.
	//
	// GET /api/galleries/tagged
	SearchByTagID(ctx context.Context, params SearchByTagIDParams) (SearchByTagIDRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	cfg config
}

func NewServer(h Handler, opts ...Option) *Server {
	srv := &Server{
		h:   h,
		cfg: newConfig(opts...),
	}
	return srv
}
