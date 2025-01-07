package cmn

type Config struct {
	Environment  *uint8    `yaml:"environment" validate:"gte=1,lte=2"`
	FrontendPath *string   `yaml:"frontendPath" validate:"required"`
	Address      *string   `yaml:"address" validate:"required"`
	Cors         []*string `yaml:"cors" validate:"required"`
}
