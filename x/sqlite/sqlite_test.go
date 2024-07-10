package x

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSqliteTablesStructList(t *testing.T) {
	// 打开SQLite数据库文件
	db, err := sql.Open("sqlite3", "sqlite_v1.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 验证数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 执行查询以获取所有表的名称
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// 处理查询结果
	fmt.Println("Tables in the database:")
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		getSqliteTableStruct(db, tableName)
	}
}

func getSqliteTableStruct(db *sql.DB, tableName string) {

	// 定义结构用于存储表结构信息
	type TableInfo struct {
		CID          int
		Name         string
		Type         string
		NotNull      int
		DefaultValue sql.NullString
		PrimaryKey   int
	}
	// 执行查询
	rows, err := db.Query("PRAGMA table_info(" + tableName + ")")
	if err != nil {
		log.Fatal(err)
	}
	// 处理查询结果
	var tableInfoList []TableInfo
	for rows.Next() {
		var info TableInfo
		err := rows.Scan(&info.CID, &info.Name, &info.Type, &info.NotNull, &info.DefaultValue, &info.PrimaryKey)
		if err != nil {
			log.Fatal(err)
		}
		tableInfoList = append(tableInfoList, info)
	}

	// 输出表结构信息
	fmt.Printf("Table Structure for %s:\n", tableName)
	fmt.Printf("%-4s %-20s %-10s %-8s %-20s %-10s\n", "CID", "Name", "Type", "NotNull", "DefaultValue", "PrimaryKey")
	for _, info := range tableInfoList {
		fmt.Printf("%-4d %-20s %-10s %-8d %-20s %-10d\n", info.CID, info.Name, info.Type, info.NotNull, info.DefaultValue.String, info.PrimaryKey)
	}

}

func TestSqliteQyuery(t *testing.T) {
	// 打开SQLite数据库文件
	db, err := sql.Open("sqlite3", "sqlite_v1.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 验证数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 执行查询
	rows, err := db.Query("SELECT * FROM cluster_nodes where status=1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		var column1, column2, column3, column4 string
		err := rows.Scan(&column1, &column2, &column3, &column4)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Column1: %s, Column2: %s,Column3: %s, Column4: %s\n", column1, column2, column3, column4)
	}

	// 检查查询过程中是否出现错误
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
