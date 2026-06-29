package xuuid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/parocube/backroom/xerr"
	"github.com/parocube/backroom/xlog"
)

// ParseStringToUuid 将字符串解析为 UUID，解析成功返回 UUID 和 true，失败返回 uuid.Nil 和 false。
func ParseStringToUuid(s string) (uuid.UUID, bool) {
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, false
	}

	return id, true
}

// GetIdFromParam 从 Gin 路由参数 id 中获取 UUID。
func GetIdFromParam(c *gin.Context) (uuid.UUID, bool) {
	return getUuidFromContextParam(c, "id")
}

// getUuidFromContextParam 从 Gin 路由参数中按指定 key 获取 UUID，失败时记录错误并写入 invalid uuid 错误。
func getUuidFromContextParam(c *gin.Context, key string) (uuid.UUID, bool) {
	id, err := uuid.Parse(c.Param(key))
	if err != nil {
		xlog.Error(c.Request.Context(), err, "xuuid: get uuid from context param failed: "+key)
		_ = c.Error(xerr.ErrInvalidUUID)
		return uuid.Nil, false
	}

	return id, true
}

// GetTraceIdFromHeader 从 Gin 上下文中获取 traceId，缺失或解析失败时生成默认 UUID。
func GetTraceIdFromHeader(c *gin.Context) (uuid.UUID, bool) {
	return getUuidFromContextHeader(c, "traceId", true)
}

// getUuidFromContextHeader 从 Gin 上下文中按指定 key 获取 UUID，可按需在失败时生成默认 UUID。
func getUuidFromContextHeader(c *gin.Context, key string, needDefault bool) (uuid.UUID, bool) {
	id, err := uuid.Parse(c.GetString(key))
	if err != nil {
		if needDefault {
			return uuid.Must(uuid.NewV7()), true
		}
		xlog.Error(c.Request.Context(), err, "xuuid: get uuid from context string failed: "+key)
		_ = c.Error(xerr.ErrInvalidUUID)
		return uuid.Nil, false
	}
	return id, true
}
