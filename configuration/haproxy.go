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
	// EmptyBalancerType is used to indicate that there is supplied balancer name
	EmptyBalancerType = ""

	// hostnameAclFormat indicates the filter for hostname rules
	hostnameAclFormat = "hdr(host) -i %v"
)

// BalancerType contains information about the traffic that bamboo is routing
func (h HAProxy) BalancerType() BalancerType {
	if h.HostnameLabel == nil {
		return EmptyBalancerType
	}

	return BalancerType(*h.HostnameLabel)
}

// AclFormat returns a formatted host acl for HAProxy
func AclFormat(hostname string) string {
	return fmt.Sprintf(hostnameAclFormat, hostname)
}
