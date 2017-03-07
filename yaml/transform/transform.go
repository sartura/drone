package transform

import "github.com/sartura/drone/yaml"

// TransformFunc defines an operation for transforming the Yaml file.
type TransformFunc func(*yaml.Config) error
