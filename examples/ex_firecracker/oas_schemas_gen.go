// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/netip"
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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/nonrecording"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = bytes.NewReader
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = io.Copy
	_ = math.Mod
	_ = big.Rat{}
	_ = bits.LeadingZeros64
	_ = net.IP{}
	_ = http.MethodGet
	_ = netip.Addr{}
	_ = url.URL{}
	_ = regexp.MustCompile
	_ = sort.Ints
	_ = strconv.ParseInt
	_ = strings.Builder{}
	_ = sync.Pool{}
	_ = time.Time{}

	_ = errors.Is
	_ = jx.Null
	_ = uuid.UUID{}
	_ = otel.GetTracerProvider
	_ = attribute.KeyValue{}
	_ = codes.Unset
	_ = metric.MeterConfig{}
	_ = syncint64.Counter(nil)
	_ = nonrecording.NewNoopMeterProvider
	_ = trace.TraceIDFromHex

	_ = conv.ToInt32
	_ = ht.NewRequest
	_ = json.Marshal
	_ = ogenerrors.SecurityError{}
	_ = otelogen.Version
	_ = uri.PathEncoder{}
	_ = validate.Int{}
)

// Balloon device descriptor.
// Ref: #/components/schemas/Balloon
type Balloon struct {
	// Target balloon size in MiB.
	AmountMib int "json:\"amount_mib\""
	// Whether the balloon should deflate when the guest has memory pressure.
	DeflateOnOom bool "json:\"deflate_on_oom\""
	// Interval in seconds between refreshing statistics. A non-zero value will enable the statistics.
	// Defaults to 0.
	StatsPollingIntervalS OptInt "json:\"stats_polling_interval_s\""
}

func (*Balloon) describeBalloonConfigRes() {}

// Describes the balloon device statistics.
// Ref: #/components/schemas/BalloonStats
type BalloonStats struct {
	// Target number of pages the device aims to hold.
	TargetPages int "json:\"target_pages\""
	// Actual number of pages the device is holding.
	ActualPages int "json:\"actual_pages\""
	// Target amount of memory (in MiB) the device aims to hold.
	TargetMib int "json:\"target_mib\""
	// Actual amount of memory (in MiB) the device is holding.
	ActualMib int "json:\"actual_mib\""
	// The amount of memory that has been swapped in (in bytes).
	SwapIn OptInt64 "json:\"swap_in\""
	// The amount of memory that has been swapped out to disk (in bytes).
	SwapOut OptInt64 "json:\"swap_out\""
	// The number of major page faults that have occurred.
	MajorFaults OptInt64 "json:\"major_faults\""
	// The number of minor page faults that have occurred.
	MinorFaults OptInt64 "json:\"minor_faults\""
	// The amount of memory not being used for any purpose (in bytes).
	FreeMemory OptInt64 "json:\"free_memory\""
	// The total amount of memory available (in bytes).
	TotalMemory OptInt64 "json:\"total_memory\""
	// An estimate of how much memory is available (in bytes) for starting new applications, without
	// pushing the system to swap.
	AvailableMemory OptInt64 "json:\"available_memory\""
	// The amount of memory, in bytes, that can be quickly reclaimed without additional I/O. Typically
	// these pages are used for caching files from disk.
	DiskCaches OptInt64 "json:\"disk_caches\""
	// The number of successful hugetlb page allocations in the guest.
	HugetlbAllocations OptInt64 "json:\"hugetlb_allocations\""
	// The number of failed hugetlb page allocations in the guest.
	HugetlbFailures OptInt64 "json:\"hugetlb_failures\""
}

func (*BalloonStats) describeBalloonStatsRes() {}

// Update the statistics polling interval, with the first statistics update scheduled immediately.
// Statistics cannot be turned on/off after boot.
// Ref: #/components/schemas/BalloonStatsUpdate
type BalloonStatsUpdate struct {
	// Interval in seconds between refreshing statistics.
	StatsPollingIntervalS int "json:\"stats_polling_interval_s\""
}

