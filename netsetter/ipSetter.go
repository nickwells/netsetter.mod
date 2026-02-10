package netsetter

import (
	"fmt"
	"net"

	"github.com/nickwells/param.mod/v7/psetter"
)

// IP allows you to give a parameter that can be used to set a network IP
// address
type IP struct {
	psetter.ValueReqMandatory

	Value *net.IP
}

// SetWithVal will attempt to parse the value into an IP address using the
// net.ParseIP function and will return an error if the parameter cannot be
// converted into an IP address (if net.ParseIP returns nil). Only if the
// value is parsed successfully and no checks are violated is the Value set.
func (s IP) SetWithVal(_, paramVal string) error {
	addr := net.ParseIP(paramVal)
	if addr == nil {
		return fmt.Errorf("could not convert %q into an IP address", paramVal)
	}

	*s.Value = addr

	return nil
}

// AllowedValues returns a string describing the allowed values
func (s IP) AllowedValues() string {
	return "any value that can be interpretted as an IP address." +
		" Either IPv4 ('192.167.0.1') or" +
		" IPv6 ('2001:db8::68') forms are allowed"
}

// CurrentValue returns the current setting of the parameter value
func (s IP) CurrentValue() string {
	return fmt.Sprintf("%v", *s.Value)
}

// CheckSetter panics if the setter has not been properly created - if the
// Value is nil.
func (s IP) CheckSetter(name string) {
	if s.Value == nil {
		panic(psetter.NilValueMessage(name, "netsetter.IP"))
	}
}
