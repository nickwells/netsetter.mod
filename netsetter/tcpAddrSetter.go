package netsetter

import (
	"fmt"
	"net"

	"github.com/nickwells/param.mod/v6/psetter"
)

// TCPAddr allows you to give a parameter that can be used to set a
// net.TCPAddr address. Note that the Value is a pointer to a pointer so if
// you have a
//
//	var addr *net.TCPAddr
//
// then you should set the Value as follows:
//
//	netsetter.TCPAddr{Value: &addr}
//
// The standard library net package uses a pointer to a TCPAddr throughout so
// this is likely the most useful way of setting this value.
type TCPAddr struct {
	psetter.ValueReqMandatory

	Value **net.TCPAddr
}

// SetWithVal will attempt to parse the value into an TCPAddr address using
// the net.ResolveTCPAddr function and will return an error if the parameter
// cannot be converted into an TCPAddr address. Only if the parameter is
// successfully converted into a TCPAddr is the Value updated.
func (s TCPAddr) SetWithVal(_, paramVal string) error {
	addr, err := net.ResolveTCPAddr("tcp", paramVal)
	if err != nil {
		return fmt.Errorf("Could not convert %q into a TCP address: %w",
			paramVal, err)
	}

	*s.Value = addr
	return nil
}

// AllowedValues returns a string describing the allowed values
func (s TCPAddr) AllowedValues() string {
	return "any value that can be interpretted as a TCP address." +
		" Either IPv4 ('192.167.0.1:8080') or" +
		" IPv6 ('[2001:db8::68]:8080') forms are allowed"
}

// CurrentValue returns the current setting of the parameter value
func (s TCPAddr) CurrentValue() string {
	return fmt.Sprintf("%v", *s.Value)
}

// CheckSetter panics if the setter has not been properly created - if the
// Value is nil.
func (s TCPAddr) CheckSetter(name string) {
	if s.Value == nil {
		panic(psetter.NilValueMessage(name, "netsetter.TCPAddr"))
	}
}