// Balloon device descriptor.
// Ref: #/components/schemas/BalloonUpdate
type BalloonUpdate struct {
	// Target balloon size in MiB.
	AmountMib int "json:\"amount_mib\""
}

// Boot source descriptor.
// Ref: #/components/schemas/BootSource
type BootSource struct {
	// Kernel boot arguments.
	BootArgs OptString "json:\"boot_args\""
	// Host level path to the initrd image used to boot the guest.
	InitrdPath OptString "json:\"initrd_path\""
	// Host level path to the kernel image used to boot the guest.
	KernelImagePath string "json:\"kernel_image_path\""
}

// The CPU Template defines a set of flags to be disabled from the microvm so that the features
// exposed to the guest are the same as in the selected instance type.
// Ref: #/components/schemas/CpuTemplate
type CpuTemplate string

const (
	CpuTemplateC3 CpuTemplate = "C3"
	CpuTemplateT2 CpuTemplate = "T2"
)

// CreateSnapshotNoContent is response for CreateSnapshot operation.
type CreateSnapshotNoContent struct{}

func (*CreateSnapshotNoContent) createSnapshotRes() {}

// CreateSyncActionNoContent is response for CreateSyncAction operation.
type CreateSyncActionNoContent struct{}

func (*CreateSyncActionNoContent) createSyncActionRes() {}

// Ref: #/components/schemas/Drive
type Drive struct {
	DriveID string "json:\"drive_id\""
	// Represents the caching strategy for the block device.
	CacheType    OptString "json:\"cache_type\""
	IsReadOnly   bool      "json:\"is_read_only\""
	IsRootDevice bool      "json:\"is_root_device\""
	// Represents the unique id of the boot partition of this device. It is optional and it will be taken
	// into account only if the is_root_device field is true.
	Partuuid OptString "json:\"partuuid\""
	// Host level path for the guest drive.
	PathOnHost  string         "json:\"path_on_host\""
	RateLimiter OptRateLimiter "json:\"rate_limiter\""
}

// Ref: #/components/schemas/Error
type Error struct {
	// A description of the error condition.
	FaultMessage OptString "json:\"fault_message\""
}

func (*Error) createSnapshotRes()                 {}
func (*Error) createSyncActionRes()               {}
func (*Error) describeBalloonConfigRes()          {}
func (*Error) describeBalloonStatsRes()           {}
func (*Error) loadSnapshotRes()                   {}
func (*Error) mmdsConfigPutRes()                  {}
func (*Error) mmdsGetRes()                        {}
func (*Error) mmdsPatchRes()                      {}
func (*Error) mmdsPutRes()                        {}
func (*Error) patchBalloonRes()                   {}
func (*Error) patchBalloonStatsIntervalRes()      {}
func (*Error) patchGuestDriveByIDRes()            {}
func (*Error) patchGuestNetworkInterfaceByIDRes() {}
func (*Error) patchMachineConfigurationRes()      {}
func (*Error) patchVmRes()                        {}
func (*Error) putBalloonRes()                     {}
func (*Error) putGuestBootSourceRes()             {}
func (*Error) putGuestDriveByIDRes()              {}
func (*Error) putGuestNetworkInterfaceByIDRes()   {}
func (*Error) putGuestVsockRes()                  {}
func (*Error) putLoggerRes()                      {}
func (*Error) putMachineConfigurationRes()        {}
func (*Error) putMetricsRes()                     {}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

