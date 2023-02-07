package load

import (
	"encoding/json"

	"github.com/jeffweiss/gopsutil/internal/common"
)

var invoke common.Invoker

func init() {
	invoke = common.Invoke{}
}

type AvgStat struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

func (l AvgStat) String() string {
	s, _ := json.Marshal(l)
	return string(s)
}

type MiscStat struct {
	ProcsRunning int `json:"procsRunning"`
	ProcsBlocked int `json:"procsBlocked"`
	Ctxt         int `json:"ctxt"`
}

func (m MiscStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
