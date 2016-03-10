package configuration

import "fmt"

type HAProxy struct {
	TemplatePath            string
	OutputPath              string
	ReloadCommand           string
	ReloadValidationCommand string
	ReloadCleanupCommand    string
	ShutdownCommand         string
	GraceSeconds            int
	HostnameLabel           *string
}

// BalancerType indicates whether there bamboo is dealing with external or internal traffic
type BalancerType string

const (
	EmptyBalancerType    = ""
	InternalBalancerType = "internal"
	ExternalBalancerType = "external"
)

func (h HAProxy) BalancerType() BalancerType {
	if h.HostnameLabel == nil {
		return EmptyBalancerType
	}

	return BalancerType(*h.HostnameLabel)
}

const hostnameAclFormat = "hdr(host) -i %v"

// AclFormat returns a formatted acl
func AclFormat(hostname string) string {
	return fmt.Sprintf(hostnameAclFormat, hostname)
}
