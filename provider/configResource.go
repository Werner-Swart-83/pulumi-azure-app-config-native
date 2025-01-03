package provider

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type ConfigResource struct{}

type ConfigResourceArgs struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

type ConfigResourceState struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

func (c ConfigResource) Create(ctx context.Context, name string, input ConfigResourceArgs, preview bool) (string, ConfigResourceState, error) {

	config := infer.GetConfig[Config](ctx)

	_, err := config.AppConfigClient.AddSetting(ctx, input.Key, &input.Value, nil)
	if err != nil {
		return "", ConfigResourceState{}, err
	}

	state := ConfigResourceState(input)

	return name, state, nil
}
