package config

import (
	"github.com/greenplum-db/gpdb/gp/internal/pkg/enums"
)

type Database struct {
	DeploymentType         enums.DeploymentType `json:"deploymentType"`
	Admin                  User                 `json:"admin"`
	SegmentsPerSegmentHost int                  `json:"segmentsPerSegmentHost"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Artefact struct {
	Greenplum      string   `json:"greenplum"`
	DependencyList []string `json:"dependencyList"`
}
