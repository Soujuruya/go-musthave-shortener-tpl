package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-musthave-shortener-tpl/internal/repository"
	"go-musthave-shortener-tpl/internal/service"
)

func TestRedirectHandler(t *testing.T) {

	type want struct {
		code        int
		contentType string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code:        http.StatusTemporaryRedirect,
				contentType: "text/plain",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := repository.NewMemoryURLRepo()

			svc := service.NewURLService(repo)

			SetURLService(svc)

			//создание url
			postReq := httptest.NewRequest(http.MethodPost, "/", nil)
			postRec := httptest.NewRecorder()
			ShortenHandler(postRec, postReq)

			postRes := postRec.Result()

			defer postRes.Body.Close()

			shortURLByte, err := io.ReadAll(postRes.Body)
			require.NoError(t, err)
			shortURL := string(shortURLByte)

			id := shortURL[len("http://localhost:8080/"):]

			request := httptest.NewRequest(http.MethodGet, "/"+id, nil)
			w := httptest.NewRecorder()

			RedirectUrlHandler(w, request)

			res := w.Result()
			assert.Equal(t, res.StatusCode, test.want.code)

			loc := res.Header.Get("Location")
			assert.NotEmpty(t, loc)
		})
	}
}
