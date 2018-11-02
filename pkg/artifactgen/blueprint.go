package artifactgen

import (
    "github.com/kun-lun/digester/pkg/common"
)

type NonIaaSPart struct {
    ProgrammingLanguage string
    Framework string
    Databases []common.Database
}

type IaaSPart struct {
    VMGroup []common.VM
}

type Blueprint struct {
    NonIaaSPart NonIaaSPart
    IaaSPart    IaaSPart
}

func (b Blueprint) ExposeArtifects() {
}
