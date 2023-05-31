package common

// gracefully deprecating this package
// moving forward to use constant package instead

import (
	"github.com/bittorrent/go-common/v2/constant"
)

const (
	ContextHandlerKey = constant.HandlerNameContext
	ContextHTTPURLKey = constant.HTTPURLContext
)
