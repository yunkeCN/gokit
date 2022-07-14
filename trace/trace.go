package trace

import (
	"encoding/json"
	"sync"

	"github.com/yunkeCN/gokit/util"
)

const (
	TraceName = "_trace_"
)

type T interface {
	AppendSQL(sql *SQL) *Trace
}

// Trace 记录的参数
type Trace struct {
	mux  sync.Mutex
	SQLs []*SQL `json:"sqls"` // 执行的 SQL 信息
}

func (t *Trace) String() string {
	if t == nil {
		return ""
	}
	bytes, _ := json.Marshal(t)
	return util.Byte2Str(bytes)
}

// AppendSQL 追加 SQL
func (t *Trace) AppendSQL(sql *SQL) *Trace {
	if sql == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.SQLs = append(t.SQLs, sql)
	return t
}
