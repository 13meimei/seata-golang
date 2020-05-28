package undo

import (
	_struct "github.com/dk-lockdown/seata-golang/client/rm_datasource/sql/struct"
	"github.com/dk-lockdown/seata-golang/client/sqlparser"
)

type SqlUndoLog struct {
	SqlType sqlparser.SQLType
	TableName string
	BeforeImage *_struct.TableRecords
	AfterImage *_struct.TableRecords
}

type BranchUndoLog struct {
	Xid string

	BranchId int64

	SqlUndoLogs []*SqlUndoLog
}