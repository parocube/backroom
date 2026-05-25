package xuuid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/parocube/backroom/xerr"
	"github.com/parocube/backroom/xlog"
)

func GetIdFromParam(c *gin.Context) (uuid.UUID, bool) {
	return getUuidFromContextParam(c, "id")
}

func getUuidFromContextParam(c *gin.Context, key string) (uuid.UUID, bool) {
	id, err := uuid.Parse(c.Param(key))
	if err != nil {
		xlog.Error(c.Request.Context(), err, "xuuid: get uuid from context param failed: "+key)
		_ = c.Error(xerr.ErrInvalidUUID)
		return uuid.Nil, false
	}

	return id, true
}

func GetTraceIdFromHeader(c *gin.Context) (uuid.UUID, bool) {
	return getUuidFromContextHeader(c, "traceId", true)
}

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
