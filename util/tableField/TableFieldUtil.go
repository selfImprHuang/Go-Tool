/*
 *  @Author : huangzj
 *  @Time : 2020/4/27 11:43
 *  @Description：
 */

package tableField

import (
	"database/sql"
	"log"
	"strings"
)

/*
 * @param
 * @return
 * @description 在连接数据库的情况下进行数据库表的字段信息查询
 * 	            这边的sql查询语句是：show full columns from [数据库名].[表名]
 */
func FindColumnMessage(dbName string, tableName string, db *sql.DB) ([]ColumnMessage, error) {
	var cms []ColumnMessage

	row, err := db.Query(strings.Join([]string{"show full columns from ", dbName, ".", tableName}, ""))
	if err != nil {
		log.Printf("数据库查询出错：" + err.Error())
		return nil, err
	}

	for row.Next() {
		var cm ColumnMessage
		err := row.Scan(&cm.Field, &cm.Type, &cm.Collation, &cm.Null, &cm.Key, &cm.Default, &cm.Extra, &cm.Privileges, &cm.Comment)
		if err != nil {
			log.Printf("字段赋值错误了" + err.Error())
			return nil, err
		}
		cms = append(cms, cm)
	}
	return cms, nil
}
