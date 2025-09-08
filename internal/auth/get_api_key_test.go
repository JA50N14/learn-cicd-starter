package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  error
	}{
		"authPresent": {
			input: http.Header{"Host": []string{"notely"}, "Accept": []string{"application/json"}, "Authorization": []string{"ApiKey ashnjasdkfj3asdf35d"}},
			want:  nil,
		},
		"authMissing": {
			input: http.Header{"Host": []string{"notely"}, "User-Agent": []string{"curl/8.6.6"}, "Accept": []string{"application/json"}},
			want:  ErrNoAuthHeaderIncluded,
		},
		"authWrong": {
			input: http.Header{"Host": []string{"notely"}, "User-Agent": []string{"curl/8.6.6"}, "Accept": []string{"application/json"}, "Authorization": []string{"Bearer adf3jd5"}},
			want:  ErrMalformedAuthHeader,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, got := GetAPIKey(tc.input)
			if !errors.Is(tc.want, got) {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
