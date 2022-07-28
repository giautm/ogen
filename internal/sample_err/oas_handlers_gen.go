// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// HandleDataCreateRequest handles dataCreate operation.
//
// POST /data
func (s *Server) handleDataCreateRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("dataCreate"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DataCreate",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "DataCreate",
			ID:   "dataCreate",
		}
	)
	request, close, err := s.decodeDataCreateRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer close()

	response, err := s.h.DataCreate(ctx, request)
	if err != nil {
		recordError("Internal", err)
		var errRes *ErrorStatusCode
		if errors.As(err, &errRes) {
			encodeErrorResponse(*errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeDataCreateResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleDataGetRequest handles dataGet operation.
//
// GET /data
func (s *Server) handleDataGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("dataGet"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DataGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err error
	)

	response, err := s.h.DataGet(ctx)
	if err != nil {
		recordError("Internal", err)
		var errRes *ErrorStatusCode
		if errors.As(err, &errRes) {
			encodeErrorResponse(*errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeDataGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
