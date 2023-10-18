package router

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	t.Run("should return 404 NOT FOUND if path is unknown", func(t *testing.T) {
		// given
		router := New()
		router.GET("/the/route/without/params", func(w http.ResponseWriter, r *http.Request) {})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/unknown/route", nil)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should route to correct handler without params", func(t *testing.T) {
		// given
		router := New()
		router.GET("/the/route/without/params", func(w http.ResponseWriter, r *http.Request) {})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/the/route/without/params", nil)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should route to correct handler with params", func(t *testing.T) {
		// given
		router := New()
		var ctx context.Context
		router.GET("/the/:route/with/:params", func(w http.ResponseWriter, r *http.Request) {
			ctx = r.Context()
		})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/the/route/with/params", nil)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "route", ctx.Value("route"))
		assert.Equal(t, "params", ctx.Value("params"))
	})

	t.Run("test post request", func(t *testing.T) {
		// given
		router := New()
		var ctx context.Context
		router.POST("/the/:route/with/:params", func(w http.ResponseWriter, r *http.Request) {
			ctx = r.Context()
		})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/the/route/with/params", nil)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "route", ctx.Value("route"))
		assert.Equal(t, "params", ctx.Value("params"))
	})

	t.Run("test put request", func(t *testing.T) {
		// given
		router := New()
		var ctx context.Context
		router.PUT("/the/:route/with/:params", func(w http.ResponseWriter, r *http.Request) {
			ctx = r.Context()
		})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/the/route/with/params", nil)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "route", ctx.Value("route"))
		assert.Equal(t, "params", ctx.Value("params"))
	})

	t.Run("test delete request", func(t *testing.T) {
		// given
		router := New()
		var ctx context.Context
		router.DELETE("/the/:route/with/:params", func(w http.ResponseWriter, r *http.Request) {
			ctx = r.Context()
		})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("DELETE", "/the/route/with/params", nil)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "route", ctx.Value("route"))
		assert.Equal(t, "params", ctx.Value("params"))
	})
}
