// Package snpn provides a map of service names and port numbers
// as registered with IANA per RFC6335.
package snpn

//go:generate go run gendata.go

// Service represents a service registered with IANA per RFC6335.
type Service struct {
	Name        string
	Ports       []string // "port/proto" or "port" if no protocol specified
	Description string
}

// FromName returns the Service assigned to name.
func FromName(name string) Service {
	return services[name]
}
