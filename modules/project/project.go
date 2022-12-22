package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
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

type Project struct {
	Args *ProjectArgs
        Name string
}

func (project *Project) Create(ctx *pulumi.Context) error {
	gcpProject, err := organizations.NewProject(ctx, project.Name, *project.Args)
	fmt.Println(gcpProject)
	return err
}