func (*ErrorStatusCode) createSnapshotRes()                 {}
func (*ErrorStatusCode) createSyncActionRes()               {}
func (*ErrorStatusCode) describeBalloonConfigRes()          {}
func (*ErrorStatusCode) describeBalloonStatsRes()           {}
func (*ErrorStatusCode) describeInstanceRes()               {}
func (*ErrorStatusCode) getExportVmConfigRes()              {}
func (*ErrorStatusCode) getMachineConfigurationRes()        {}
func (*ErrorStatusCode) loadSnapshotRes()                   {}
func (*ErrorStatusCode) mmdsConfigPutRes()                  {}
func (*ErrorStatusCode) mmdsGetRes()                        {}
func (*ErrorStatusCode) mmdsPatchRes()                      {}
func (*ErrorStatusCode) mmdsPutRes()                        {}
func (*ErrorStatusCode) patchBalloonRes()                   {}
func (*ErrorStatusCode) patchBalloonStatsIntervalRes()      {}
func (*ErrorStatusCode) patchGuestDriveByIDRes()            {}
func (*ErrorStatusCode) patchGuestNetworkInterfaceByIDRes() {}
func (*ErrorStatusCode) patchMachineConfigurationRes()      {}
func (*ErrorStatusCode) patchVmRes()                        {}
func (*ErrorStatusCode) putBalloonRes()                     {}
func (*ErrorStatusCode) putGuestBootSourceRes()             {}
func (*ErrorStatusCode) putGuestDriveByIDRes()              {}
func (*ErrorStatusCode) putGuestNetworkInterfaceByIDRes()   {}
func (*ErrorStatusCode) putGuestVsockRes()                  {}
func (*ErrorStatusCode) putLoggerRes()                      {}
func (*ErrorStatusCode) putMachineConfigurationRes()        {}
func (*ErrorStatusCode) putMetricsRes()                     {}

// Ref: #/components/schemas/FullVmConfiguration
type FullVmConfiguration struct {
	BalloonDevice OptBalloon "json:\"balloon_device\""
	// Configurations for all block devices.
	BlockDevices  []Drive                 "json:\"block_devices\""
	BootSource    OptBootSource           "json:\"boot_source\""
	Logger        OptLogger               "json:\"logger\""
	MachineConfig OptMachineConfiguration "json:\"machine_config\""
	Metrics       OptMetrics              "json:\"metrics\""
	MmdsConfig    OptMmdsConfig           "json:\"mmds_config\""
	// Configurations for all net devices.
	NetDevices  []NetworkInterface "json:\"net_devices\""
	VsockDevice OptVsock           "json:\"vsock_device\""
}

func (*FullVmConfiguration) getExportVmConfigRes() {}

// Variant wrapper containing the real action.
// Ref: #/components/schemas/InstanceActionInfo
type InstanceActionInfo struct {
	// Enumeration indicating what type of action is contained in the payload.
	ActionType InstanceActionInfoActionType "json:\"action_type\""
}

// Enumeration indicating what type of action is contained in the payload.
type InstanceActionInfoActionType string

const (
	InstanceActionInfoActionTypeFlushMetrics   InstanceActionInfoActionType = "FlushMetrics"
	InstanceActionInfoActionTypeInstanceStart  InstanceActionInfoActionType = "InstanceStart"
	InstanceActionInfoActionTypeSendCtrlAltDel InstanceActionInfoActionType = "SendCtrlAltDel"
)

// Describes MicroVM instance information.
// Ref: #/components/schemas/InstanceInfo
type InstanceInfo struct {
	// Application name.
	AppName string "json:\"app_name\""
	// MicroVM / instance ID.
	ID string "json:\"id\""
	// The current detailed state (Not started, Running, Paused) of the Firecracker instance. This value
	// is read-only for the control-plane.
	State InstanceInfoState "json:\"state\""
	// MicroVM hypervisor build version.
	VmmVersion string "json:\"vmm_version\""
}

func (*InstanceInfo) describeInstanceRes() {}

// The current detailed state (Not started, Running, Paused) of the Firecracker instance. This value
// is read-only for the control-plane.
type InstanceInfoState string

const (
	InstanceInfoStateNotStarted InstanceInfoState = "Not started"
	InstanceInfoStateRunning    InstanceInfoState = "Running"
	InstanceInfoStatePaused     InstanceInfoState = "Paused"
)

// LoadSnapshotNoContent is response for LoadSnapshot operation.
type LoadSnapshotNoContent struct{}

func (*LoadSnapshotNoContent) loadSnapshotRes() {}

