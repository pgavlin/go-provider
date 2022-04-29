module github.com/pgavlin/go-provider

go 1.14

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/hashicorp/hcl/v2 v2.6.0
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi/pkg/v2 v2.10.3-0.20200924002346-9363f606b69c
	github.com/pulumi/pulumi/sdk/v2 v2.2.1
	github.com/rakyll/statik v0.1.7
	golang.org/x/tools v0.0.0-20200922173257-82fe25c37531
)

replace github.com/pulumi/pulumi/sdk/v2 => ../../pulumi/pulumi/sdk

replace github.com/pulumi/pulumi/pkg/v2 => ../../pulumi/pulumi/pkg
