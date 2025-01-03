package main

import (
	"github.com/Werner-Swart-83/pulumi-azure-app-config-native/sdk/go/azure-app-config"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := azure-app-config.NewRandom(ctx, "myRandomResource", &azure-app-config.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", map[string]interface{}{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
