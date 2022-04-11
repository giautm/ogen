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

func decodeMarketCandlesGetParams(args [0]string, r *http.Request) (MarketCandlesGetParams, error) {
	var (
		params    MarketCandlesGetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: figi.
	{
		values, ok := queryArgs["figi"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
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

				params.Figi = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		} else {
			return params, errors.New("query: figi: not specified")
		}
	}
	// Decode query: from.
	{
		values, ok := queryArgs["from"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToTime(s)
				if err != nil {
					return err
				}

				params.From = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: from: parse")
			}
		} else {
			return params, errors.New("query: from: not specified")
		}
	}
	// Decode query: to.
	{
		values, ok := queryArgs["to"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToTime(s)
				if err != nil {
					return err
				}

				params.To = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: to: parse")
			}
		} else {
			return params, errors.New("query: to: not specified")
		}
	}
	// Decode query: interval.
	{
		values, ok := queryArgs["interval"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
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

				params.Interval = CandleResolution(c)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: interval: parse")
			}
			if err := func() error {
				if err := params.Interval.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: interval: invalid")
			}
		} else {
			return params, errors.New("query: interval: not specified")
		}
	}
	return params, nil
}

func decodeMarketOrderbookGetParams(args [0]string, r *http.Request) (MarketOrderbookGetParams, error) {
	var (
		params    MarketOrderbookGetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: figi.
	{
		values, ok := queryArgs["figi"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
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

				params.Figi = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		} else {
			return params, errors.New("query: figi: not specified")
		}
	}
	// Decode query: depth.
	{
		values, ok := queryArgs["depth"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt32(s)
				if err != nil {
					return err
				}

				params.Depth = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: depth: parse")
			}
		} else {
			return params, errors.New("query: depth: not specified")
		}
	}
	return params, nil
}

func decodeMarketSearchByFigiGetParams(args [0]string, r *http.Request) (MarketSearchByFigiGetParams, error) {
	var (
		params    MarketSearchByFigiGetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: figi.
	{
		values, ok := queryArgs["figi"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
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

				params.Figi = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		} else {
			return params, errors.New("query: figi: not specified")
		}
	}
	return params, nil
}

func decodeMarketSearchByTickerGetParams(args [0]string, r *http.Request) (MarketSearchByTickerGetParams, error) {
	var (
		params    MarketSearchByTickerGetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: ticker.
	{
		values, ok := queryArgs["ticker"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
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

				params.Ticker = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: ticker: parse")
			}
		} else {
			return params, errors.New("query: ticker: not specified")
		}
	}
	return params, nil
}

func decodeOperationsGetParams(args [0]string, r *http.Request) (OperationsGetParams, error) {
	var (
		params    OperationsGetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: from.
	{
		values, ok := queryArgs["from"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToTime(s)
				if err != nil {
					return err
				}

				params.From = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: from: parse")
			}
		} else {
			return params, errors.New("query: from: not specified")
		}
	}
	// Decode query: to.
	{
		values, ok := queryArgs["to"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToTime(s)
				if err != nil {
					return err
				}

				params.To = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: to: parse")
			}
		} else {
			return params, errors.New("query: to: not specified")
		}
	}
	// Decode query: figi.
	{
		values, ok := queryArgs["figi"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotFigiVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotFigiVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Figi.SetTo(paramsDotFigiVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		}
	}
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeOrdersCancelPostParams(args [0]string, r *http.Request) (OrdersCancelPostParams, error) {
	var (
		params    OrdersCancelPostParams
		queryArgs = r.URL.Query()
	)
	// Decode query: orderId.
	{
		values, ok := queryArgs["orderId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
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

				params.OrderId = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: orderId: parse")
			}
		} else {
			return params, errors.New("query: orderId: not specified")
		}
	}
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeOrdersGetParams(args [0]string, r *http.Request) (OrdersGetParams, error) {
	var (
		params    OrdersGetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeOrdersLimitOrderPostParams(args [0]string, r *http.Request) (OrdersLimitOrderPostParams, error) {
	var (
		params    OrdersLimitOrderPostParams
		queryArgs = r.URL.Query()
	)
	// Decode query: figi.
	{
		values, ok := queryArgs["figi"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
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

				params.Figi = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		} else {
			return params, errors.New("query: figi: not specified")
		}
	}
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeOrdersMarketOrderPostParams(args [0]string, r *http.Request) (OrdersMarketOrderPostParams, error) {
	var (
		params    OrdersMarketOrderPostParams
		queryArgs = r.URL.Query()
	)
	// Decode query: figi.
	{
		values, ok := queryArgs["figi"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
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

				params.Figi = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		} else {
			return params, errors.New("query: figi: not specified")
		}
	}
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodePortfolioCurrenciesGetParams(args [0]string, r *http.Request) (PortfolioCurrenciesGetParams, error) {
	var (
		params    PortfolioCurrenciesGetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodePortfolioGetParams(args [0]string, r *http.Request) (PortfolioGetParams, error) {
	var (
		params    PortfolioGetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeSandboxClearPostParams(args [0]string, r *http.Request) (SandboxClearPostParams, error) {
	var (
		params    SandboxClearPostParams
		queryArgs = r.URL.Query()
	)
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeSandboxCurrenciesBalancePostParams(args [0]string, r *http.Request) (SandboxCurrenciesBalancePostParams, error) {
	var (
		params    SandboxCurrenciesBalancePostParams
		queryArgs = r.URL.Query()
	)
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeSandboxPositionsBalancePostParams(args [0]string, r *http.Request) (SandboxPositionsBalancePostParams, error) {
	var (
		params    SandboxPositionsBalancePostParams
		queryArgs = r.URL.Query()
	)
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeSandboxRemovePostParams(args [0]string, r *http.Request) (SandboxRemovePostParams, error) {
	var (
		params    SandboxRemovePostParams
		queryArgs = r.URL.Query()
	)
	// Decode query: brokerAccountId.
	{
		values, ok := queryArgs["brokerAccountId"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(s)
					if err != nil {
						return err
					}

					paramsDotBrokerAccountIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.BrokerAccountId.SetTo(paramsDotBrokerAccountIdVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}
