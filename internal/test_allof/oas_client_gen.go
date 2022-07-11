// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	cfg       config
	requests  syncint64.Counter
	errors    syncint64.Counter
	duration  syncint64.Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		serverURL: u,
	}
	if c.requests, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.SyncInt64().Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

// NullableStrings invokes nullableStrings operation.
//
// Nullable strings.
//
// POST /nullableStrings
func (c *Client) NullableStrings(ctx context.Context, request string) (res NullableStringsOK, err error) {
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"],
		}).Validate(string(request)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("nullableStrings"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "NullableStrings",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/nullableStrings"

	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeNullableStringsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeNullableStringsResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ObjectsWithConflictingArrayProperty invokes objectsWithConflictingArrayProperty operation.
//
// Objects with conflicting array property.
//
// POST /objectsWithConflictingArrayProperty
func (c *Client) ObjectsWithConflictingArrayProperty(ctx context.Context, request ObjectsWithConflictingArrayPropertyReq) (res ObjectsWithConflictingArrayPropertyOK, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("objectsWithConflictingArrayProperty"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "ObjectsWithConflictingArrayProperty",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/objectsWithConflictingArrayProperty"

	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeObjectsWithConflictingArrayPropertyRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeObjectsWithConflictingArrayPropertyResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ObjectsWithConflictingProperties invokes objectsWithConflictingProperties operation.
//
// Objects with conflicting properties.
//
// POST /objectsWithConflictingProperties
func (c *Client) ObjectsWithConflictingProperties(ctx context.Context, request ObjectsWithConflictingPropertiesReq) (res ObjectsWithConflictingPropertiesOK, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("objectsWithConflictingProperties"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "ObjectsWithConflictingProperties",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/objectsWithConflictingProperties"

	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeObjectsWithConflictingPropertiesRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeObjectsWithConflictingPropertiesResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReferencedAllof invokes referencedAllof operation.
//
// Referenced allOf.
//
// POST /referencedAllof
func (c *Client) ReferencedAllof(ctx context.Context, request ReferencedAllofReq) (res ReferencedAllofOK, err error) {
	switch request := request.(type) {
	case *ReferencedAllofApplicationJSON:
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	case *ReferencedAllofMultipartFormData:
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("referencedAllof"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "ReferencedAllof",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/referencedAllof"

	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeReferencedAllofRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeReferencedAllofResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReferencedAllofOptional invokes referencedAllofOptional operation.
//
// Referenced allOf, but requestBody is not required.
//
// POST /referencedAllofOptional
func (c *Client) ReferencedAllofOptional(ctx context.Context, request ReferencedAllofOptionalReq) (res ReferencedAllofOptionalOK, err error) {
	switch request := request.(type) {
	case *OptRobot:
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	case *OptRobot:
		if err := func() error {
			if request.Set {
				if err := func() error {
					if err := request.Value.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "validate")
		}
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("referencedAllofOptional"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "ReferencedAllofOptional",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/referencedAllofOptional"

	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeReferencedAllofOptionalRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeReferencedAllofOptionalResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SimpleInteger invokes simpleInteger operation.
//
// Simple integers with validation.
//
// POST /simpleInteger
func (c *Client) SimpleInteger(ctx context.Context, request int) (res SimpleIntegerOK, err error) {
	if err := func() error {
		if err := (validate.Int{
			MinSet:        true,
			Min:           -5,
			MaxSet:        true,
			Max:           5,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    0,
		}).Validate(int64(request)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("simpleInteger"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "SimpleInteger",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/simpleInteger"

	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeSimpleIntegerRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSimpleIntegerResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SimpleObjects invokes simpleObjects operation.
//
// Simple objects.
//
// POST /simpleObjects
func (c *Client) SimpleObjects(ctx context.Context, request SimpleObjectsReq) (res SimpleObjectsOK, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("simpleObjects"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "SimpleObjects",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/simpleObjects"

	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeSimpleObjectsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSimpleObjectsResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
