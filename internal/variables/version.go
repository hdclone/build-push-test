package variables

import "fmt"

var (
	Version = "v0.0-0-g0000000"
)

func Service(name string) string {
	return fmt.Sprintf("Symbiosis %s", name)
}

func Banner(name string) string {
	return fmt.Sprintf("%s %s", Service(name), Version)
}
