import pulumi
import pulumi_azure-app-config as azure-app-config

my_random_resource = azure-app-config.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
