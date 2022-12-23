package main

import (
	"fmt"
	"log"
	"github.com/garybowers/b1-infra-new/modules/project"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	var project project.Project

	pulumi.Run(func(ctx *pulumi.Context) error {
		project.Args.Name = "b1-services"
		project.Args.FolderId = "folders/415061719873"
		project.Args.BillingAccount = "01504C-A2522F-2110FA"
		project.Args.ProjectId = "b1-services-607482"
		project.Args.AutoCreateNetwork = false

		prj, err := project.Create(ctx)
		if err != nil {
			log.Println(err)
		}


		args := &compute.NetworkArgs{}
		args.Project = prj.ProjectId
		vpc, err := compute.NewNetwork(ctx, "vpc-1", &compute.NetworkArgs{
			Project: prj.ProjectId,
		})
		if err != nil {
			log.Println(err)
		}
		fmt.Println(vpc)
		return nil
	})
}
