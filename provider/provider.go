// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/appconfig/azappconfig"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "azure-app-config"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single resource and component
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[Random, RandomArgs, RandomState](),
			infer.Resource[ConfigResource, ConfigResourceArgs, ConfigResourceState](),
		},
		Components: []infer.InferredComponent{
			infer.Component[*RandomComponent, RandomComponentArgs, *RandomComponentState](),
		},
		Config: infer.Config[Config](),
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index", // required because the folder with everything in is "provider"
		},
	})
}

// Define some provider-level configuration
type Config struct {
	Scream          *bool `pulumi:"itsasecret,optional"`
	AppConfigClient *azappconfig.Client
}

func (c *Config) Configure(ctx context.Context) error {

	cred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		// TODO: handle error
	}

	if c.AppConfigClient == nil {
		c.AppConfigClient, err = azappconfig.NewClient("https://myappconfig.azconfig.io", cred, nil)

		if err != nil {
			// TODO: handle error
		}
	}
	return nil
}
