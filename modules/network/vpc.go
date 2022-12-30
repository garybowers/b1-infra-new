package network

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"fmt"
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

	var projectId string
	vpc.Args.ProjectId.ApplyT(func(p string) error {
		projectId = fmt.Sprintf("%s", p)
		return nil
	})

	fmt.Println(projectId)

	vpcNetwork, err = compute.NewNetwork(ctx, fmt.Sprintf("%s-%s",projectId,args.Name), args)
	ctx.Export("vpc", vpcNetwork)
	return vpcNetwork, err
}
