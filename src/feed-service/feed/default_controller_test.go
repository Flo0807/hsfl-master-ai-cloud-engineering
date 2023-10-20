package feed

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultController_GetFeed(t *testing.T) {

	controller := NewDefaultController()

	t.Run("GetFeed", func(t *testing.T) {
		t.Run("should return 500 INTERNAL SERVER ERROR if query failed", func(t *testing.T) {

			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/feed", nil)

			// when
			controller.GetFeed(w, r)

			// then
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

	})
}
