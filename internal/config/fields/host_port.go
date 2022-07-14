package fields

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"net"
	"strconv"
)

type HostPortConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (r *HostPortConfig) IsEmpty() bool {
	return r.Host == "" || r.Port == 0
}

func (r *HostPortConfig) UnmarshalYAML(value *yaml.Node) error {
	return r.parse(value.Value)
}

func (r *HostPortConfig) parse(value string) error {
	if host, portStr, err := net.SplitHostPort(value); err != nil {
		return err
	} else if port, err := strconv.ParseInt(portStr, 10, 32); err != nil {
		return err
	} else {
		r.Port = int(port)
		if len(host) == 0 {
			r.Host = "127.0.0.1"
		} else {
			r.Host = host
		}
	}
	return nil
}

func (r *HostPortConfig) Decode(value string) error {
	return r.parse(value)
}

func (r *HostPortConfig) ToString() string {
	return fmt.Sprintf("%v:%v", r.Host, r.Port)
}
