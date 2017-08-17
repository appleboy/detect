package detect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var agent = "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us; Silk/1.1.0-80) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16 Silk-Accelerated=true"

func TestSetPlatform(t *testing.T) {
	u := New(agent)
	get := u.setPlatform()
	assert.True(t, get)
	assert.Equal(t, "Mac OS X", u.PlatForm)
}
