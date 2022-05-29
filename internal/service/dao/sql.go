// dao层sql操作

package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

// QuerySql demo返回sql错误
// 和sql标准库协作，用 errors.Warp保存堆栈信息
func QuerySql() error {
	// xxx 省略真实查询操作，直接返回sql.ErrNoRows错误
	err := sql.ErrNoRows
	return errors.Wrapf(err, "Query sql failed")
}
