package main

import (
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/garybowers/b1-infra-new/modules/project"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := project.Create()
		fmt.Prinltn(err)
	})
}
