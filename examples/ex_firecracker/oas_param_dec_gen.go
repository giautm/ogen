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
	_ = otelogen.Version
	_ = uri.PathEncoder{}
	_ = validate.Int{}
)

func decodePatchGuestDriveByIDParams(args [1]string, r *http.Request) (PatchGuestDriveByIDParams, error) {
	var (
		params PatchGuestDriveByIDParams
	)
	// Decode path: drive_id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "drive_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.DriveID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: drive_id: not specified")
		}
	}
	return params, nil
}

func decodePatchGuestNetworkInterfaceByIDParams(args [1]string, r *http.Request) (PatchGuestNetworkInterfaceByIDParams, error) {
	var (
		params PatchGuestNetworkInterfaceByIDParams
	)
	// Decode path: iface_id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "iface_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.IfaceID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: iface_id: not specified")
		}
	}
	return params, nil
}

func decodePutGuestDriveByIDParams(args [1]string, r *http.Request) (PutGuestDriveByIDParams, error) {
	var (
		params PutGuestDriveByIDParams
	)
	// Decode path: drive_id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "drive_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.DriveID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: drive_id: not specified")
		}
	}
	return params, nil
}

func decodePutGuestNetworkInterfaceByIDParams(args [1]string, r *http.Request) (PutGuestNetworkInterfaceByIDParams, error) {
	var (
		params PutGuestNetworkInterfaceByIDParams
	)
	// Decode path: iface_id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "iface_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.IfaceID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: iface_id: not specified")
		}
	}
	return params, nil
}
