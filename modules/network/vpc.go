package network

import (
	"fmt"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"log"
)

type VpcArgs struct {
	Name                        string
	Description                 string
	ProjectId                   pulumi.StringOutput
	RoutingMode                 string
	AutoCreateSubnetworks       bool
	DeleteDefaultRoutesOnCreate bool
	EnableUIaInternalIpv6       bool
	InternalIpv6Range           bool
}

type Vpc struct {
	Args VpcArgs
	Name string
}

func (vpc *Vpc) Create(ctx *pulumi.Context) (vpcNetwork *compute.Network, err error) {
	args := &compute.NetworkArgs{}
	args.Name = pulumi.String(vpc.Args.Name)
	args.Project = vpc.Args.ProjectId
	args.AutoCreateSubnetworks = pulumi.Bool(vpc.Args.AutoCreateSubnetworks)

	vpc.Args.ProjectId.ApplyT(func(pid string) error {
		vpcNetwork, err = compute.NewNetwork(ctx, fmt.Sprintf("%s-%s", pid, args.Name), args)
		return err
	})
	if err != nil {
		log.Println(err)
	}

	ctx.Export("vpc", vpcNetwork)
	return vpcNetwork, err
}
