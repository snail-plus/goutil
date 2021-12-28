package optional_test

import (
	"testing"

	"github.com/snail-plus/goutil/dump"
	"github.com/snail-plus/goutil/stdutil/optional"
)

func TestOptFactory_of(t *testing.T) {
	opt := optional.Of(nil)

	dump.P(opt.OrElseGet(34))
}
