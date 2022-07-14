package fields

type EnvironmentName string

const (
	Dev = EnvironmentName("dev")
)

func (environmentName EnvironmentName) IsEmpty() bool {
	return len(environmentName) == 0
}

func (environmentName EnvironmentName) IsDev() bool {
	return environmentName == Dev
}
