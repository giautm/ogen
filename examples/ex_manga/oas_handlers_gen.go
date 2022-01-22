// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

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
	_ = bits.LeadingZeros64
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

// HandleGetBookRequest handles getBook operation.
//
// GET /api/gallery/{book_id}
func (s *Server) handleGetBookRequest(args map[string]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `GetBook`,
		trace.WithAttributes(otelogen.OperationID(`getBook`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeGetBookParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.GetBook(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeGetBookResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleGetPageCoverImageRequest handles getPageCoverImage operation.
//
// GET /galleries/{media_id}/cover.{format}
func (s *Server) handleGetPageCoverImageRequest(args map[string]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `GetPageCoverImage`,
		trace.WithAttributes(otelogen.OperationID(`getPageCoverImage`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeGetPageCoverImageParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.GetPageCoverImage(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeGetPageCoverImageResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleGetPageImageRequest handles getPageImage operation.
//
// GET /galleries/{media_id}/{page}.{format}
func (s *Server) handleGetPageImageRequest(args map[string]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `GetPageImage`,
		trace.WithAttributes(otelogen.OperationID(`getPageImage`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeGetPageImageParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.GetPageImage(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeGetPageImageResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleGetPageThumbnailImageRequest handles getPageThumbnailImage operation.
//
// GET /galleries/{media_id}/{page}t.{format}
func (s *Server) handleGetPageThumbnailImageRequest(args map[string]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `GetPageThumbnailImage`,
		trace.WithAttributes(otelogen.OperationID(`getPageThumbnailImage`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeGetPageThumbnailImageParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.GetPageThumbnailImage(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeGetPageThumbnailImageResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleSearchRequest handles search operation.
//
// GET /api/galleries/search
func (s *Server) handleSearchRequest(args map[string]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `Search`,
		trace.WithAttributes(otelogen.OperationID(`search`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeSearchParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.Search(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeSearchResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleSearchByTagIDRequest handles searchByTagID operation.
//
// GET /api/galleries/tagged
func (s *Server) handleSearchByTagIDRequest(args map[string]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `SearchByTagID`,
		trace.WithAttributes(otelogen.OperationID(`searchByTagID`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeSearchByTagIDParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.SearchByTagID(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeSearchByTagIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}
