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
	// CreateSnapshot implements createSnapshot operation.
	//
	// PUT /snapshot/create
	CreateSnapshot(ctx context.Context, req SnapshotCreateParams) (CreateSnapshotRes, error)
	// CreateSyncAction implements createSyncAction operation.
	//
	// PUT /actions
	CreateSyncAction(ctx context.Context, req InstanceActionInfo) (CreateSyncActionRes, error)
	// DescribeBalloonConfig implements describeBalloonConfig operation.
	//
	// GET /balloon
	DescribeBalloonConfig(ctx context.Context) (DescribeBalloonConfigRes, error)
	// DescribeBalloonStats implements describeBalloonStats operation.
	//
	// GET /balloon/statistics
	DescribeBalloonStats(ctx context.Context) (DescribeBalloonStatsRes, error)
	// DescribeInstance implements describeInstance operation.
	//
	// GET /
	DescribeInstance(ctx context.Context) (DescribeInstanceRes, error)
	// GetExportVmConfig implements getExportVmConfig operation.
	//
	// GET /vm/config
	GetExportVmConfig(ctx context.Context) (GetExportVmConfigRes, error)
	// GetMachineConfiguration implements getMachineConfiguration operation.
	//
	// GET /machine-config
	GetMachineConfiguration(ctx context.Context) (GetMachineConfigurationRes, error)
	// LoadSnapshot implements loadSnapshot operation.
	//
	// PUT /snapshot/load
	LoadSnapshot(ctx context.Context, req SnapshotLoadParams) (LoadSnapshotRes, error)
	// MmdsConfigPut implements  operation.
	//
	// PUT /mmds/config
	MmdsConfigPut(ctx context.Context, req MmdsConfig) (MmdsConfigPutRes, error)
	// MmdsGet implements  operation.
	//
	// GET /mmds
	MmdsGet(ctx context.Context) (MmdsGetRes, error)
	// MmdsPatch implements  operation.
	//
	// PATCH /mmds
	MmdsPatch(ctx context.Context, req MmdsPatchReq) (MmdsPatchRes, error)
	// MmdsPut implements  operation.
	//
	// PUT /mmds
	MmdsPut(ctx context.Context, req MmdsPutReq) (MmdsPutRes, error)
	// PatchBalloon implements patchBalloon operation.
	//
	// PATCH /balloon
	PatchBalloon(ctx context.Context, req BalloonUpdate) (PatchBalloonRes, error)
	// PatchBalloonStatsInterval implements patchBalloonStatsInterval operation.
	//
	// PATCH /balloon/statistics
	PatchBalloonStatsInterval(ctx context.Context, req BalloonStatsUpdate) (PatchBalloonStatsIntervalRes, error)
	// PatchGuestDriveByID implements patchGuestDriveByID operation.
	//
	// PATCH /drives/{drive_id}
	PatchGuestDriveByID(ctx context.Context, req PartialDrive, params PatchGuestDriveByIDParams) (PatchGuestDriveByIDRes, error)
	// PatchGuestNetworkInterfaceByID implements patchGuestNetworkInterfaceByID operation.
	//
	// PATCH /network-interfaces/{iface_id}
	PatchGuestNetworkInterfaceByID(ctx context.Context, req PartialNetworkInterface, params PatchGuestNetworkInterfaceByIDParams) (PatchGuestNetworkInterfaceByIDRes, error)
	// PatchMachineConfiguration implements patchMachineConfiguration operation.
	//
	// PATCH /machine-config
	PatchMachineConfiguration(ctx context.Context, req MachineConfiguration) (PatchMachineConfigurationRes, error)
	// PatchVm implements patchVm operation.
	//
	// PATCH /vm
	PatchVm(ctx context.Context, req VM) (PatchVmRes, error)
	// PutBalloon implements putBalloon operation.
	//
	// PUT /balloon
	PutBalloon(ctx context.Context, req Balloon) (PutBalloonRes, error)
	// PutGuestBootSource implements putGuestBootSource operation.
	//
	// PUT /boot-source
	PutGuestBootSource(ctx context.Context, req BootSource) (PutGuestBootSourceRes, error)
	// PutGuestDriveByID implements putGuestDriveByID operation.
	//
	// PUT /drives/{drive_id}
	PutGuestDriveByID(ctx context.Context, req Drive, params PutGuestDriveByIDParams) (PutGuestDriveByIDRes, error)
	// PutGuestNetworkInterfaceByID implements putGuestNetworkInterfaceByID operation.
	//
	// PUT /network-interfaces/{iface_id}
	PutGuestNetworkInterfaceByID(ctx context.Context, req NetworkInterface, params PutGuestNetworkInterfaceByIDParams) (PutGuestNetworkInterfaceByIDRes, error)
	// PutGuestVsock implements putGuestVsock operation.
	//
	// PUT /vsock
	PutGuestVsock(ctx context.Context, req Vsock) (PutGuestVsockRes, error)
	// PutLogger implements putLogger operation.
	//
	// PUT /logger
	PutLogger(ctx context.Context, req Logger) (PutLoggerRes, error)
	// PutMachineConfiguration implements putMachineConfiguration operation.
	//
	// PUT /machine-config
	PutMachineConfiguration(ctx context.Context, req MachineConfiguration) (PutMachineConfigurationRes, error)
	// PutMetrics implements putMetrics operation.
	//
	// PUT /metrics
	PutMetrics(ctx context.Context, req Metrics) (PutMetricsRes, error)
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

