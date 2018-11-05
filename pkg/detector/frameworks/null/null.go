package null

import (
    "github.com/kun-lun/digester/pkg/common"
)

type framework struct{}

func New() common.Framework {
    return &framework{}
}

func (f *framework) GetName() common.FrameworkName {
    return ""
}

func (f *framework) GetProgrammingLanguage() common.ProgrammingLanguage {
    return ""
}

func (f *framework) DetectConfig(path string) []common.Database {
    return []common.Database{}
}
