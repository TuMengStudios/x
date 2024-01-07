package errors

import (
	"fmt"
	"net/http"
)

// CodeMsg is a struct that contains a code and a message.
// It implements the error interface.
type CodeMsg struct {
	ErrNo      int    `json:"err_no"`
	ErrMsg     string `json:"err_msg"`
	StatusCode int    `json:"-"`
}

func (c *CodeMsg) Error() string {
	return fmt.Sprintf("err_no: %d, err_msg: %s", c.ErrNo, c.ErrMsg)
}

// New creates a new CodeMsg.
func New(statusCode int, errNo int, errMsg string) error {
	if _, ok := statusMap[statusCode]; !ok {
		fmt.Println("unexpected status code", statusCode)
		statusCode = 200
	}
	return &CodeMsg{ErrNo: errNo, ErrMsg: errMsg, StatusCode: statusCode}
}

var (
	statusMap = map[int]struct{}{
		http.StatusContinue:           {},
		http.StatusSwitchingProtocols: {},
		http.StatusProcessing:         {},
		http.StatusEarlyHints:         {},

		http.StatusOK:                   {},
		http.StatusCreated:              {},
		http.StatusAccepted:             {},
		http.StatusNonAuthoritativeInfo: {},
		http.StatusNoContent:            {},
		http.StatusResetContent:         {},
		http.StatusPartialContent:       {},
		http.StatusMultiStatus:          {},
		http.StatusAlreadyReported:      {},
		http.StatusIMUsed:               {},

		http.StatusMultipleChoices:   {},
		http.StatusMovedPermanently:  {},
		http.StatusFound:             {},
		http.StatusSeeOther:          {},
		http.StatusNotModified:       {},
		http.StatusUseProxy:          {},
		http.StatusTemporaryRedirect: {},
		http.StatusPermanentRedirect: {},

		http.StatusBadRequest:                   {},
		http.StatusUnauthorized:                 {},
		http.StatusPaymentRequired:              {},
		http.StatusForbidden:                    {},
		http.StatusNotFound:                     {},
		http.StatusMethodNotAllowed:             {},
		http.StatusNotAcceptable:                {},
		http.StatusProxyAuthRequired:            {},
		http.StatusRequestTimeout:               {},
		http.StatusConflict:                     {},
		http.StatusGone:                         {},
		http.StatusLengthRequired:               {},
		http.StatusPreconditionFailed:           {},
		http.StatusRequestEntityTooLarge:        {},
		http.StatusRequestURITooLong:            {},
		http.StatusUnsupportedMediaType:         {},
		http.StatusRequestedRangeNotSatisfiable: {},
		http.StatusExpectationFailed:            {},
		http.StatusTeapot:                       {},
		http.StatusMisdirectedRequest:           {},
		http.StatusUnprocessableEntity:          {},
		http.StatusLocked:                       {},
		http.StatusFailedDependency:             {},
		http.StatusTooEarly:                     {},
		http.StatusUpgradeRequired:              {},
		http.StatusPreconditionRequired:         {},
		http.StatusTooManyRequests:              {},
		http.StatusRequestHeaderFieldsTooLarge:  {},
		http.StatusUnavailableForLegalReasons:   {},

		http.StatusInternalServerError:           {},
		http.StatusNotImplemented:                {},
		http.StatusBadGateway:                    {},
		http.StatusServiceUnavailable:            {},
		http.StatusGatewayTimeout:                {},
		http.StatusHTTPVersionNotSupported:       {},
		http.StatusVariantAlsoNegotiates:         {},
		http.StatusInsufficientStorage:           {},
		http.StatusLoopDetected:                  {},
		http.StatusNotExtended:                   {},
		http.StatusNetworkAuthenticationRequired: {},
	}
)
