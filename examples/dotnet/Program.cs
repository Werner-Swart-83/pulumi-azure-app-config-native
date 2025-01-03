using System.Collections.Generic;
using System.Linq;
using Pulumi;
using AzureAppConfig = Pulumi.AzureAppConfig;

return await Deployment.RunAsync(() =>
{
    var myRandomResource = new AzureAppConfig.Random("myRandomResource", new()
    {
        Length = 24,
    });

    AzureAppConfig.ConfigResource config = new("config", new()
    {
        Key = "myRandomResource",
        Value = "myValue"
    });

    return new Dictionary<string, object?>
    {
        ["output"] = new Dictionary<string, object?>
        {
            { "value", myRandomResource.Result },
        },
    };
});

