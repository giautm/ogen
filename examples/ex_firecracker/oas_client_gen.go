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

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
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

// CreateSnapshot invokes createSnapshot operation.
//
// Creates a snapshot of the microVM state. The microVM should be in the `Paused` state.
//
// PUT /snapshot/create
func (c *Client) CreateSnapshot(ctx context.Context, request SnapshotCreateParams) (res CreateSnapshotRes, err error) {
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
		otelogen.OperationID("createSnapshot"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateSnapshot",
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
	u.Path += "/snapshot/create"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodeCreateSnapshotRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCreateSnapshotResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateSyncAction invokes createSyncAction operation.
//
// Creates a synchronous action.
//
// PUT /actions
func (c *Client) CreateSyncAction(ctx context.Context, request InstanceActionInfo) (res CreateSyncActionRes, err error) {
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
		otelogen.OperationID("createSyncAction"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateSyncAction",
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
	u.Path += "/actions"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodeCreateSyncActionRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCreateSyncActionResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DescribeBalloonConfig invokes describeBalloonConfig operation.
//
// Returns the current balloon device configuration.
//
// GET /balloon
func (c *Client) DescribeBalloonConfig(ctx context.Context) (res DescribeBalloonConfigRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("describeBalloonConfig"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "DescribeBalloonConfig",
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
	u.Path += "/balloon"

	r := ht.NewRequest(ctx, "GET", u, nil)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDescribeBalloonConfigResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DescribeBalloonStats invokes describeBalloonStats operation.
//
// Returns the latest balloon device statistics, only if enabled pre-boot.
//
// GET /balloon/statistics
func (c *Client) DescribeBalloonStats(ctx context.Context) (res DescribeBalloonStatsRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("describeBalloonStats"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "DescribeBalloonStats",
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
	u.Path += "/balloon/statistics"

	r := ht.NewRequest(ctx, "GET", u, nil)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDescribeBalloonStatsResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DescribeInstance invokes describeInstance operation.
//
// Returns general information about an instance.
//
// GET /
func (c *Client) DescribeInstance(ctx context.Context) (res DescribeInstanceRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("describeInstance"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "DescribeInstance",
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
	u.Path += "/"

	r := ht.NewRequest(ctx, "GET", u, nil)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDescribeInstanceResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetExportVmConfig invokes getExportVmConfig operation.
//
// Gets configuration for all VM resources.
//
// GET /vm/config
func (c *Client) GetExportVmConfig(ctx context.Context) (res GetExportVmConfigRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getExportVmConfig"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "GetExportVmConfig",
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
	u.Path += "/vm/config"

	r := ht.NewRequest(ctx, "GET", u, nil)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeGetExportVmConfigResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetMachineConfiguration invokes getMachineConfiguration operation.
//
// Gets the machine configuration of the VM. When called before the PUT operation, it will return the
// default values for the vCPU count (=1), memory size (=128 MiB). By default Hyperthreading is
// disabled and there is no CPU Template.
//
// GET /machine-config
func (c *Client) GetMachineConfiguration(ctx context.Context) (res GetMachineConfigurationRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getMachineConfiguration"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "GetMachineConfiguration",
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
	u.Path += "/machine-config"

	r := ht.NewRequest(ctx, "GET", u, nil)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeGetMachineConfigurationResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// LoadSnapshot invokes loadSnapshot operation.
//
// Loads the microVM state from a snapshot. Only accepted on a fresh Firecracker process (before
// configuring any resource other than the Logger and Metrics).
//
// PUT /snapshot/load
func (c *Client) LoadSnapshot(ctx context.Context, request SnapshotLoadParams) (res LoadSnapshotRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("loadSnapshot"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "LoadSnapshot",
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
	u.Path += "/snapshot/load"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodeLoadSnapshotRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeLoadSnapshotResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MmdsConfigPut invokes PUT /mmds/config operation.
//
// Creates MMDS configuration to be used by the MMDS network stack.
//
// PUT /mmds/config
func (c *Client) MmdsConfigPut(ctx context.Context, request MmdsConfig) (res MmdsConfigPutRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MmdsConfigPut",
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
	u.Path += "/mmds/config"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodeMmdsConfigPutRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMmdsConfigPutResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MmdsGet invokes GET /mmds operation.
//
// Get the MMDS data store.
//
// GET /mmds
func (c *Client) MmdsGet(ctx context.Context) (res MmdsGetRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MmdsGet",
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
	u.Path += "/mmds"

	r := ht.NewRequest(ctx, "GET", u, nil)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMmdsGetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MmdsPatch invokes PATCH /mmds operation.
//
// Updates the MMDS data store.
//
// PATCH /mmds
func (c *Client) MmdsPatch(ctx context.Context, request *MmdsPatchReq) (res MmdsPatchRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MmdsPatch",
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
	u.Path += "/mmds"

	r := ht.NewRequest(ctx, "PATCH", u, nil)
	if err := encodeMmdsPatchRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMmdsPatchResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MmdsPut invokes PUT /mmds operation.
//
// Creates a MMDS (Microvm Metadata Service) data store.
//
// PUT /mmds
func (c *Client) MmdsPut(ctx context.Context, request *MmdsPutReq) (res MmdsPutRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := c.cfg.Tracer.Start(ctx, "MmdsPut",
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
	u.Path += "/mmds"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodeMmdsPutRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeMmdsPutResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchBalloon invokes patchBalloon operation.
//
// Updates an existing balloon device, before or after machine startup. Will fail if update is not
// possible.
//
// PATCH /balloon
func (c *Client) PatchBalloon(ctx context.Context, request BalloonUpdate) (res PatchBalloonRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("patchBalloon"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PatchBalloon",
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
	u.Path += "/balloon"

	r := ht.NewRequest(ctx, "PATCH", u, nil)
	if err := encodePatchBalloonRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchBalloonResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchBalloonStatsInterval invokes patchBalloonStatsInterval operation.
//
// Updates an existing balloon device statistics interval, before or after machine startup. Will fail
// if update is not possible.
//
// PATCH /balloon/statistics
func (c *Client) PatchBalloonStatsInterval(ctx context.Context, request BalloonStatsUpdate) (res PatchBalloonStatsIntervalRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("patchBalloonStatsInterval"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PatchBalloonStatsInterval",
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
	u.Path += "/balloon/statistics"

	r := ht.NewRequest(ctx, "PATCH", u, nil)
	if err := encodePatchBalloonStatsIntervalRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchBalloonStatsIntervalResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchGuestDriveByID invokes patchGuestDriveByID operation.
//
// Updates the properties of the drive with the ID specified by drive_id path parameter. Will fail if
// update is not possible.
//
// PATCH /drives/{drive_id}
func (c *Client) PatchGuestDriveByID(ctx context.Context, request PartialDrive, params PatchGuestDriveByIDParams) (res PatchGuestDriveByIDRes, err error) {
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
		otelogen.OperationID("patchGuestDriveByID"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PatchGuestDriveByID",
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
	u.Path += "/drives/"
	{
		// Encode "drive_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "drive_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.DriveID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "PATCH", u, nil)
	if err := encodePatchGuestDriveByIDRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchGuestDriveByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchGuestNetworkInterfaceByID invokes patchGuestNetworkInterfaceByID operation.
//
// Updates the rate limiters applied to a network interface.
//
// PATCH /network-interfaces/{iface_id}
func (c *Client) PatchGuestNetworkInterfaceByID(ctx context.Context, request PartialNetworkInterface, params PatchGuestNetworkInterfaceByIDParams) (res PatchGuestNetworkInterfaceByIDRes, err error) {
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
		otelogen.OperationID("patchGuestNetworkInterfaceByID"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PatchGuestNetworkInterfaceByID",
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
	u.Path += "/network-interfaces/"
	{
		// Encode "iface_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "iface_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.IfaceID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "PATCH", u, nil)
	if err := encodePatchGuestNetworkInterfaceByIDRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchGuestNetworkInterfaceByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchMachineConfiguration invokes patchMachineConfiguration operation.
//
// Partially updates the Virtual Machine Configuration with the specified input. If any of the
// parameters has an incorrect value, the whole update fails.
//
// PATCH /machine-config
func (c *Client) PatchMachineConfiguration(ctx context.Context, request OptMachineConfiguration) (res PatchMachineConfigurationRes, err error) {
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
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("patchMachineConfiguration"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PatchMachineConfiguration",
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
	u.Path += "/machine-config"

	r := ht.NewRequest(ctx, "PATCH", u, nil)
	if err := encodePatchMachineConfigurationRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchMachineConfigurationResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PatchVm invokes patchVm operation.
//
// Sets the desired state (Paused or Resumed) for the microVM.
//
// PATCH /vm
func (c *Client) PatchVm(ctx context.Context, request VM) (res PatchVmRes, err error) {
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
		otelogen.OperationID("patchVm"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PatchVm",
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
	u.Path += "/vm"

	r := ht.NewRequest(ctx, "PATCH", u, nil)
	if err := encodePatchVmRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePatchVmResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutBalloon invokes putBalloon operation.
//
// Creates a new balloon device if one does not already exist, otherwise updates it, before machine
// startup. This will fail after machine startup. Will fail if update is not possible.
//
// PUT /balloon
func (c *Client) PutBalloon(ctx context.Context, request Balloon) (res PutBalloonRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putBalloon"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PutBalloon",
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
	u.Path += "/balloon"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodePutBalloonRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutBalloonResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutGuestBootSource invokes putGuestBootSource operation.
//
// Creates new boot source if one does not already exist, otherwise updates it. Will fail if update
// is not possible.
//
// PUT /boot-source
func (c *Client) PutGuestBootSource(ctx context.Context, request BootSource) (res PutGuestBootSourceRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putGuestBootSource"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PutGuestBootSource",
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
	u.Path += "/boot-source"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodePutGuestBootSourceRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutGuestBootSourceResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutGuestDriveByID invokes putGuestDriveByID operation.
//
// Creates new drive with ID specified by drive_id path parameter. If a drive with the specified ID
// already exists, updates its state based on new input. Will fail if update is not possible.
//
// PUT /drives/{drive_id}
func (c *Client) PutGuestDriveByID(ctx context.Context, request Drive, params PutGuestDriveByIDParams) (res PutGuestDriveByIDRes, err error) {
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
		otelogen.OperationID("putGuestDriveByID"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PutGuestDriveByID",
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
	u.Path += "/drives/"
	{
		// Encode "drive_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "drive_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.DriveID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodePutGuestDriveByIDRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutGuestDriveByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutGuestNetworkInterfaceByID invokes putGuestNetworkInterfaceByID operation.
//
// Creates new network interface with ID specified by iface_id path parameter.
//
// PUT /network-interfaces/{iface_id}
func (c *Client) PutGuestNetworkInterfaceByID(ctx context.Context, request NetworkInterface, params PutGuestNetworkInterfaceByIDParams) (res PutGuestNetworkInterfaceByIDRes, err error) {
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
		otelogen.OperationID("putGuestNetworkInterfaceByID"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PutGuestNetworkInterfaceByID",
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
	u.Path += "/network-interfaces/"
	{
		// Encode "iface_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "iface_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.IfaceID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodePutGuestNetworkInterfaceByIDRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutGuestNetworkInterfaceByIDResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutGuestVsock invokes putGuestVsock operation.
//
// The first call creates the device with the configuration specified in body. Subsequent calls will
// update the device configuration. May fail if update is not possible.
//
// PUT /vsock
func (c *Client) PutGuestVsock(ctx context.Context, request Vsock) (res PutGuestVsockRes, err error) {
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
		otelogen.OperationID("putGuestVsock"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PutGuestVsock",
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
	u.Path += "/vsock"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodePutGuestVsockRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutGuestVsockResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutLogger invokes putLogger operation.
//
// Initializes the logger by specifying a named pipe or a file for the logs output.
//
// PUT /logger
func (c *Client) PutLogger(ctx context.Context, request Logger) (res PutLoggerRes, err error) {
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
		otelogen.OperationID("putLogger"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PutLogger",
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
	u.Path += "/logger"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodePutLoggerRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutLoggerResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutMachineConfiguration invokes putMachineConfiguration operation.
//
// Updates the Virtual Machine Configuration with the specified input. Firecracker starts with
// default values for vCPU count (=1) and memory size (=128 MiB). With Hyperthreading enabled, the
// vCPU count is restricted to be 1 or an even number, otherwise there are no restrictions regarding
// the vCPU count. If any of the parameters has an incorrect value, the whole update fails.
//
// PUT /machine-config
func (c *Client) PutMachineConfiguration(ctx context.Context, request OptMachineConfiguration) (res PutMachineConfigurationRes, err error) {
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
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putMachineConfiguration"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PutMachineConfiguration",
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
	u.Path += "/machine-config"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodePutMachineConfigurationRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutMachineConfigurationResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutMetrics invokes putMetrics operation.
//
// Initializes the metrics system by specifying a named pipe or a file for the metrics output.
//
// PUT /metrics
func (c *Client) PutMetrics(ctx context.Context, request Metrics) (res PutMetricsRes, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putMetrics"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "PutMetrics",
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
	u.Path += "/metrics"

	r := ht.NewRequest(ctx, "PUT", u, nil)
	if err := encodePutMetricsRequestJSON(request, r); err != nil {
		return res, err
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodePutMetricsResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
