package workers

import (
	"github.com/drewsonne/lunaform/server/models"
	"fmt"
	"os/exec"
	"io"
	"time"
)

type Terraform struct{}

type TerraformAction struct {
	action string
	cmd    *exec.Cmd
	out    *TerraformOutput
}
type TerraformOutput struct {
	Stderr io.ReadCloser
	Stdout io.ReadCloser
}

// NewTerraformClient return a struct which behaves like the cli terraform client.
//
// Between the unreliability of the internal interfaces in the terraform library and then
// need to communicate with providers, we'll wrap the terraform command in bash, rather
// than importing the `github.com/hashicorp/terraform` library and calling methods
// directly. See https://github.com/hashicorp/terraform/issues/12582 for more info.
func NewTerraformClient() *Terraform {
	return &Terraform{}
}

func (tf *Terraform) Plan(s *models.ResourceTfStack) (a *TerraformAction, out *TerraformOutput) {

	fmt.Print("Called Terraform.Plan")
	fmt.Print(s)

	out = &TerraformOutput{}

	a = &TerraformAction{
		action: "plan",
		out:    out,
	}
	return
}

func (a *TerraformAction) Run() (err error) {
	a.cmd = exec.Command("terraform " + a.action)
	time.Sleep(1000 * time.Millisecond)
	if a.out.Stderr, err = a.cmd.StderrPipe(); err != nil {
		return
	}
	if a.out.Stdout, err = a.cmd.StdoutPipe(); err != nil {
		return
	}
	return
}
