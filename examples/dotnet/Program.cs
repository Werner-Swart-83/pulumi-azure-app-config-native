using System.Collections.Generic;
using System.Linq;
using Pulumi;
using azure-app-config = Pulumi.azure-app-config;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new azure-app-config.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

