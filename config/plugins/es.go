package plugins

type ES struct {
	Url   string `mapstructure:"url" json:"url" yaml:"url"`
	Index string `mapstructure:"index" json:"index" yaml:"index"`
}
