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
	"go.opentelemetry.io/otel/codes"
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
	_ = big.Rat{}
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
	_ = codes.Unset
)

type GetBookParams struct {
	// ID of book.
	BookID int
}

type SearchParams struct {
	// Search query.
	// * You can search for multiple terms at the same time, and this will return only galleries that
	// contain both terms. For example, rust cox finds all galleries that contain both rust and cox.
	// * You can exclude terms by prefixing them with -. For example, rust cox -tokio matches all
	// galleries matching rust and cox but not tokio.
	// * Exact searches can be performed by wrapping terms in double quotes. For example, "big dogs" only
	// matches galleries with "big dogs" somewhere in the title or in tags.
	// * These can be combined with tag namespaces for finer control over the query: parodies:railgun
	// -tag:"big dogs".
	Query string
	// Number of result page.
	Page OptInt
}

type SearchByTagIDParams struct {
	// Tag ID.
	TagID int
	// Number of result page.
	Page OptInt
}
