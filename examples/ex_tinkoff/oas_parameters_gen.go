// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"time"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
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

				c, err := conv.ToDateTime(val)
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

				c, err := conv.ToDateTime(val)
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

type MarketOrderbookGetParams struct {
	// FIGI.
	Figi string
	// Глубина стакана [1..20].
	Depth int32
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

type MarketSearchByFigiGetParams struct {
	// FIGI.
	Figi string
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

type MarketSearchByTickerGetParams struct {
	// Тикер инструмента.
	Ticker string
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

				c, err := conv.ToDateTime(val)
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

				c, err := conv.ToDateTime(val)
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

type OrdersCancelPostParams struct {
	// ID заявки.
	OrderId string
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
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

type OrdersGetParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
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

type OrdersLimitOrderPostParams struct {
	// FIGI инструмента.
	Figi string
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
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

type OrdersMarketOrderPostParams struct {
	// FIGI инструмента.
	Figi string
	// Уникальный идентификатор счета (по умолчанию -
	// Тинькофф).
	BrokerAccountId OptString
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

type PortfolioCurrenciesGetParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
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

type PortfolioGetParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
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

type SandboxClearPostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
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

type SandboxCurrenciesBalancePostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
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

type SandboxPositionsBalancePostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
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

type SandboxRemovePostParams struct {
	// Номер счета (по умолчанию - Тинькофф).
	BrokerAccountId OptString
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
