// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azureappconfig

import (
	"context"
	"reflect"

	"errors"
	"example.com/pulumi-azure-app-config/sdk/go/azure-app-config/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigResource struct {
	pulumi.CustomResourceState

	Key   pulumi.StringOutput `pulumi:"key"`
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewConfigResource registers a new resource with the given unique name, arguments, and options.
func NewConfigResource(ctx *pulumi.Context,
	name string, args *ConfigResourceArgs, opts ...pulumi.ResourceOption) (*ConfigResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigResource
	err := ctx.RegisterResource("azure-app-config:index:ConfigResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigResource gets an existing ConfigResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigResourceState, opts ...pulumi.ResourceOption) (*ConfigResource, error) {
	var resource ConfigResource
	err := ctx.ReadResource("azure-app-config:index:ConfigResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigResource resources.
type configResourceState struct {
}

type ConfigResourceState struct {
}

func (ConfigResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*configResourceState)(nil)).Elem()
}

type configResourceArgs struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ConfigResource resource.
type ConfigResourceArgs struct {
	Key   pulumi.StringInput
	Value pulumi.StringInput
}

func (ConfigResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configResourceArgs)(nil)).Elem()
}

type ConfigResourceInput interface {
	pulumi.Input

	ToConfigResourceOutput() ConfigResourceOutput
	ToConfigResourceOutputWithContext(ctx context.Context) ConfigResourceOutput
}

func (*ConfigResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigResource)(nil)).Elem()
}

func (i *ConfigResource) ToConfigResourceOutput() ConfigResourceOutput {
	return i.ToConfigResourceOutputWithContext(context.Background())
}

func (i *ConfigResource) ToConfigResourceOutputWithContext(ctx context.Context) ConfigResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigResourceOutput)
}

type ConfigResourceOutput struct{ *pulumi.OutputState }

func (ConfigResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigResource)(nil)).Elem()
}

func (o ConfigResourceOutput) ToConfigResourceOutput() ConfigResourceOutput {
	return o
}

func (o ConfigResourceOutput) ToConfigResourceOutputWithContext(ctx context.Context) ConfigResourceOutput {
	return o
}

func (o ConfigResourceOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigResource) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o ConfigResourceOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigResource) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigResourceInput)(nil)).Elem(), &ConfigResource{})
	pulumi.RegisterOutputType(ConfigResourceOutput{})
}
