package awstpr

import (
	"github.com/giantswarm/awstpr/aws"
	"github.com/giantswarm/clustertpr"
)

type Spec struct {
	Cluster clustertpr.Cluster `json:"cluster" yaml:"cluster"`
	Aws     aws.Aws            `json:"aws" yaml:"aws"`
}