// Describes the configuration option for the logging capability.
// Ref: #/components/schemas/Logger
type Logger struct {
	// Set the level. The possible values are case-insensitive.
	Level OptLoggerLevel "json:\"level\""
	// Path to the named pipe or file for the human readable log output.
	LogPath string "json:\"log_path\""
	// Whether or not to output the level in the logs.
	ShowLevel OptBool "json:\"show_level\""
	// Whether or not to include the file path and line number of the log's origin.
	ShowLogOrigin OptBool "json:\"show_log_origin\""
}

// Set the level. The possible values are case-insensitive.
type LoggerLevel string

const (
	LoggerLevelError   LoggerLevel = "Error"
	LoggerLevelWarning LoggerLevel = "Warning"
	LoggerLevelInfo    LoggerLevel = "Info"
	LoggerLevelDebug   LoggerLevel = "Debug"
)

// Describes the number of vCPUs, memory size, Hyperthreading capabilities and the CPU template.
// Ref: #/components/schemas/MachineConfiguration
type MachineConfiguration struct {
	CPUTemplate OptCpuTemplate "json:\"cpu_template\""
	// Flag for enabling/disabling Hyperthreading.
	HtEnabled bool "json:\"ht_enabled\""
	// Memory size of VM.
	MemSizeMib int "json:\"mem_size_mib\""
	// Enable dirty page tracking. If this is enabled, then incremental guest memory snapshots can be
	// created. These belong to diff snapshots, which contain, besides the microVM state, only the memory
	// dirtied since a previous snapshot. Full snapshots each contain a full copy of the guest memory.
	TrackDirtyPages OptBool "json:\"track_dirty_pages\""
	// Number of vCPUs (either 1 or an even number).
	VcpuCount int "json:\"vcpu_count\""
}

func (*MachineConfiguration) getMachineConfigurationRes() {}

// Describes the configuration option for the metrics capability.
// Ref: #/components/schemas/Metrics
type Metrics struct {
	// Path to the named pipe or file where the JSON-formatted metrics are flushed.
	MetricsPath string "json:\"metrics_path\""
}

// Defines the MMDS configuration.
// Ref: #/components/schemas/MmdsConfig
type MmdsConfig struct {
	// A valid IPv4 link-local address.
	Ipv4Address OptString "json:\"ipv4_address\""
}

// MmdsConfigPutNoContent is response for MmdsConfigPut operation.
type MmdsConfigPutNoContent struct{}

func (*MmdsConfigPutNoContent) mmdsConfigPutRes() {}

type MmdsGetOK struct{}

func (*MmdsGetOK) mmdsGetRes() {}

// MmdsPatchNoContent is response for MmdsPatch operation.
type MmdsPatchNoContent struct{}

func (*MmdsPatchNoContent) mmdsPatchRes() {}

type MmdsPatchReq struct{}

// MmdsPutNoContent is response for MmdsPut operation.
type MmdsPutNoContent struct{}

func (*MmdsPutNoContent) mmdsPutRes() {}

type MmdsPutReq struct{}

// Defines a network interface.
// Ref: #/components/schemas/NetworkInterface
type NetworkInterface struct {
	// If this field is set, the device model will reply to HTTP GET requests sent to the MMDS address
	// via this interface. In this case, both ARP requests for 169.254.169.254 and TCP segments heading
	// to the same address are intercepted by the device model, and do not reach the associated TAP
	// device.
	AllowMmdsRequests OptBool   "json:\"allow_mmds_requests\""
	GuestMAC          OptString "json:\"guest_mac\""
	// Host level path for the guest network interface.
	HostDevName   string         "json:\"host_dev_name\""
	IfaceID       string         "json:\"iface_id\""
	RxRateLimiter OptRateLimiter "json:\"rx_rate_limiter\""
	TxRateLimiter OptRateLimiter "json:\"tx_rate_limiter\""
}

// NewOptBalloon returns new OptBalloon with value set to v.
func NewOptBalloon(v Balloon) OptBalloon {
	return OptBalloon{
		Value: v,
		Set:   true,
	}
}

