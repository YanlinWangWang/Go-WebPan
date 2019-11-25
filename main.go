package main

import (
	"fmt"
	"net/http"

	"./handler"
)
//主函数当中有路由信息
func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/success", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("服务器错误%s", err.Error())
	}
}
