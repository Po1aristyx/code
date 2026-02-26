package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//定义数据库连接信息
type DbConn struct {
	Dsn string //数据库驱动连接字符串
	Db  *sql.DB
}

//user_info表的映射对象
type UserTable struct {
	Uid        int
	Username   string
	Department string
	Created    string
}

func main() {
	var err error
	dbConn := DbConn{
		//Dsn: "root:root@tcp(127.0.0.1:3306)/testdb?charset=utf8",
		Dsn: "root:@tcp(127.0.0.1:3306)/testdb?charset=utf8",
	}
	dbConn.Db, err = sql.Open("mysql", dbConn.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer dbConn.Db.Close()

	//1、测试封装的ExecData()方法
	execData(&dbConn)

	//2、测试封装的PreExecData()方法
	preExecData(&dbConn)

	//3、查询单行数据
	//查询最后一条数据的信息
	result := dbConn.QueryRowData("select * from user_info where uid=(select max(uid) from user_info)")
	fmt.Println(result)

	//4、查询多行数据
	result1 := dbConn.QueryData("select * from user_info where uid<10")
	fmt.Println(len(result1))
	////遍历查询的结果集
	//for k, v := range result1 {
	//	fmt.Println("uid：", k,  v)
	//}

	//5、查询多行数据
	result2 := dbConn.PreQueryData("select * from user_info where uid<? order by uid desc" , 10)
	fmt.Println(len(result2))
	////遍历查询的结果集
	//for k, v := range result2 {
	//	fmt.Println("uid：", k,  v)
	//}

	dbConn.PreQueryData2("select * from user_info where uid<? order by uid desc" , 10)
	fmt.Println(len(result2))
	////遍历查询的结果集
	//for k, v := range result2 {
	//	fmt.Println("uid：", k,  v)
	//}
}

//一、测试封装的ExecData()函数
func execData(dbConn *DbConn) {
	count, id, err := dbConn.ExecData("INSERT user_info(username , departname , created)  VALUES ('Josh','business group','2018-07-3')")
	//count , err := execData("UPDATE user_info SET created='2018-06-30' WHERE uid=14")
	//count , err := execData("DELETE FROM user_info WHERE uid=10")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("受影响行数：", count)
		fmt.Println("新添加数据的id：", id)
	}
}

//二、测试封装的PreExecData()函数
func preExecData(dbConn *DbConn) {
	count, id, err := dbConn.PreExecData("INSERT user_info(username , departname , created)  VALUES (?,?,?)", "Jackson", "Education Department", "2017-10-8")
	//count, id, err := PreExecData("Delete from user_info WHERE uid<?", 4)
	//count, id, err := PreExecData("UPDATE user_info set departname=? WHERE departname = ?", "BC Group", "blockchain")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("受影响行数：", count)
		fmt.Println("新添加数据的id：", id)
	}
}

//1、封装增删改数据的函数，该函数直接使用DB的Exec()方法实现数据操作
func (dbConn *DbConn) ExecData(sqlString string) (count, id int64, err error) {
	result, err := dbConn.Db.Exec(sqlString)
	if err != nil {
		panic(err)
		return
	}
	if id, err = result.LastInsertId(); err != nil {
		panic(err)
		return
	}
	if count, err = result.RowsAffected(); err != nil {
		panic(err)
		return
	}
	return count, id, nil
}

//2、封装增删改数据的函数，该函数使用预编译语句加Exec()方法实现增删改数据
func (dbConn *DbConn) PreExecData(sqlString string, args ...interface{}) (count, id int64, err error) {
	stmt, err := dbConn.Db.Prepare(sqlString)
	defer stmt.Close()
	if err != nil {
		panic(err)
		return
	}
	result, err := stmt.Exec(args ...)
	if err != nil {
		panic(err)
		return
	}
	if id, err = result.LastInsertId(); err != nil {
		panic(err)
		return
	}
	if count, err = result.RowsAffected(); err != nil {
		panic(err)
		return
	}
	return count, id, nil
}

//3、查询当行数据
func (dbConn *DbConn) QueryRowData(sqlString string) (data UserTable) {
	user := new(UserTable)
	err := dbConn.Db.QueryRow(sqlString).Scan(&user.Uid, &user.Username, &user.Department, &user.Created)
	if err != nil {
		panic(err)
		return
	}
	return *user
}

//4、未使用预编译，直接查询多行数据
func (dbConn *DbConn) QueryData(sqlString string) (resultSet map[int]UserTable) {
	rows, err := dbConn.Db.Query(sqlString)
	defer rows.Close()
	if err != nil {
		panic(err)
		return
	}
	resultSet = make(map[int]UserTable)
	user := new(UserTable)
	for rows.Next() {
		err := rows.Scan(&user.Uid, &user.Username, &user.Department, &user.Created)
		if err != nil {
			panic(err)
			continue
		}
		resultSet[user.Uid] = *user
	}
	return resultSet
}

//5、使用预编译语句进行查询多行数据
func (dbConn *DbConn) PreQueryData(sqlString string , args ...interface{}) (resultSet map[int]UserTable) {
	stmt, err := dbConn.Db.Prepare(sqlString)
	defer stmt.Close()
	if err != nil {
		panic(err)
		return
	}
	rows, err := stmt.Query(args ...)
	defer rows.Close()
	if err != nil {
		panic(err)
		return
	}

	resultSet = make(map[int]UserTable)
	user := new(UserTable)
	for rows.Next() {
		err := rows.Scan(&user.Uid, &user.Username, &user.Department, &user.Created)
		if err != nil {
			panic(err)
			continue
		}
		resultSet[user.Uid] = *user
	}
	return resultSet
}


//无返回值，只打印输出，用于测试
func (dbConn *DbConn) PreQueryData2(sqlString string , args ...interface{})  {
	stmt, err := dbConn.Db.Prepare(sqlString)
	defer stmt.Close()
	if err != nil {
		panic(err)
		return
	}
	rows, err := stmt.Query(args ...)
	defer rows.Close()
	if err != nil {
		panic(err)
		return
	}

	user := new(UserTable)
	for rows.Next() {
		err := rows.Scan(&user.Uid, &user.Username, &user.Department, &user.Created)
		if err != nil {
			panic(err)
			continue
		}
		fmt.Println(*user)
	}
}


