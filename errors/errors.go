package errors

import "fmt"

// CodeMsg is a struct that contains a code and a message.
// It implements the error interface.
type CodeMsg struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
}

func (c *CodeMsg) Error() string {
	return fmt.Sprintf("err_no: %d, err_msg: %s", c.ErrNo, c.ErrMsg)
}

// New creates a new CodeMsg.
func New(errNo int, errMsg string) error {
	return &CodeMsg{ErrNo: errNo, ErrMsg: errMsg}
}
