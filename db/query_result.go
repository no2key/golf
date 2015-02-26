// query_result
package db

import (
	"database/sql"

	"github.com/wdsgyj/golf/conv"
)

type QueryResult struct {
	rows    *sql.Rows
	buf     []interface{}
	values  map[string]*interface{}
	columns []string
	filled  bool
}

func NewQueryResult(rows *sql.Rows) *QueryResult {
	rs := &QueryResult{rows: rows}
	rs.fillColumns() // 填充列名
	rs.filled = false
	rs.buf = conv.NewInterfacePtrs(rs.ColumnCount())
	return rs
}

func (self *QueryResult) fillColumns() error {
	var err error
	self.columns, err = self.rows.Columns()
	return err
}

func (self *QueryResult) Next() bool {
	b := self.rows.Next()
	if b {
		if e := self.rows.Scan(self.buf...); e == nil {
			self.filled = true
			for i, name := range self.columns {
				self.values[name] = self.buf[i].(*interface{})
			}
		} else {
			self.filled = false
		}
	} else {
		self.filled = false
	}
	return b
}

func (self *QueryResult) Close() error {
	self.filled = false
	return self.rows.Close()
}

func (self *QueryResult) ColumnCount() int {
	return len(self.columns)
}

func (self *QueryResult) ColumnNames() []string {
	return self.columns
}

func (self *QueryResult) getOf(i int) interface{} {
	return *(self.buf[i].(*interface{}))
}

func (self *QueryResult) getColumn(column string) interface{} {
	return *(self.values[column])
}

func (self *QueryResult) BoolOf(i int) (bool, bool) {
	v := self.getOf(i)
	if v == nil {
		return false, false
	}
	return any2Bool(v)
}

func (self *QueryResult) BoolOfColumn(column string) (bool, bool) {
	v := self.getColumn(column)
	if v == nil {
		return false, false
	}
	return any2Bool(v)
}

func (self *QueryResult) IntOf(i int) (int64, bool) {
	v := self.getOf(i)
	if v == nil {
		return 0, false
	}
	return any2Int(v)
}

func (self *QueryResult) IntOfColumn(column string) (int64, bool) {
	v := self.getColumn(column)
	if v == nil {
		return 0, false
	}
	return any2Int(v)
}

func (self *QueryResult) FloatOf(i int) (float64, bool) {
	v := self.getOf(i)
	if v == nil {
		return 0, false
	}
	return any2Float(v)
}

func (self *QueryResult) FloatOfColumn(column string) (float64, bool) {
	v := self.getColumn(column)
	if v == nil {
		return 0, false
	}
	return any2Float(v)
}

func (self *QueryResult) BytesOf(i int) ([]byte, bool) {
	v := self.getOf(i)
	if v == nil {
		return nil, false
	}
	return any2Bytes(v)
}

func (self *QueryResult) BytesOfColumn(column string) ([]byte, bool) {
	v := self.getColumn(column)
	if v == nil {
		return nil, false
	}
	return any2Bytes(v)
}

func (self *QueryResult) StringOf(i int) (string, bool) {
	v := self.getOf(i)
	if v == nil {
		return "", false
	}
	return any2String(v)
}

func (self *QueryResult) StringOfColumn(column string) (string, bool) {
	v := self.getColumn(column)
	if v == nil {
		return "", false
	}
	return any2String(v)
}

func (self *QueryResult) ValueOf(i int) interface{} {
	return self.getOf(i)
}

func (self *QueryResult) ValueOfColumn(column string) interface{} {
	return self.getColumn(column)
}
