package example

type EchoArgs struct {
	// A value to echo.
	InputValue interface{} `pulumi:"inputValue"`
}

type EchoResult struct {
	// The input value.
	OutputValue interface{} `pulumi:"outputValue"'`
}

// Echo returns its input argument as-is.
//
//pulumi:function
func Echo(p *Provider, args *EchoArgs, options interface{}) (*EchoResult, error) {
	return &EchoResult{OutputValue: args.InputValue}, nil
}
