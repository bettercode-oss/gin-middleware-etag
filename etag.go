package etag

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ETag         = "ETag"
	CacheControl = "Cache-Control"
	IfNoneMatch  = "If-None-Match"
	MaxAge       = "max-age"
)

func HttpEtagCache(maxAge uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		c.Next()

		eTag := generateMD5Hash(w.body.String())
		w.ResponseWriter.Header().Set(ETag, eTag)

		if len(c.Request.Header.Get(IfNoneMatch)) > 0 {
			if eTag == c.Request.Header.Get(IfNoneMatch) {
				c.Writer.Header().Set(CacheControl, fmt.Sprintf("%s=%d", MaxAge, maxAge))
				c.Status(http.StatusNotModified)
				return
			}
		}

		_, err := w.ResponseWriter.Write(w.body.Bytes())
		if err != nil {
			panic(err)
		}
	}
}

func generateMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	return r.body.Write(b)
}
