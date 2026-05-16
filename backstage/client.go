package backstage

import (
	"fmt"
	"net/http"

	"github.com/datolabs-io/go-backstage/v3"
)

func getClient(host, token string) (*backstage.Client, error) {
	if host == "" {
		return nil, fmt.Errorf("host must be configured")
	}
	if token == "" {
		return nil, fmt.Errorf("token must be configured")
	}

	httpClient := &http.Client{
		Transport: &tokenRoundTripper{
			token:  token,
			client: http.DefaultTransport,
		},
	}
	client, err := backstage.NewClient(host, "", httpClient)
	if err != nil {
		return nil, fmt.Errorf("error creating backstage client: %v", err)
	}

	return client, nil
}

type tokenRoundTripper struct {
	token  string
	client http.RoundTripper
}

func (t *tokenRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.token))
	return t.client.RoundTrip(req)
}
