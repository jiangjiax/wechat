package context

import (
	"github.com/jiangjiax/wechat/v2/credential"
	"github.com/jiangjiax/wechat/v2/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
