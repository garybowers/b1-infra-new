package project

import (
	"crypto/rand"
	"fmt"
	"io"
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

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func encodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

type Project struct {
	Args ProjectArgs
	Name string
}

func (project *Project) Create(ctx *pulumi.Context) (gcpProject *organizations.Project, err error) {
	args := &organizations.ProjectArgs{}
	args.Name = pulumi.String(project.Args.Name)
	if project.Args.ProjectId == "" {
		postfix := encodeToString(6)
		args.ProjectId = pulumi.String(project.Args.Name + "-" + postfix)
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

	fmt.Println(args)

	gcpProject, err = organizations.NewProject(ctx, project.Args.Name, args)
	fmt.Println(gcpProject)
	ctx.Export("project", gcpProject)
	return gcpProject, err
}
