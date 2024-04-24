package config

import (
	"github.com/greenplum-db/gpdb/gp/internal/pkg/enums"
)

type InfraConfig interface {
	GetRequestPort() int
	GetPublishPort() int
	GetCoordinator() HostConfig
	GetStandby() HostConfig
	GetSegmentHost() SegmentHostsConfig
}
type Infra struct {
	RequestPort  int                `json:"requestPort"`
	PublishPort  int                `json:"publishPort"`
	Coordinator  HostConfig         `json:"coordinatorHost"`
	Stanby       HostConfig         `json:"standbyHost"`
	SegmentHosts SegmentHostsConfig `json:"segmentHost"`
}

func (i Infra) GetRequestPort() int {
	return i.RequestPort
}
func (i Infra) GetPublishPort() int {
	return i.PublishPort
}
func (i Infra) GetCoordinator() HostConfig {
	return i.Coordinator
}
func (i Infra) GetStandby() HostConfig {
	return i.Stanby
}
func (i Infra) GetSegmentHost() SegmentHostsConfig {
	return i.SegmentHosts
}

type AuthenticationConfig interface {
	GetType() enums.AuthType
	GetPassword() string
}

type Authentication struct {
	Type     enums.AuthType `json:"type"`
	Password string         `json:"password"`
}

func (a Authentication) GetType() enums.AuthType {
	return a.Type
}
func (a Authentication) GetPassword() string {
	return a.Password
}

type HostConfig interface {
	GetNetwork() NetworkConfig
	GetHostname() string
	GetDomainName() string
	GetAuth() AuthenticationConfig
}

type Host struct {
	Network    NetworkConfig        `json:"coordinatorNetwork"`
	Hostname   string               `json:"hostname"`
	DomainName string               `json:"domainName"`
	Auth       AuthenticationConfig `json:"authentication"`
}

func (h Host) GetNetwork() NetworkConfig {
	return h.Network
}
func (h Host) GetHostname() string {
	return h.Hostname
}
func (h Host) GetDomainName() string {
	return h.DomainName
}
func (h Host) GetAuth() AuthenticationConfig {
	return h.Auth
}

type SegmentHostsConfig interface {
	GetSegmentHostsCount() int
	GetNetwork() SegmentHostsNetworkConfig
	GetAuthentication() AuthenticationConfig
	GetHostnamePrefix() string
	GetDomainName() string
}
type SegmentHosts struct {
	SegmentHostsCount int                       `json:"segmentHostCount"`
	Network           SegmentHostsNetworkConfig `json:"network"`
	Authentication    AuthenticationConfig      `json:"authentication"`
	HostnamePrefix    string                    `json:"hostnamePrefix"`
	DomainName        string                    `json:"domainName"`
}

func (s SegmentHosts) GetSegmentHostsCount() int {
	return s.SegmentHostsCount
}
func (s SegmentHosts) GetNetwork() SegmentHostsNetworkConfig {
	return s.Network
}
func (s SegmentHosts) GetAuthentication() AuthenticationConfig {
	return s.Authentication
}
func (s SegmentHosts) GetHostnamePrefix() string {
	return s.HostnamePrefix
}
func (s SegmentHosts) GetDomainName() string {
	return s.DomainName
}

type NetworkConfig interface {
	GetExternalAddress() string
	GetInternalAddress() string
}

type Network struct {
	ExternalAddress string `json:"externalIp"`
	InternalAddress string `json:"internalIp"`
}

func (n Network) GetExternalAddress() string {
	return n.ExternalAddress
}
func (n Network) GetInternalAddress() string {
	return n.InternalAddress
}

type SegmentHostsNetworkConfig interface {
	GetInternalCidr() string
	GetIpRange() IpRangeConfig
	GetIpList() []string
}

type SegmentHostsNetwork struct {
	InternalCidr string        `json:"internalCidr"`
	IpRange      IpRangeConfig `json:"ipRange"`
	IpList       []string      `json:"ipList"`
}

func (sn SegmentHostsNetwork) GetInternalCidr() string {
	return sn.InternalCidr
}
func (sn SegmentHostsNetwork) GetIpRange() IpRangeConfig {
	return sn.IpRange
}
func (sn SegmentHostsNetwork) GetIpList() []string {
	return sn.IpList
}

type IpRangeConfig interface {
	GetFirstIp() string
	GetLastIp() string
}
type IpRange struct {
	FirstIp string `json:"first"`
	LastIp  string `json:"last"`
}

func (ip IpRange) GetFirstIp() string {
	return ip.FirstIp
}
func (ip IpRange) GetLastIp() string {
	return ip.LastIp
}
