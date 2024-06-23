package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ast := assert.New(t)
	c := New(http.StatusOK, 1, "test")
	cm, ok := c.(*CodeMsg)
	ast.True(ok)
	ast.NotNil(cm)
	ast.Equal(int(1), cm.ErrNo)
	ast.Equal("test", cm.ErrMsg)
}

func TestCodeMsg_Error(t *testing.T) {
	ast := assert.New(t)
	c := New(http.StatusOK, 1, "test")
	cm, ok := c.(*CodeMsg)
	ast.True(ok)
	ast.NotNil(cm)
	ast.NotEmpty(cm.Error())
}
