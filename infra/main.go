package main

import (
	"fmt"
	"log"
	"github.com/garybowers/b1-infra-new/modules/project"
	"github.com/garybowers/b1-infra-new/modules/vpc"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	var project project.Project
	var vpc  vpc.Vpc

	pulumi.Run(func(ctx *pulumi.Context) error {
		project.Args.Name = "b1-services"
		project.Args.FolderId = "folders/415061719873"
		project.Args.BillingAccount = "01504C-A2522F-2110FA"
		project.Args.AutoCreateNetwork = false

		prj, err := project.Create(ctx)
		if err != nil {
			log.Println(err)
		}

		vpc.Args.Name = "vpc-1"
		vpc.Args.Project = prj.ProjectId

		vpcNetwork, err := vpc.Create(ctx)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(vpcNetwork)
		return nil
	})
}
