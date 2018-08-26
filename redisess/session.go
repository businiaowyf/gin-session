package redisess

import (
	"github.com/boj/redistore"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var SessName string // the key name in cookies
var SessStore sessions.Store

func Init(name string, maxIdle int, address, password string, keyPairs ...[]byte) error {
	SessName = name
	store, err := redistore.NewRediStore(maxIdle, "tcp", address, password, keyPairs...)
	if err != nil {
		glog.Error("NewRediStore error=%v", err)
		return err
	}
	SessStore = store
	return nil
}

func NewSessionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// to avoid memory leak
		defer context.Clear(c.Request)
		c.Next()
	}
}
