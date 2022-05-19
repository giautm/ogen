// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func decodeMarketCandlesGetParams(args [0]string, r *http.Request) (params MarketCandlesGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: figi.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Figi = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: from.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToTime(val)
				if err != nil {
					return err
				}

				params.From = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: from: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: to.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToTime(val)
				if err != nil {
					return err
				}

				params.To = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: to: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: interval.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "interval",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Interval = CandleResolution(c)
				return nil
			}); err != nil {
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
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}

func decodeMarketOrderbookGetParams(args [0]string, r *http.Request) (params MarketOrderbookGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: figi.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Figi = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: depth.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "depth",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt32(val)
				if err != nil {
					return err
				}

				params.Depth = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: depth: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}

func decodeMarketSearchByFigiGetParams(args [0]string, r *http.Request) (params MarketSearchByFigiGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: figi.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Figi = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}

func decodeMarketSearchByTickerGetParams(args [0]string, r *http.Request) (params MarketSearchByTickerGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: ticker.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "ticker",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Ticker = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: ticker: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}

func decodeOperationsGetParams(args [0]string, r *http.Request) (params OperationsGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: from.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToTime(val)
				if err != nil {
					return err
				}

				params.From = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: from: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: to.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToTime(val)
				if err != nil {
					return err
				}

				params.To = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: to: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: figi.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFigiVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		}
	}
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeOrdersCancelPostParams(args [0]string, r *http.Request) (params OrdersCancelPostParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: orderId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "orderId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.OrderId = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: orderId: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeOrdersGetParams(args [0]string, r *http.Request) (params OrdersGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeOrdersLimitOrderPostParams(args [0]string, r *http.Request) (params OrdersLimitOrderPostParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: figi.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Figi = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeOrdersMarketOrderPostParams(args [0]string, r *http.Request) (params OrdersMarketOrderPostParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: figi.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "figi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Figi = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: figi: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodePortfolioCurrenciesGetParams(args [0]string, r *http.Request) (params PortfolioCurrenciesGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodePortfolioGetParams(args [0]string, r *http.Request) (params PortfolioGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeSandboxClearPostParams(args [0]string, r *http.Request) (params SandboxClearPostParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeSandboxCurrenciesBalancePostParams(args [0]string, r *http.Request) (params SandboxCurrenciesBalancePostParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeSandboxPositionsBalancePostParams(args [0]string, r *http.Request) (params SandboxPositionsBalancePostParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}

func decodeSandboxRemovePostParams(args [0]string, r *http.Request) (params SandboxRemovePostParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: brokerAccountId.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "brokerAccountId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBrokerAccountIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
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
			}); err != nil {
				return params, errors.Wrap(err, "query: brokerAccountId: parse")
			}
		}
	}
	return params, nil
}
