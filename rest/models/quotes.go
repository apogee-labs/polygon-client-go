package models

// ListQuotesParams is the set of parameters for the ListQuotes method.
type ListQuotesParams struct {
	Ticker string `validate:"required" path:"ticker"`

	TimestampEQ  *Nanos `query:"timestamp"`
	TimestampLT  *Nanos `query:"timestamp.lt"`
	TimestampLTE *Nanos `query:"timestamp.lte"`
	TimestampGT  *Nanos `query:"timestamp.gt"`
	TimestampGTE *Nanos `query:"timestamp.gte"`

	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
	Sort  *Sort  `query:"sort"`
}

func (p ListQuotesParams) WithTimestamp(c Comparator, q Nanos) *ListQuotesParams {
	if c == EQ {
		p.TimestampEQ = &q
	} else if c == LT {
		p.TimestampLT = &q
	} else if c == LTE {
		p.TimestampLTE = &q
	} else if c == GT {
		p.TimestampGT = &q
	} else if c == GTE {
		p.TimestampGTE = &q
	}
	return &p
}

func (p ListQuotesParams) WithSort(q Sort) *ListQuotesParams {
	p.Sort = &q
	return &p
}

func (p ListQuotesParams) WithOrder(q Order) *ListQuotesParams {
	p.Order = &q
	return &p
}

func (p ListQuotesParams) WithLimit(q int) *ListQuotesParams {
	p.Limit = &q
	return &p
}

// ListQuotesResponse is the response returned by the ListQuotes method.
type ListQuotesResponse struct {
	BaseResponse
	Results []Quote `json:"results,omitempty"`
}

// GetLastQuoteParams is the set of parameters for the GetLastQuote method.
type GetLastQuoteParams struct {
	Ticker string `validate:"required" path:"ticker"`
}

// GetLastQuoteResponse is the response returned by the GetLastQuote method.
type GetLastQuoteResponse struct {
	BaseResponse
	Results LastQuote `json:"results,omitempty"`
}

// GetLastForexQuoteParams is the set of parameters for the GetLastForexQuote method.
type GetLastForexQuoteParams struct {
	From string `validate:"required" path:"from"`
	To   string `validate:"required" path:"to"`
}

// GetLastForexQuoteResponse is the response returned by the GetLastForexQuote method.
type GetLastForexQuoteResponse struct {
	BaseResponse
	Symbol string     `json:"symbol,omitempty"`
	Last   ForexQuote `json:"last,omitempty"`
}

// GetRealTimeCurrencyConversionParams is the set of parameters for the GetRealTimeCurrencyConversion method.
type GetRealTimeCurrencyConversionParams struct {
	From string `validate:"required" path:"from"`
	To   string `validate:"required" path:"to"`
}

// GetRealTimeCurrencyConversionResponse is the response returned by the GetRealTimeCurrencyConversion method.
type GetRealTimeCurrencyConversionResponse struct {
	BaseResponse
	InitialAmount float64    `json:"initialAmount,omitempty"`
	Converted     float64    `json:"converted,omitempty"`
	From          string     `json:"from,omitempty"`
	To            string     `json:"to,omitempty"`
	LastQuote     ForexQuote `json:"last,omitempty"`
}

// Quote is an NBBO for a ticker symbol in a given time range.
type Quote struct {
	AskExchange          int     `json:"ask_exchange,omitempty"`
	AskPrice             float64 `json:"ask_price,omitempty"`
	AskSize              float64 `json:"ask_size,omitempty"`
	BidExchange          int     `json:"bid_exchange,omitempty"`
	BidPrice             float64 `json:"bid_price,omitempty"`
	BidSize              float64 `json:"bid_size,omitempty"`
	Conditions           []int32 `json:"conditions,omitempty"`
	Indicators           []int32 `json:"indicators,omitempty"`
	ParticipantTimestamp Nanos   `json:"participant_timestamp,omitempty"`
	SequenceNumber       Nanos   `json:"sequence_number,omitempty"`
	SipTimestamp         Nanos   `json:"sip_timestamp,omitempty"`
	Tape                 int32   `json:"tape,omitempty"`
	TrfTimestamp         Nanos   `json:"trf_timestamp,omitempty"`
}

// LastQuote is the most recent NBBO for a ticker symbol.
type LastQuote struct {
	Ticker               string  `json:"T,omitempty"`
	AskExchange          int     `json:"X,omitempty"`
	AskPrice             float64 `json:"P,omitempty"`
	AskSize              float64 `json:"S,omitempty"`
	BidExchange          int     `json:"x,omitempty"`
	BidPrice             float64 `json:"p,omitempty"`
	BidSize              float64 `json:"s,omitempty"`
	Conditions           []int32 `json:"c,omitempty"`
	Indicators           []int32 `json:"i,omitempty"`
	ParticipantTimestamp Nanos   `json:"y,omitempty"`
	SequenceNumber       Nanos   `json:"q,omitempty"`
	SipTimestamp         Nanos   `json:"t,omitempty"`
	Tape                 int32   `json:"z,omitempty"`
	TrfTimestamp         Nanos   `json:"f,omitempty"`
}

// ForexQuote is a BBO for a forex currency pair.
type ForexQuote struct {
	Ask       float64 `json:"ask,omitempty"`
	Bid       float64 `json:"bid,omitempty"`
	Exchange  int     `json:"exchange,omitempty"`
	Timestamp Nanos   `json:"timestamp,omitempty"`
}
