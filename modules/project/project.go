package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/orgnizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectArgs struct {
	ProjectId         string
	Name              string
	AutoCreateNetwork bool
	BillingAccount    string
	FolderId          string
	Labels            map[string]string
	OrgId             string
}

func Create(ctx *pulumi.Context) error {
	folder, err := organizations.NewProject(ctx, *Project)
}
