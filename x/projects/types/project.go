package types

import (
	"fmt"

	commontypes "github.com/lavanet/lava/common/types"
)

const (
	ADMIN_PROJECT_NAME        = "admin"
	ADMIN_PROJECT_DESCRIPTION = "default admin project"
)

func ProjectIndex(subscriptionAddress string, projectName string) string {
	return subscriptionAddress + "-" + projectName
}

func NewProject(subscriptionAddress string, projectName string, description string, enable bool) (Project, error) {
	if !ValidateProjectNameAndDescription(projectName, description) {
		return Project{}, fmt.Errorf("project name must be ASCII, cannot contain \",\" and its length must be less than %d."+
			" Name: %s. The project's description must also be ASCII and its length must be less than %d",
			MAX_PROJECT_NAME_LEN, projectName, MAX_PROJECT_DESCRIPTION_LEN)
	}

	return Project{
		Index:              ProjectIndex(subscriptionAddress, projectName),
		Subscription:       subscriptionAddress,
		Description:        description,
		ProjectKeys:        []ProjectKey{},
		AdminPolicy:        nil,
		SubscriptionPolicy: nil,
		UsedCu:             0,
		Enabled:            enable,
	}, nil
}

func ValidateProjectNameAndDescription(name string, description string) bool {
	if !commontypes.ValidateString(name, commontypes.NAME_RESTRICTIONS, nil) ||
		len(name) > MAX_PROJECT_NAME_LEN || len(description) > MAX_PROJECT_DESCRIPTION_LEN ||
		!commontypes.ValidateString(description, commontypes.DESCRIPTION_RESTRICTIONS, nil) {
		return false
	}

	return true
}

func NewProjectKey(key string) ProjectKey {
	return ProjectKey{Key: key}
}

func (projectKey ProjectKey) AddType(kind ProjectKey_Type) ProjectKey {
	projectKey.Kinds |= uint32(kind)
	return projectKey
}

func ProjectAdminKey(key string) ProjectKey {
	return NewProjectKey(key).AddType(ProjectKey_ADMIN)
}

func ProjectDeveloperKey(key string) ProjectKey {
	return NewProjectKey(key).AddType(ProjectKey_DEVELOPER)
}

func (projectKey ProjectKey) IsType(kind ProjectKey_Type) bool {
	return projectKey.Kinds&uint32(kind) != 0x0
}

func (projectKey ProjectKey) IsTypeValid() bool {
	const keyKindsAll = (uint32(ProjectKey_ADMIN) | uint32(ProjectKey_DEVELOPER))

	return projectKey.Kinds != uint32(ProjectKey_NONE) &&
		(projectKey.Kinds & ^keyKindsAll) == 0x0
}

func (project *Project) GetKey(key string) ProjectKey {
	for _, projectKey := range project.ProjectKeys {
		if projectKey.Key == key {
			return projectKey
		}
	}
	return ProjectKey{}
}

func (project *Project) AppendKey(key ProjectKey) {
	for i, projectKey := range project.ProjectKeys {
		if projectKey.Key == key.Key {
			project.ProjectKeys[i].Kinds |= key.Kinds
			return
		}
	}
	project.ProjectKeys = append(project.ProjectKeys, key)
}

func (project *Project) IsAdminKey(key string) bool {
	return project.Subscription == key || project.GetKey(key).IsType(ProjectKey_ADMIN)
}
