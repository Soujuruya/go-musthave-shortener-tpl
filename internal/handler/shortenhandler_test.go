package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-musthave-shortener-tpl/internal/repository"
	"go-musthave-shortener-tpl/internal/service"
)

func TestShortenHandler(t *testing.T) {

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
				code:        http.StatusCreated,
				contentType: "text/plain",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			repo := repository.NewMemoryURLRepo()

			svc := service.NewURLService(repo)

			// передаём в обработчик
			SetURLService(svc)

			request := httptest.NewRequest(http.MethodPost, "/", nil)

			w := httptest.NewRecorder()
			ShortenHandler(w, request)

			res := w.Result()

			assert.Equal(t, res.StatusCode, test.want.code)

			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)

			pattern := `^http://localhost:8080/.+$`
			matched, err := regexp.MatchString(pattern, string(resBody))

			if err != nil {
				t.Fatalf("regexp error: %v", err)
			}

			if !matched {
				t.Errorf("URL does not match pattern: %s", resBody)
			}

			assert.Equal(t, res.Header.Get("Content-Type"), test.want.contentType)
			assert.True(t, matched, "Response does not match pattern. Got: %s", resBody)
		})
	}
}
