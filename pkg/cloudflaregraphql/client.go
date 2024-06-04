package cloudflaregraphql

import (
	"errors"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/zeet-dev/pkg/utils"
	"github.com/zeet-dev/pkg/utils/options"
	"k8s.io/client-go/transport"
)

type Client struct {
	gqlV4 graphql.Client

	opt ClientOption
}

const (
	defaultServerURL = "https://api.cloudflare.com/client/v4"
	defaultUserAgent = "zeet-dev/cloudflare-graphql-go"
)

type ClientOption struct {
	ServerURL          string // optional: cloudflare API server URL
	CloudflareAPIToken string // required: cloudflare API token

	Debug     bool   // optional: enable debug logging
	UserAgent string // optional: user agent string
}

func New(opts ...options.MustOption[ClientOption]) (*Client, error) {
	opt := options.MustNewWithDefaults(ClientOption{
		ServerURL: defaultServerURL,
		UserAgent: defaultUserAgent,
		Debug:     false,
	}, opts...)

	if opt.CloudflareAPIToken == "" {
		return nil, errors.New("CloudflareAPIToken is required")
	}

	httpClient := newHTTPClient(opt)

	gqlPath := utils.URLJoin(opt.ServerURL, "graphql")

	return &Client{
		gqlV4: graphql.NewClient(gqlPath, httpClient),
		opt:   opt,
	}, nil
}

func newHTTPClient(opt ClientOption) *http.Client {
	tp := http.DefaultTransport
	if opt.Debug {
		tp = utils.LoggingHttpTransport
	}
	tp = transport.NewUserAgentRoundTripper(opt.UserAgent, tp)

	return &http.Client{
		Transport: transport.NewBearerAuthRoundTripper(opt.CloudflareAPIToken, tp),
	}
}
