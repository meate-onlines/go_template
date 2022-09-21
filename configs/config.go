package configs

type Server struct {
	System        System        `mapstructure:"system" json:"system" yaml:"system"`
	Elasticsearch Elasticsearch `mapstructure:"elasticsearch" json:"elasticsearch" yaml:"elasticsearch"`
	Xxl           Xxl           `json:"xxl" yaml:"xxl" mapstructure:"xxl"`
}