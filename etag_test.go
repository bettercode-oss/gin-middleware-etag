package etag

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpEtagCache_None_If_None_Match_Header(t *testing.T) {
	// given
	router := gin.Default()
	router.GET("/", HttpEtagCache(120), getTestProducts)

	// when
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "712d7783b1dc390cb7febeb72075e04f", w.Header().Get("ETag"))

	expected, _ := json.Marshal([]map[string]any{
		{
			"id":        1,
			"name":      "큰 잔",
			"listPrice": 1000,
		},
		{
			"id":        2,
			"name":      "작은 잔",
			"listPrice": 2000,
		},
	})
	assert.JSONEqf(t, string(expected), w.Body.String(), "aaa")
}

func TestHttpEtagCache_With_If_None_Match_Header(t *testing.T) {
	// given
	router := gin.Default()
	router.GET("/", HttpEtagCache(120), getTestProducts)

	// when
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("If-None-Match", "712d7783b1dc390cb7febeb72075e04f")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusNotModified, w.Code)
	assert.Equal(t, "712d7783b1dc390cb7febeb72075e04f", w.Header().Get("ETag"))
	assert.Equal(t, "max-age=120", w.Header().Get("Cache-Control"))
	assert.Nil(t, w.Body.Bytes())
}

func TestHttpEtagCache_With_If_None_Match_Header_And_CacheControl(t *testing.T) {
	// given
	router := gin.Default()
	router.GET("/", HttpEtagCache(500), getTestProducts)

	// when
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("If-None-Match", "712d7783b1dc390cb7febeb72075e04f")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusNotModified, w.Code)
	assert.Equal(t, "712d7783b1dc390cb7febeb72075e04f", w.Header().Get("ETag"))
	assert.Equal(t, "max-age=500", w.Header().Get("Cache-Control"))
	assert.Nil(t, w.Body.Bytes())
}

func getTestProducts(c *gin.Context) {
	products := []map[string]any{
		{
			"id":        1,
			"name":      "큰 잔",
			"listPrice": 1000,
		},
		{
			"id":        2,
			"name":      "작은 잔",
			"listPrice": 2000,
		},
	}

	c.JSON(http.StatusOK, products)
}
