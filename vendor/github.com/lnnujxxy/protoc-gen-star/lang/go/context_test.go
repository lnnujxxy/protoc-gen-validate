package pgsgo

import (
	"testing"

	"github.com/lnnujxxy/protoc-gen-star"
	"github.com/stretchr/testify/assert"
)

func TestContext_Params(t *testing.T) {
	t.Parallel()

	p := pgs.Parameters{}
	p.SetStr("foo", "bar")
	ctx := InitContext(p)

	params := ctx.Params()
	assert.Equal(t, "bar", params.Str("foo"))
}
