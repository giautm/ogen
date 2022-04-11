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

type MarketCandlesGetParams struct {
	// FIGI.
	Figi string
	// Начало временного промежутка.
	From time.Time
	// Конец временного промежутка.
	To time.Time
	// Интервал свечи.
	Interval CandleResolution
}

type MarketOrderbookGetParams struct {
	// FIGI.
	Figi string
	// Глубина стакана [1..20].
	Depth int32
}

type MarketSearchByFigiGetParams struct {
	// FIGI.
	Figi string
}

type MarketSearchByTickerGetParams struct {
	// Тикер инструмента.
	Ticker string
}

type OperationsGetParams struct {
	// Начало временного промежутка.
	From time.Time
	// Конец временного промежутка.
	To time.Time
	// Figi инструмента для фильтрации.
	Figi OptString
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type OrdersCancelPostParams struct {
	// ID заявки.
	OrderId string
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type OrdersGetParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type OrdersLimitOrderPostParams struct {
	// FIGI инструмента.
	Figi string
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type OrdersMarketOrderPostParams struct {
	// FIGI инструмента.
	Figi string
	// Уникальный идентификатор счета (по умолчанию -
	// Тинькофф).
	BrokerAccountId OptString
}

type PortfolioCurrenciesGetParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type PortfolioGetParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type SandboxClearPostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type SandboxCurrenciesBalancePostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type SandboxPositionsBalancePostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}

type SandboxRemovePostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
}