// Register request handlers in router.
func (s *Server) Register(r chi.Router) {
	r.MethodFunc("PUT", "/snapshot/create", s.HandleCreateSnapshotRequest)
	r.MethodFunc("PUT", "/actions", s.HandleCreateSyncActionRequest)
	r.MethodFunc("GET", "/balloon", s.HandleDescribeBalloonConfigRequest)
	r.MethodFunc("GET", "/balloon/statistics", s.HandleDescribeBalloonStatsRequest)
	r.MethodFunc("GET", "/", s.HandleDescribeInstanceRequest)
	r.MethodFunc("GET", "/vm/config", s.HandleGetExportVmConfigRequest)
	r.MethodFunc("GET", "/machine-config", s.HandleGetMachineConfigurationRequest)
	r.MethodFunc("PUT", "/snapshot/load", s.HandleLoadSnapshotRequest)
	r.MethodFunc("PUT", "/mmds/config", s.HandleMmdsConfigPutRequest)
	r.MethodFunc("GET", "/mmds", s.HandleMmdsGetRequest)
	r.MethodFunc("PATCH", "/mmds", s.HandleMmdsPatchRequest)
	r.MethodFunc("PUT", "/mmds", s.HandleMmdsPutRequest)
	r.MethodFunc("PATCH", "/balloon", s.HandlePatchBalloonRequest)
	r.MethodFunc("PATCH", "/balloon/statistics", s.HandlePatchBalloonStatsIntervalRequest)
	r.MethodFunc("PATCH", "/drives/{drive_id}", s.HandlePatchGuestDriveByIDRequest)
	r.MethodFunc("PATCH", "/network-interfaces/{iface_id}", s.HandlePatchGuestNetworkInterfaceByIDRequest)
	r.MethodFunc("PATCH", "/machine-config", s.HandlePatchMachineConfigurationRequest)
	r.MethodFunc("PATCH", "/vm", s.HandlePatchVmRequest)
	r.MethodFunc("PUT", "/balloon", s.HandlePutBalloonRequest)
	r.MethodFunc("PUT", "/boot-source", s.HandlePutGuestBootSourceRequest)
	r.MethodFunc("PUT", "/drives/{drive_id}", s.HandlePutGuestDriveByIDRequest)
	r.MethodFunc("PUT", "/network-interfaces/{iface_id}", s.HandlePutGuestNetworkInterfaceByIDRequest)
	r.MethodFunc("PUT", "/vsock", s.HandlePutGuestVsockRequest)
	r.MethodFunc("PUT", "/logger", s.HandlePutLoggerRequest)
	r.MethodFunc("PUT", "/machine-config", s.HandlePutMachineConfigurationRequest)
	r.MethodFunc("PUT", "/metrics", s.HandlePutMetricsRequest)
}

// DefaultMux returns new *chi.Mux with called Register method on it.
func (s *Server) DefaultMux() *chi.Mux {
	mux := chi.NewMux()
	s.Register(mux)
	return mux
}
