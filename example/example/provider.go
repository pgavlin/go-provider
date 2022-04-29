package example

type ProviderArgs struct {
	// The name of the provider.
	Name string `pulumi:"name"`
}

// This is an example Pulumi provider.
//
//pulumi:provider
type Provider struct{}

func (p *Provider) Configure(args *ProviderArgs, options interface{}) error {
	return nil
}
