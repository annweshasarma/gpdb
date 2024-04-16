package config

import (
	"github.com/greenplum-db/gpdb/gp/internal/pkg/enums"
)

type Infra struct {
	RequestPort  int             `json:"requestPort"`
	PublishPort  int             `json:"publishPort"`
	Auth         *Authentication `json:"auth"`
	Coordinator  Host            `json:"coordinator"`
	SegmentHosts []Host          `json:"segmentHosts"`
}

type Authentication struct {
	Type     enums.AuthType `json:"type"`
	Password string         `json:"password"`
}

type Host struct {
	Network  Network         `json:"network"`
	Hostname string          `json:"hostname"`
	Auth     *Authentication `json:"auth"`
}

type Network struct {
	ExternalAddress string `json:"externalAddress"`
	InternalAddress string `json:"internalAddress"`
}
