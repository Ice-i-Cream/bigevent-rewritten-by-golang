package interceptors

import (
	"big_event/anno"
	"big_event/utils"
	"github.com/beego/beego/v2/server/web/context"
	"strings"
)

func isExcludedPath(urlPath string) bool {
	excludedPrefixes := []string{"/image"}
	excludedPaths := []string{"/", "/user/login", "/user/register"}

	for _, prefix := range excludedPrefixes {
		if strings.HasPrefix(urlPath, prefix) {
			return false
		}
	}
	for _, path := range excludedPaths {
		if urlPath == path {
			return false
		}
	}
	return true
}

var LoginInterceptor = func(c *context.Context) {
	urlPath := c.Request.URL.Path
	if isExcludedPath(urlPath) {
		ctx := anno.Ctx
		rdb := anno.RedisDb
		token := c.Request.Header.Get("Authorization")
		n, _ := rdb.Exists(ctx, token).Result()
		if n > 0 {
			claims, _ := utils.ParseToken(token)
			anno.Thread.Set(claims)
		} else {
			c.ResponseWriter.WriteHeader(401)
		}
	}
}
