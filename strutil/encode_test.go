package strutil_test

import (
	"testing"

	"github.com/snail-plus/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestEscape(t *testing.T) {
	tests := struct{ give, want string }{
		"<p>some text</p>",
		"&lt;p&gt;some text&lt;/p&gt;",
	}

	assert.Equal(t, tests.want, strutil.EscapeHTML(tests.give))

	ret := strutil.EscapeJS("<script>var a = 23;</script>")
	assert.NotContains(t, ret, "<script>")
	assert.NotContains(t, ret, "</script>")
}

func TestBase64(t *testing.T) {

}
