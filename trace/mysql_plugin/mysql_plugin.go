package mysql_plugin

import (
	"time"

	"github.com/yunkeCN/gokit/trace"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
)

const (
	callBackBeforeName = "trace:before"
	callBackAfterName  = "trace:after"
	startTime          = "_start_time"
	CSTLayout          = "2006-01-02 15:04:05" // CSTLayout China Standard Time Layout
)

type TracePlugin struct{}

func (op *TracePlugin) Name() string {
	return "tracePlugin"
}

func (op *TracePlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	_ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// 结束后
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

var _ gorm.Plugin = &TracePlugin{}

func before(db *gorm.DB) {
	db.InstanceSet(startTime, time.Now())
}

func after(db *gorm.DB) {
	ctx := db.Statement.Context
	ctxValue := ctx.Value(trace.TraceName)

	t, ok := ctxValue.(trace.T)
	if !ok {
		return
	}

	_ts, isExist := db.InstanceGet(startTime)
	if !isExist {
		return
	}

	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}

	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)

	sqlInfo := new(trace.SQL)
	sqlInfo.Timestamp = ts.Local().Format(CSTLayout)
	sqlInfo.SQL = sql
	sqlInfo.Stack = utils.FileWithLineNum()
	sqlInfo.Rows = db.Statement.RowsAffected
	sqlInfo.CostSeconds = time.Since(ts).Seconds()

	if t == nil {
		return
	}

	t.AppendSQL(sqlInfo)
}