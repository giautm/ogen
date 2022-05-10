// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	cfg       config
	requests  syncint64.Counter
	errors    syncint64.Counter
	duration  syncint64.Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		sec:       sec,
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

// MarketBondsGet invokes  operation.
//
// GET /market/bonds
func (c *Client) MarketBondsGet(ctx context.Context) (res MarketBondsGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketBondsGet",
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
	u.Path += "/market/bonds"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "MarketBondsGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketBondsGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketCandlesGet invokes  operation.
//
// GET /market/candles
func (c *Client) MarketCandlesGet(ctx context.Context, params MarketCandlesGetParams) (res MarketCandlesGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketCandlesGet",
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
	u.Path += "/market/candles"

	q := uri.NewQueryEncoder()
	{
		// Encode "figi" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Figi))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.TimeToString(params.From))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "to" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.TimeToString(params.To))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "interval" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "interval",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(params.Interval)))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "MarketCandlesGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketCandlesGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketCurrenciesGet invokes  operation.
//
// GET /market/currencies
func (c *Client) MarketCurrenciesGet(ctx context.Context) (res MarketCurrenciesGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketCurrenciesGet",
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
	u.Path += "/market/currencies"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "MarketCurrenciesGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketCurrenciesGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketEtfsGet invokes  operation.
//
// GET /market/etfs
func (c *Client) MarketEtfsGet(ctx context.Context) (res MarketEtfsGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketEtfsGet",
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
	u.Path += "/market/etfs"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "MarketEtfsGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketEtfsGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketOrderbookGet invokes  operation.
//
// GET /market/orderbook
func (c *Client) MarketOrderbookGet(ctx context.Context, params MarketOrderbookGetParams) (res MarketOrderbookGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketOrderbookGet",
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
	u.Path += "/market/orderbook"

	q := uri.NewQueryEncoder()
	{
		// Encode "figi" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Figi))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "depth" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "depth",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.Int32ToString(params.Depth))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "MarketOrderbookGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketOrderbookGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketSearchByFigiGet invokes  operation.
//
// GET /market/search/by-figi
func (c *Client) MarketSearchByFigiGet(ctx context.Context, params MarketSearchByFigiGetParams) (res MarketSearchByFigiGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketSearchByFigiGet",
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
	u.Path += "/market/search/by-figi"

	q := uri.NewQueryEncoder()
	{
		// Encode "figi" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Figi))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "MarketSearchByFigiGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketSearchByFigiGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketSearchByTickerGet invokes  operation.
//
// GET /market/search/by-ticker
func (c *Client) MarketSearchByTickerGet(ctx context.Context, params MarketSearchByTickerGetParams) (res MarketSearchByTickerGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketSearchByTickerGet",
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
	u.Path += "/market/search/by-ticker"

	q := uri.NewQueryEncoder()
	{
		// Encode "ticker" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "ticker",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Ticker))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "MarketSearchByTickerGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketSearchByTickerGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MarketStocksGet invokes  operation.
//
// GET /market/stocks
func (c *Client) MarketStocksGet(ctx context.Context) (res MarketStocksGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MarketStocksGet",
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
	u.Path += "/market/stocks"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "MarketStocksGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMarketStocksGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OperationsGet invokes  operation.
//
// GET /operations
func (c *Client) OperationsGet(ctx context.Context, params OperationsGetParams) (res OperationsGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "OperationsGet",
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
	u.Path += "/operations"

	q := uri.NewQueryEncoder()
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.TimeToString(params.From))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "to" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.TimeToString(params.To))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "figi" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Figi.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "OperationsGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOperationsGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OrdersCancelPost invokes  operation.
//
// POST /orders/cancel
func (c *Client) OrdersCancelPost(ctx context.Context, params OrdersCancelPostParams) (res OrdersCancelPostRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "OrdersCancelPost",
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
	u.Path += "/orders/cancel"

	q := uri.NewQueryEncoder()
	{
		// Encode "orderId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "orderId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.OrderId))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "OrdersCancelPost", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOrdersCancelPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OrdersGet invokes  operation.
//
// GET /orders
func (c *Client) OrdersGet(ctx context.Context, params OrdersGetParams) (res OrdersGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "OrdersGet",
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
	u.Path += "/orders"

	q := uri.NewQueryEncoder()
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "OrdersGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOrdersGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OrdersLimitOrderPost invokes  operation.
//
// POST /orders/limit-order
func (c *Client) OrdersLimitOrderPost(ctx context.Context, request LimitOrderRequest, params OrdersLimitOrderPostParams) (res OrdersLimitOrderPostRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "OrdersLimitOrderPost",
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeOrdersLimitOrderPostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutEncoder(buf)
	reqBody = bytes.NewReader(buf.Bytes())

	u := uri.Clone(c.serverURL)
	u.Path += "/orders/limit-order"

	q := uri.NewQueryEncoder()
	{
		// Encode "figi" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Figi))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	if err := c.securitySSOAuth(ctx, "OrdersLimitOrderPost", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOrdersLimitOrderPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OrdersMarketOrderPost invokes  operation.
//
// POST /orders/market-order
func (c *Client) OrdersMarketOrderPost(ctx context.Context, request MarketOrderRequest, params OrdersMarketOrderPostParams) (res OrdersMarketOrderPostRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "OrdersMarketOrderPost",
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeOrdersMarketOrderPostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutEncoder(buf)
	reqBody = bytes.NewReader(buf.Bytes())

	u := uri.Clone(c.serverURL)
	u.Path += "/orders/market-order"

	q := uri.NewQueryEncoder()
	{
		// Encode "figi" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Figi))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	if err := c.securitySSOAuth(ctx, "OrdersMarketOrderPost", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeOrdersMarketOrderPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PortfolioCurrenciesGet invokes  operation.
//
// GET /portfolio/currencies
func (c *Client) PortfolioCurrenciesGet(ctx context.Context, params PortfolioCurrenciesGetParams) (res PortfolioCurrenciesGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "PortfolioCurrenciesGet",
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
	u.Path += "/portfolio/currencies"

	q := uri.NewQueryEncoder()
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "PortfolioCurrenciesGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePortfolioCurrenciesGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PortfolioGet invokes  operation.
//
// GET /portfolio
func (c *Client) PortfolioGet(ctx context.Context, params PortfolioGetParams) (res PortfolioGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "PortfolioGet",
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
	u.Path += "/portfolio"

	q := uri.NewQueryEncoder()
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "PortfolioGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePortfolioGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SandboxClearPost invokes  operation.
//
// Удаление всех позиций клиента.
//
// POST /sandbox/clear
func (c *Client) SandboxClearPost(ctx context.Context, params SandboxClearPostParams) (res SandboxClearPostRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "SandboxClearPost",
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
	u.Path += "/sandbox/clear"

	q := uri.NewQueryEncoder()
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "SandboxClearPost", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSandboxClearPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SandboxCurrenciesBalancePost invokes  operation.
//
// POST /sandbox/currencies/balance
func (c *Client) SandboxCurrenciesBalancePost(ctx context.Context, request SandboxSetCurrencyBalanceRequest, params SandboxCurrenciesBalancePostParams) (res SandboxCurrenciesBalancePostRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "SandboxCurrenciesBalancePost",
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeSandboxCurrenciesBalancePostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutEncoder(buf)
	reqBody = bytes.NewReader(buf.Bytes())

	u := uri.Clone(c.serverURL)
	u.Path += "/sandbox/currencies/balance"

	q := uri.NewQueryEncoder()
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	if err := c.securitySSOAuth(ctx, "SandboxCurrenciesBalancePost", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSandboxCurrenciesBalancePostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SandboxPositionsBalancePost invokes  operation.
//
// POST /sandbox/positions/balance
func (c *Client) SandboxPositionsBalancePost(ctx context.Context, request SandboxSetPositionBalanceRequest, params SandboxPositionsBalancePostParams) (res SandboxPositionsBalancePostRes, err error) {
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "SandboxPositionsBalancePost",
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeSandboxPositionsBalancePostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutEncoder(buf)
	reqBody = bytes.NewReader(buf.Bytes())

	u := uri.Clone(c.serverURL)
	u.Path += "/sandbox/positions/balance"

	q := uri.NewQueryEncoder()
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	if err := c.securitySSOAuth(ctx, "SandboxPositionsBalancePost", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSandboxPositionsBalancePostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SandboxRegisterPost invokes  operation.
//
// Создание счета и валютных позиций для клиента.
//
// POST /sandbox/register
func (c *Client) SandboxRegisterPost(ctx context.Context, request OptSandboxRegisterRequest) (res SandboxRegisterPostRes, err error) {
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
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "SandboxRegisterPost",
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
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeSandboxRegisterPostRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutEncoder(buf)
	reqBody = bytes.NewReader(buf.Bytes())

	u := uri.Clone(c.serverURL)
	u.Path += "/sandbox/register"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	if err := c.securitySSOAuth(ctx, "SandboxRegisterPost", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSandboxRegisterPostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SandboxRemovePost invokes  operation.
//
// Удаление счета клиента.
//
// POST /sandbox/remove
func (c *Client) SandboxRemovePost(ctx context.Context, params SandboxRemovePostParams) (res SandboxRemovePostRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "SandboxRemovePost",
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
	u.Path += "/sandbox/remove"

	q := uri.NewQueryEncoder()
	{
		// Encode "brokerAccountId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.BrokerAccountId.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	r := ht.NewRequest(ctx, "POST", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "SandboxRemovePost", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSandboxRemovePostResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UserAccountsGet invokes  operation.
//
// GET /user/accounts
func (c *Client) UserAccountsGet(ctx context.Context) (res UserAccountsGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "UserAccountsGet",
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
	u.Path += "/user/accounts"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	if err := c.securitySSOAuth(ctx, "UserAccountsGet", r); err != nil {
		return res, errors.Wrap(err, "security")
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeUserAccountsGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
