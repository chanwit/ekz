package manifests

import _ "embed"

//go:embed storageclass.yaml
var StorageClass string

//go:embed metallb.yaml
var MetalLB string

//go:embed metallb_config.yaml
var MetalLBConfig string
