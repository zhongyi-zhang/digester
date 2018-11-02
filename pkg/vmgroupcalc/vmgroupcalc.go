package vmgroupcalc

import (
    "github.com/kun-lun/digester/pkg/common"
)

type Requirment struct {
    ConcurrentUserNumber int
}

func Calc(r Requirment) []common.VM {
    res := []common.VM{}
    res = append(res, common.VM{"Standard_DS1_v2"})
    x := r.ConcurrentUserNumber
    if x >= 500 {

    }
    if x >= 1500 {
        res = append(res, common.VM{"Standard_DS1_v2"})
    }
    if x >= 2500 {
        res = append(res, common.VM{"Standard_DS1_v2"})
        res = append(res, common.VM{"Standard_DS1_v2"})
    }
    return res
}
