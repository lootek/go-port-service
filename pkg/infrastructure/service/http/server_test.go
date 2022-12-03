package http

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lootek/go-port-service/pkg/core/application"
	"github.com/lootek/go-port-service/pkg/infrastructure/repository/memory"
	"github.com/stretchr/testify/require"
)

func TestServer_Upsert(t *testing.T) {
	testCases := []struct {
		name         string
		body         string
		wantCode     int
		wantResponse string
	}{
		{
			"empty",
			"",
			http.StatusBadRequest,
			"",
		},

		// TODO: Add test cases with mocked JSON file
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewServer(application.NewPorts(memory.NewStorage()))

			req, err := http.NewRequest("POST", "/ports", bytes.NewBufferString(tc.body))
			require.NoError(t, err)

			w := httptest.NewRecorder()
			s.httpRouter.ServeHTTP(w, req)

			responseData, err := io.ReadAll(w.Body)
			require.NoError(t, err)

			require.Equal(t, tc.wantCode, w.Code)
			require.Equal(t, tc.wantResponse, string(responseData))
		})
	}
}

func TestNewServer(t *testing.T) {

}
