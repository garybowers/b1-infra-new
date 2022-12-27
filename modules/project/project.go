package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"log"
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
	Args ProjectArgs
	Name string
}

func (project *Project) Create(ctx *pulumi.Context) (gcpProject *organizations.Project, err error) {
	args := &organizations.ProjectArgs{}
	args.Name = pulumi.String(project.Args.Name)
	if project.Args.ProjectId == "" {
		postfix, err := random.NewRandomString(ctx, "random", &random.RandomStringArgs{
			Length:  pulumi.Int(6),
			Special: pulumi.Bool(false),
			Upper:   pulumi.Bool(false),
			Lower:   pulumi.Bool(false),
		})
		if err != nil {
			log.Println(err)
		}

		args.ProjectId = pulumi.Sprintf("%s-%s", project.Args.Name, postfix.Result)
	} else {
		args.ProjectId = pulumi.String(project.Args.ProjectId)
	}

	args.BillingAccount = pulumi.String(project.Args.BillingAccount)
	args.AutoCreateNetwork = pulumi.Bool(project.Args.AutoCreateNetwork)

	if project.Args.FolderId == "" {
		args.OrgId = pulumi.String(project.Args.OrgId)
	} else {
		args.FolderId = pulumi.String(project.Args.FolderId)
	}

	//args.Labels = project.Args.Labels

	gcpProject, err = organizations.NewProject(ctx, project.Args.Name, args)
	ctx.Export("project", gcpProject)
	return gcpProject, err
}
