package strip

import (
	"flag"
	"net"
)

type StringNetIP string

func (sip StringNetIP) NetIP() net.IP {
	return net.ParseIP(string(sip))
}

func (sip *StringNetIP) FromNetIP(ip net.IP) {
	*sip = StringNetIP(ip.String())
}

func (sip *StringNetIP) Set(val string) error {
	*sip = StringNetIP(net.ParseIP(val))
	return nil
}

func (sip *StringNetIP) Get() interface{} {
	return sip
}

func (sip *StringNetIP) String() string {
	return sip.NetIP().String()
}

func (sip StringNetIP) MarshalText() ([]byte, error) {
	return []byte(sip.String()), nil
}

func (sip *StringNetIP) UnmarshalText(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	return sip.Set(string(b))
}

func FromNetIP(ip net.IP) StringNetIP {
	return StringNetIP(ip.String())
}

func ConfinatorFlagVarTypeFunc(fs *flag.FlagSet, varPtr interface{}, name, usage string) {
	fs.Var(varPtr.(*StringNetIP), name, usage)
}
