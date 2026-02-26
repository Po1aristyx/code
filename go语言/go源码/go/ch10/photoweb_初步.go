package main

import (
	"net/http"
	"io"
	"log"
	"os"
	"fmt"
	"mime"
	"path"
	"io/ioutil"
	"strings"
)

const UPLOAD_DIR = "./public/upload"

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/", listHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("uploadHandler----------")
	if r.Method == "GET" {
		fmt.Println("get===========")
		content := `
<!DOCTYPE html>
<html>

	<head>
		<meta charset="utf-8">
		<title>用户注册</title>
		<style type="text/css">
			form#reg {
				width:700px;
				font-size: 14px;
				margin:0 auto;
			}
			
			form#reg fieldset {
				border: dashed 1px gray;
			}
			
			form#reg fieldset legend {
				margin: 10px auto;
			}
			
			form#reg div.input-group {
				width:100%;
				height:45px;
				line-height: 45px;
				border-bottom:solid 1px #E2E2E2;
			}
			
			form#reg div.input-group:hover {
				background-color: #E2E2E2;
			}
			
			form#reg div.input-group label {
				display: inline-block;
				width:25%;
				text-align: right;
				margin-right: 25px;
			}
		</style>
	</head>

	<body>
		<form id="reg"  action="/upload" method="post" enctype="multipart/form-data">
			<fieldset id="">
				<legend>用户注册</legend>

				<!--文本输入框-->
				<div class="input-group">
					<label>用户名：</label>
					<input type="text" value="" name="username" placeholder="请输入6位中文用户名"  />
				</div>
	
				<!--文件域-->
				<div class="input-group">
					<label>上传照片：</label>
					<input type="file" name="uploadfile"/>
				</div>

				<!--提交按钮-->
				<div class="input-group">
					<label></label>
					<input type="reset" value="重置" />
					<input type="submit" value="注册" />
				</div>
			</fieldset>
		</form>
	</body>

</html>
`
		io.WriteString(w, content)
		return
	}

	if r.Method == "POST" {
		fmt.Println("post===========")
		srcFile, fileHeader, err := r.FormFile("uploadfile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := fileHeader.Filename
		defer srcFile.Close()

		destFile, err := os.Create(UPLOAD_DIR + "/" + filename)

		fmt.Println(srcFile, destFile, filename)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listHandler----------")
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	listHtml := ""
	for _, fileInfo := range fileInfoArr {
		fileName := fileInfo.Name()
		if !strings.HasPrefix(fileName , ".") {
			listHtml += fmt.Sprintf("<li><a href='/view?id=%s'>%s</a></li>", fileName , fileName)
		}
	}
	io.WriteString(w, "<body><ul>"+listHtml+"</ul></body>")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("viewHandler----------")
	fileId := r.FormValue("id")
	filePath := UPLOAD_DIR + "/" + fileId

	mimetype := mime.TypeByExtension(path.Ext(filePath))
	fmt.Println(mimetype)

	w.Header().Set("Content-Type", mimetype)
	http.ServeFile(w, r, filePath)
}

//func isExists(filePath string) bool {
//	_, err := os.Stat(filePath)
//	if err == nil {
//		return true
//	}
//	return os.IsExist(err)
//}
