import * as pulumi from "@pulumi/pulumi";
import * as azure-app-config from "@pulumi/azure-app-config";

const myRandomResource = new azure-app-config.Random("myRandomResource", {length: 24});
export const output = {
    value: myRandomResource.result,
};
