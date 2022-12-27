package vpc

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
)

type VpcArgs struct {
	Name                        string
	Description                 string
	Project                     string
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

func (vpc *Vpc) Create(ctx *pulumi.Context) (vpc *compute.Network, err error) {

	args := &compute.NetworkArgs{}
	args.Name = pulumi.String(vpc.Args.Name)
	args.Project = pulumi.String(vpc.Args.Project)

	vpcNetwork, err := compute.NewNetwork(ctx, vpc.Args.Name, args)
	ctx.Export("vpc", vpcNetwork)
	return vpcNetwork, err

}
