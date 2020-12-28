package context

import (
	"github.com/jiangjiax/wechat/v2/credential"
	"github.com/jiangjiax/wechat/v2/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
