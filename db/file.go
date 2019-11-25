package db


import (
	mydb "../db/mysql"
	"fmt"
)

// OnFileUploadFinished : 文件上传完成，保存meta
func OnFileUploadFinished(filehash string, filename string,
	filesize int64, fileaddr string) bool {
		//准备数据
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file (`file_sha1`,`file_name`,`file_size`," +
			"`file_addr`,`status`) values (?,?,?,?,1)")
	if err != nil {
		fmt.Println("数据准备失败, err:" + err.Error())
		return false
	}
	defer stmt.Close()

	//执行数据 插入四个数据
	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	//返回影响的结果行数
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("hash值为:%s 的文件已经更新", filehash)
		}
		return true
	}
	return false
}