// OptBalloon is optional Balloon.
type OptBalloon struct {
	Value Balloon
	Set   bool
}

// IsSet returns true if OptBalloon was set.
func (o OptBalloon) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBalloon) Reset() {
	var v Balloon
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBalloon) SetTo(v Balloon) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBalloon) Get() (v Balloon, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBalloon) Or(d Balloon) Balloon {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptBootSource returns new OptBootSource with value set to v.
func NewOptBootSource(v BootSource) OptBootSource {
	return OptBootSource{
		Value: v,
		Set:   true,
	}
}

// OptBootSource is optional BootSource.
type OptBootSource struct {
	Value BootSource
	Set   bool
}

// IsSet returns true if OptBootSource was set.
func (o OptBootSource) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBootSource) Reset() {
	var v BootSource
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBootSource) SetTo(v BootSource) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBootSource) Get() (v BootSource, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBootSource) Or(d BootSource) BootSource {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCpuTemplate returns new OptCpuTemplate with value set to v.
func NewOptCpuTemplate(v CpuTemplate) OptCpuTemplate {
	return OptCpuTemplate{
		Value: v,
		Set:   true,
	}
}

// OptCpuTemplate is optional CpuTemplate.
type OptCpuTemplate struct {
	Value CpuTemplate
	Set   bool
}

// IsSet returns true if OptCpuTemplate was set.
func (o OptCpuTemplate) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCpuTemplate) Reset() {
	var v CpuTemplate
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCpuTemplate) SetTo(v CpuTemplate) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCpuTemplate) Get() (v CpuTemplate, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCpuTemplate) Or(d CpuTemplate) CpuTemplate {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt64 returns new OptInt64 with value set to v.
func NewOptInt64(v int64) OptInt64 {
	return OptInt64{
		Value: v,
		Set:   true,
	}
}

// OptInt64 is optional int64.
type OptInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if OptInt64 was set.
func (o OptInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptLogger returns new OptLogger with value set to v.
func NewOptLogger(v Logger) OptLogger {
	return OptLogger{
		Value: v,
		Set:   true,
	}
}

// OptLogger is optional Logger.
type OptLogger struct {
	Value Logger
	Set   bool
}

// IsSet returns true if OptLogger was set.
func (o OptLogger) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptLogger) Reset() {
	var v Logger
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptLogger) SetTo(v Logger) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptLogger) Get() (v Logger, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptLogger) Or(d Logger) Logger {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptLoggerLevel returns new OptLoggerLevel with value set to v.
func NewOptLoggerLevel(v LoggerLevel) OptLoggerLevel {
	return OptLoggerLevel{
		Value: v,
		Set:   true,
	}
}

// OptLoggerLevel is optional LoggerLevel.
type OptLoggerLevel struct {
	Value LoggerLevel
	Set   bool
}

// IsSet returns true if OptLoggerLevel was set.
func (o OptLoggerLevel) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptLoggerLevel) Reset() {
	var v LoggerLevel
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptLoggerLevel) SetTo(v LoggerLevel) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptLoggerLevel) Get() (v LoggerLevel, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptLoggerLevel) Or(d LoggerLevel) LoggerLevel {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMachineConfiguration returns new OptMachineConfiguration with value set to v.
func NewOptMachineConfiguration(v MachineConfiguration) OptMachineConfiguration {
	return OptMachineConfiguration{
		Value: v,
		Set:   true,
	}
}

// OptMachineConfiguration is optional MachineConfiguration.
type OptMachineConfiguration struct {
	Value MachineConfiguration
	Set   bool
}

// IsSet returns true if OptMachineConfiguration was set.
func (o OptMachineConfiguration) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMachineConfiguration) Reset() {
	var v MachineConfiguration
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMachineConfiguration) SetTo(v MachineConfiguration) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMachineConfiguration) Get() (v MachineConfiguration, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMachineConfiguration) Or(d MachineConfiguration) MachineConfiguration {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMetrics returns new OptMetrics with value set to v.
func NewOptMetrics(v Metrics) OptMetrics {
	return OptMetrics{
		Value: v,
		Set:   true,
	}
}

// OptMetrics is optional Metrics.
type OptMetrics struct {
	Value Metrics
	Set   bool
}

// IsSet returns true if OptMetrics was set.
func (o OptMetrics) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMetrics) Reset() {
	var v Metrics
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMetrics) SetTo(v Metrics) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMetrics) Get() (v Metrics, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMetrics) Or(d Metrics) Metrics {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMmdsConfig returns new OptMmdsConfig with value set to v.
func NewOptMmdsConfig(v MmdsConfig) OptMmdsConfig {
	return OptMmdsConfig{
		Value: v,
		Set:   true,
	}
}

// OptMmdsConfig is optional MmdsConfig.
type OptMmdsConfig struct {
	Value MmdsConfig
	Set   bool
}

// IsSet returns true if OptMmdsConfig was set.
func (o OptMmdsConfig) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMmdsConfig) Reset() {
	var v MmdsConfig
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMmdsConfig) SetTo(v MmdsConfig) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMmdsConfig) Get() (v MmdsConfig, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMmdsConfig) Or(d MmdsConfig) MmdsConfig {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRateLimiter returns new OptRateLimiter with value set to v.
func NewOptRateLimiter(v RateLimiter) OptRateLimiter {
	return OptRateLimiter{
		Value: v,
		Set:   true,
	}
}

// OptRateLimiter is optional RateLimiter.
type OptRateLimiter struct {
	Value RateLimiter
	Set   bool
}

// IsSet returns true if OptRateLimiter was set.
func (o OptRateLimiter) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRateLimiter) Reset() {
	var v RateLimiter
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRateLimiter) SetTo(v RateLimiter) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRateLimiter) Get() (v RateLimiter, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRateLimiter) Or(d RateLimiter) RateLimiter {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSnapshotCreateParamsSnapshotType returns new OptSnapshotCreateParamsSnapshotType with value set to v.
func NewOptSnapshotCreateParamsSnapshotType(v SnapshotCreateParamsSnapshotType) OptSnapshotCreateParamsSnapshotType {
	return OptSnapshotCreateParamsSnapshotType{
		Value: v,
		Set:   true,
	}
}

// OptSnapshotCreateParamsSnapshotType is optional SnapshotCreateParamsSnapshotType.
type OptSnapshotCreateParamsSnapshotType struct {
	Value SnapshotCreateParamsSnapshotType
	Set   bool
}

// IsSet returns true if OptSnapshotCreateParamsSnapshotType was set.
func (o OptSnapshotCreateParamsSnapshotType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSnapshotCreateParamsSnapshotType) Reset() {
	var v SnapshotCreateParamsSnapshotType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSnapshotCreateParamsSnapshotType) SetTo(v SnapshotCreateParamsSnapshotType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSnapshotCreateParamsSnapshotType) Get() (v SnapshotCreateParamsSnapshotType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSnapshotCreateParamsSnapshotType) Or(d SnapshotCreateParamsSnapshotType) SnapshotCreateParamsSnapshotType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTokenBucket returns new OptTokenBucket with value set to v.
func NewOptTokenBucket(v TokenBucket) OptTokenBucket {
	return OptTokenBucket{
		Value: v,
		Set:   true,
	}
}

// OptTokenBucket is optional TokenBucket.
type OptTokenBucket struct {
	Value TokenBucket
	Set   bool
}

// IsSet returns true if OptTokenBucket was set.
func (o OptTokenBucket) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTokenBucket) Reset() {
	var v TokenBucket
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTokenBucket) SetTo(v TokenBucket) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTokenBucket) Get() (v TokenBucket, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTokenBucket) Or(d TokenBucket) TokenBucket {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptVsock returns new OptVsock with value set to v.
func NewOptVsock(v Vsock) OptVsock {
	return OptVsock{
		Value: v,
		Set:   true,
	}
}

// OptVsock is optional Vsock.
type OptVsock struct {
	Value Vsock
	Set   bool
}

// IsSet returns true if OptVsock was set.
func (o OptVsock) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptVsock) Reset() {
	var v Vsock
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptVsock) SetTo(v Vsock) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptVsock) Get() (v Vsock, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptVsock) Or(d Vsock) Vsock {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/PartialDrive
type PartialDrive struct {
	DriveID string "json:\"drive_id\""
	// Host level path for the guest drive.
	PathOnHost  OptString      "json:\"path_on_host\""
	RateLimiter OptRateLimiter "json:\"rate_limiter\""
}

// Defines a partial network interface structure, used to update the rate limiters for that interface,
//  after microvm start.
// Ref: #/components/schemas/PartialNetworkInterface
type PartialNetworkInterface struct {
	IfaceID       string         "json:\"iface_id\""
	RxRateLimiter OptRateLimiter "json:\"rx_rate_limiter\""
	TxRateLimiter OptRateLimiter "json:\"tx_rate_limiter\""
}

// PatchBalloonNoContent is response for PatchBalloon operation.
type PatchBalloonNoContent struct{}

func (*PatchBalloonNoContent) patchBalloonRes() {}

// PatchBalloonStatsIntervalNoContent is response for PatchBalloonStatsInterval operation.
type PatchBalloonStatsIntervalNoContent struct{}

func (*PatchBalloonStatsIntervalNoContent) patchBalloonStatsIntervalRes() {}

// PatchGuestDriveByIDNoContent is response for PatchGuestDriveByID operation.
type PatchGuestDriveByIDNoContent struct{}

func (*PatchGuestDriveByIDNoContent) patchGuestDriveByIDRes() {}

// PatchGuestNetworkInterfaceByIDNoContent is response for PatchGuestNetworkInterfaceByID operation.
type PatchGuestNetworkInterfaceByIDNoContent struct{}

func (*PatchGuestNetworkInterfaceByIDNoContent) patchGuestNetworkInterfaceByIDRes() {}

// PatchMachineConfigurationNoContent is response for PatchMachineConfiguration operation.
type PatchMachineConfigurationNoContent struct{}

func (*PatchMachineConfigurationNoContent) patchMachineConfigurationRes() {}

// PatchVmNoContent is response for PatchVm operation.
type PatchVmNoContent struct{}

func (*PatchVmNoContent) patchVmRes() {}

// PutBalloonNoContent is response for PutBalloon operation.
type PutBalloonNoContent struct{}

func (*PutBalloonNoContent) putBalloonRes() {}

// PutGuestBootSourceNoContent is response for PutGuestBootSource operation.
type PutGuestBootSourceNoContent struct{}

func (*PutGuestBootSourceNoContent) putGuestBootSourceRes() {}

// PutGuestDriveByIDNoContent is response for PutGuestDriveByID operation.
type PutGuestDriveByIDNoContent struct{}

func (*PutGuestDriveByIDNoContent) putGuestDriveByIDRes() {}

// PutGuestNetworkInterfaceByIDNoContent is response for PutGuestNetworkInterfaceByID operation.
type PutGuestNetworkInterfaceByIDNoContent struct{}

func (*PutGuestNetworkInterfaceByIDNoContent) putGuestNetworkInterfaceByIDRes() {}

// PutGuestVsockNoContent is response for PutGuestVsock operation.
type PutGuestVsockNoContent struct{}

func (*PutGuestVsockNoContent) putGuestVsockRes() {}

// PutLoggerNoContent is response for PutLogger operation.
type PutLoggerNoContent struct{}

func (*PutLoggerNoContent) putLoggerRes() {}

// PutMachineConfigurationNoContent is response for PutMachineConfiguration operation.
type PutMachineConfigurationNoContent struct{}

func (*PutMachineConfigurationNoContent) putMachineConfigurationRes() {}

// PutMetricsNoContent is response for PutMetrics operation.
type PutMetricsNoContent struct{}

func (*PutMetricsNoContent) putMetricsRes() {}

// Defines an IO rate limiter with independent bytes/s and ops/s limits. Limits are defined by
// configuring each of the _bandwidth_ and _ops_ token buckets.
// Ref: #/components/schemas/RateLimiter
type RateLimiter struct {
	Bandwidth OptTokenBucket "json:\"bandwidth\""
	Ops       OptTokenBucket "json:\"ops\""
}

// Ref: #/components/schemas/SnapshotCreateParams
type SnapshotCreateParams struct {
	// Path to the file that will contain the guest memory.
	MemFilePath string "json:\"mem_file_path\""
	// Path to the file that will contain the microVM state.
	SnapshotPath string "json:\"snapshot_path\""
	// Type of snapshot to create. It is optional and by default, a full snapshot is created.
	SnapshotType OptSnapshotCreateParamsSnapshotType "json:\"snapshot_type\""
	// The microVM version for which we want to create the snapshot. It is optional and it defaults to
	// the current version.
	Version OptString "json:\"version\""
}

// Type of snapshot to create. It is optional and by default, a full snapshot is created.
type SnapshotCreateParamsSnapshotType string

const (
	SnapshotCreateParamsSnapshotTypeFull SnapshotCreateParamsSnapshotType = "Full"
	SnapshotCreateParamsSnapshotTypeDiff SnapshotCreateParamsSnapshotType = "Diff"
)

// Ref: #/components/schemas/SnapshotLoadParams
type SnapshotLoadParams struct {
	// Enable support for incremental (diff) snapshots by tracking dirty guest pages.
	EnableDiffSnapshots OptBool "json:\"enable_diff_snapshots\""
	// Path to the file that contains the guest memory to be loaded.
	MemFilePath string "json:\"mem_file_path\""
	// Path to the file that contains the microVM state to be loaded.
	SnapshotPath string "json:\"snapshot_path\""
	// When set to true, the vm is also resumed if the snapshot load is successful.
	ResumeVM OptBool "json:\"resume_vm\""
}

// Defines a token bucket with a maximum capacity (size), an initial burst size (one_time_burst) and
// an interval for refilling purposes (refill_time). The refill-rate is derived from size and
// refill_time, and it is the constant rate at which the tokens replenish. The refill process only
// starts happening after the initial burst budget is consumed. Consumption from the token bucket is
// unbounded in speed which allows for bursts bound in size by the amount of tokens available. Once
// the token bucket is empty, consumption speed is bound by the refill_rate.
// Ref: #/components/schemas/TokenBucket
type TokenBucket struct {
	// The initial size of a token bucket.
	OneTimeBurst OptInt64 "json:\"one_time_burst\""
	// The amount of milliseconds it takes for the bucket to refill.
	RefillTime int64 "json:\"refill_time\""
	// The total number of tokens this bucket can hold.
	Size int64 "json:\"size\""
}

// Defines the microVM running state. It is especially useful in the snapshotting context.
// Ref: #/components/schemas/Vm
type VM struct {
	State VMState "json:\"state\""
}

type VMState string

const (
	VMStatePaused  VMState = "Paused"
	VMStateResumed VMState = "Resumed"
)

// Defines a vsock device, backed by a set of Unix Domain Sockets, on the host side. For
// host-initiated connections, Firecracker will be listening on the Unix socket identified by the
// path `uds_path`. Firecracker will create this socket, bind and listen on it. Host-initiated
// connections will be performed by connection to this socket and issuing a connection forwarding
// request to the desired guest-side vsock port (i.e. `CONNECT 52\n`, to connect to port 52). For
// guest-initiated connections, Firecracker will expect host software to be bound and listening on
// Unix sockets at `uds_path_<PORT>`. E.g. "/path/to/host_vsock.sock_52" for port number 52.
// Ref: #/components/schemas/Vsock
type Vsock struct {
	// Guest Vsock CID.
	GuestCid int "json:\"guest_cid\""
	// Path to UNIX domain socket, used to proxy vsock connections.
	UdsPath string "json:\"uds_path\""
	VsockID string "json:\"vsock_id\""
}
