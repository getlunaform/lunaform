package workers

import (
	"github.com/getlunaform/lunaform/models"
	"fmt"
	"path/filepath"
)

type BuildSpace struct {
	module        *models.ResourceTfModule
	stack         *models.ResourceTfStack
	deployment    *models.ResourceTfDeployment
	rootDirectory string
}

//
// Initialisation functions
//

// NewBuilSpace returns an initialised build space struct
func NewBuildSpace(rootDirectory string) *BuildSpace {
	return &BuildSpace{
		rootDirectory: rootDirectory,
	}
}

// WithModule returns a BuildSpace struct after setting the module
func (bs *BuildSpace) WithModule(module *models.ResourceTfModule) *BuildSpace {
	bs.module = module
	return bs
}

// WithStack returns a BuildSpace struct after setting the stack
func (bs *BuildSpace) WithStack(stack *models.ResourceTfStack) *BuildSpace {
	bs.stack = stack
	return bs
}

// WithDeploy returns a BuildSpace struct after setting the deployment
func (bs *BuildSpace) WithDeployment(deployment *models.ResourceTfDeployment) *BuildSpace {
	bs.deployment = deployment
	return bs
}

// StackDirectory represents the (optionally) full path
func (bs *BuildSpace) StackDirectory(full bool) (path string, err error) {
	if bs.stack == nil {
		return "", fmt.Errorf("Stack can not be 'nil'")
	}
	stackDirName := fmt.Sprintf("stack-%s", bs.stack.ID)
	var rootDir string
	if full {
		rootDir = filepath.Join(bs.rootDirectory, stackDirName)
	} else {
		rootDir = stackDirName
	}
	return rootDir, nil
}

// MustStackDir
func (bs *BuildSpace) MustStackDir(full bool) (path string) {
	path, err := bs.StackDirectory(full)
	if err != nil {
		panic(err)
	}
	return path
}

func (bs *BuildSpace) DeploymentDirectory(full bool) (path string, err error) {
	if bs.deployment == nil {
		return "", fmt.Errorf("Deployment can not be 'nil'")
	}
	stackDir, err := bs.StackDirectory(full)
	if err != nil {
		return "", err
	}
	deploymentDirName := fmt.Sprintf("deployment-%s", bs.deployment.ID)
	return filepath.Join(stackDir, deploymentDirName), nil
}

func (bs *BuildSpace) MustDeploymentDirectory(full bool) (path string) {
	_, err := bs.StackDirectory(full)
	if err != nil {
		panic(err)
	}
	path, err = bs.DeploymentDirectory(full)
	if err != nil {
		panic(err)
	}
	return path
}

func (bs *BuildSpace) MustPlanPath(full bool) (path string) {
	path, err := bs.PlanPath(full)
	if err != nil {
		panic(err)
	}
	return path
}

func (bs *BuildSpace) MustProviderFilePath(full bool) (path string) {
	path, err := bs.ProviderFilePath(full)
	if err != nil {
		panic(err)
	}
	return path
}

func (bs *BuildSpace) PlanPath(full bool) (path string, err error) {
	return bs.deploymentFile("deployment.plan", full)
}

func (bs *BuildSpace) VarFilePath(full bool) (path string, err error) {
	return bs.deploymentFile("terraform.tfvars", full)
}

func (bs *BuildSpace) ProviderFilePath(full bool) (path string, err error) {
	return bs.deploymentFile("providers.tf", full)
}

func (bs *BuildSpace) deploymentFile(fileName string, full bool) (path string, err error) {
	deploymentDir, err := bs.DeploymentDirectory(full)
	if err != nil {
		return "", err
	}
	return filepath.Join(deploymentDir, fileName), nil

}
