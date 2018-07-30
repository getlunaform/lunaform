package workers

import (
	"github.com/getlunaform/lunaform/models"
	"fmt"
	"path/filepath"
)

type BuildSpace struct {
	Module        *models.ResourceTfModule
	Stack         *models.ResourceTfStack
	Deployment    *models.ResourceTfDeployment
	rootDirectory string
}

func (bs *BuildSpace) StackDirectory(full bool) (path string, err error) {
	if bs.Stack == nil {
		return "", fmt.Errorf("Stack can not be 'nil'")
	}
	stackDirName := fmt.Sprintf("stack-%", bs.Stack.ID)
	return filepath.Join(bs.rootDirectory, stackDirName), nil
}

func (bs *BuildSpace) DeploymentDirectory(full bool) (path string, err error) {
	if bs.Deployment == nil {
		return "", fmt.Errorf("Deployment can not be 'nil'")
	}
	stackDir, err := bs.StackDirectory(full)
	if err != nil {
		return "", err
	}
	deploymentDirName := fmt.Sprintf("deployment-%s", bs.Deployment.ID)
	return filepath.Join(stackDir, deploymentDirName), nil

}

func (bs *BuildSpace) PlanPath(full bool) (path string, err error) {
	return bs.deploymentFile("deployment.plan", full)
}

func (bs *BuildSpace) VarFilePath(full bool) (path string, err error) {
	return bs.deploymentFile("variables_override.tf", full)
}

func (bs *BuildSpace) deploymentFile(fileName string, full bool) (path string, err error) {
	deploymentDir, err := bs.DeploymentDirectory(full)
	if err != nil {
		return "", err
	}
	return filepath.Join(deploymentDir, fileName), nil

}
