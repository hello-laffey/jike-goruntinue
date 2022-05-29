// 主入口

package main

import (
	"database/sql"
	"fmt"
	"handlerror/internal/service/dao"
	"os"

	"github.com/pkg/errors"
)

// main 主函数
// 程序的顶部，使用%+v把堆栈打清楚
func main() {
	if err := dao.QuerySql(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("Warning: %v\n", err)
			return
		} else {
			fmt.Printf("FATAL: %+v\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("Process finish ok\n")
	return
}
