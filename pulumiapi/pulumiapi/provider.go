package pulumiapi

import (
	"fmt"
	"os"

	_ "github.com/pulumi/pulumi/sdk/v2/go/pulumi" // ??

	"github.com/pulumi/pulumi/pkg/v2/backend/httpstate"
	"github.com/pulumi/pulumi/pkg/v2/backend/httpstate/client"
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag/colors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/provider"
)

type ProviderArgs struct {
	// The URL of the Pulumi API server.
	Endpoint string `pulumi:"endpoint,optional"`
	// The access token to use for authentication.
	AccessToken string `pulumi:"accessToken,optional"`
}

// Provider implements resource operations for a Pulumi API.
//
//pulumi:provider
type Provider struct {
	sink        diag.Sink
	accessToken string
	client      *client.Client
}

func (p *Provider) Configure(ctx *provider.Context, args *ProviderArgs, options interface{}) error {
	endpoint := args.Endpoint
	if endpoint == "" {
		endpoint = httpstate.DefaultURL()
	}

	token := args.AccessToken
	if token == "" {
		account, err := workspace.GetAccount(endpoint)
		if err != nil {
			return fmt.Errorf("getting stored credentials: %v", err)
		}
		token = account.AccessToken
	}

	p.sink = diag.DefaultSink(os.Stdout, os.Stderr, diag.FormatOptions{Color: colors.Auto})
	p.client = client.NewClient(endpoint, token, p.sink)
	return nil
}
