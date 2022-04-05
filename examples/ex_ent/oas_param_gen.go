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

type CreatePetCategoriesParams struct {
	// ID of the Pet.
	ID int
}

type CreatePetFriendsParams struct {
	// ID of the Pet.
	ID int
}

type CreatePetOwnerParams struct {
	// ID of the Pet.
	ID int
}

type DeletePetParams struct {
	// ID of the Pet.
	ID int
}

type DeletePetOwnerParams struct {
	// ID of the Pet.
	ID int
}

type ListPetParams struct {
	// What page to render.
	Page OptInt32
	// Item count to render per page.
	ItemsPerPage OptInt32
}

type ListPetCategoriesParams struct {
	// ID of the Pet.
	ID int
	// What page to render.
	Page OptInt32
	// Item count to render per page.
	ItemsPerPage OptInt32
}

type ListPetFriendsParams struct {
	// ID of the Pet.
	ID int
	// What page to render.
	Page OptInt32
	// Item count to render per page.
	ItemsPerPage OptInt32
}

type ReadPetParams struct {
	// ID of the Pet.
	ID int
}

type ReadPetOwnerParams struct {
	// ID of the Pet.
	ID int
}

type UpdatePetParams struct {
	// ID of the Pet.
	ID int
}
