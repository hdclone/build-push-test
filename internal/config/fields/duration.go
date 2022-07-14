package fields

import (
	"gopkg.in/yaml.v3"
	"time"
)

type DurationConfig struct {
	duration time.Duration
}

func (d *DurationConfig) parse(value string) error {
	if duration, err := time.ParseDuration(value); err != nil {
		return err
	} else {
		d.duration = duration
	}
	return nil
}

func (d *DurationConfig) Set(duration time.Duration) {
	d.duration = duration
}

func (d DurationConfig) Duration() time.Duration {
	return d.duration
}

func (d *DurationConfig) UnmarshalYAML(value *yaml.Node) error {
	return d.parse(value.Value)
}

func (d *DurationConfig) Decode(value string) error {
	return d.parse(value)

}
