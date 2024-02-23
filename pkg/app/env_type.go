package app

type EnvironmentType uint8

const (
	DEV EnvironmentType = iota
	PROD
)

func ToEnvType(env string) EnvironmentType {
	switch env {
	case "prod", "production":
		return PROD
	default:
		return DEV
	}
}

func (e *EnvironmentType) ToString() string {
	switch *e {
	case PROD:
		return "prod"
	default:
		return "dev"
	}
}
