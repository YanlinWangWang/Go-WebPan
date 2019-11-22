package handler

import (
	"fmt"
	"io"
	"io/ioutil" //IO读取单元
	"net/http"
	"os"
)

// UploadHandler ： 处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "服务器错误")
			return
		}
		io.WriteString(w, string(data))
		// 另一种返回方式:
		// 动态文件使用http.HandleFunc设置，静态文件使用到http.FileServer设置(见main.go)
		// 所以直接redirect到http.FileServer所配置的url
		// http.Redirect(w, r, "/static/view/index.html",  http.StatusFound)
	} else if r.Method == "POST" {
		// 接收文件流及存储到本地目录
		file, head, err := r.FormFile("file")//Fromfile文件
		if err != nil {
			fmt.Printf("没收到数据, err:%s\n", err.Error())
			return
		}
		defer file.Close()//类似与堆栈 代表在执行前肯定关闭文件

		//创造新的文件
		newFile, err := os.Create("/tmp/"+head.Filename)//注意os直接向系统根目录写入东西 /tmp文件夹在系统目录下 与/home文件夹平级
		if err != nil {
			fmt.Printf("无法创建文件, err:%s\n", err.Error())
			return
		}
		defer newFile.Close()

		//拷贝文件
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("文件拷贝失败, err:%s\n", err.Error())
			return
		}

		//重定向至陈宫页面
		http.Redirect(w,r,"/file/upload/success",http.StatusFound)//htttp302代表重定向
	}
}

// UploadSucHandler : 上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "上传成功！")
}
