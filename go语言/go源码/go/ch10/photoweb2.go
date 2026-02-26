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
	"html/template"
	"strings"
)

const (
	LISTDIR       = 0X0001
	PUBLIC_DIR    = "./public/"
	UPLOAD_DIR    = "./public/upload/"
	TEMPLATES_DIR = "./views/"
)

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/public/", PUBLIC_DIR, 0)
	mux.HandleFunc("/list", safeHandler(listHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	err := http.ListenAndServe(":8082", mux)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & LISTDIR) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("uploadHandler----------")
	if r.Method == "GET" {
		fmt.Println("get===========")
		err := renderHtml(w, "upload", nil)
		showError(err, w)
		return
	}

	if r.Method == "POST" {
		fmt.Println("post===========")
		srcFile, fileHeader, err := r.FormFile("uploadfile")
		showError(err, w)

		filename := fileHeader.Filename
		defer srcFile.Close()

		destFile, err := os.Create(UPLOAD_DIR + filename)

		fmt.Println(srcFile, destFile, filename)

		showError(err, w)

		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		showError(err, w)

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listHandler----------")
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	showError(err, w)
	//check(err)

	locals := make(map[string]interface{})
	files := []string{}
	for _, fileInfo := range fileInfoArr {
		fileName := fileInfo.Name()
		if !strings.HasPrefix(fileName, ".") {
			files = append(files, fileName)
		}
	}
	locals["files"] = files
	fmt.Println("==========", files)
	err = renderHtml(w, "list", locals)
	showError(err, w)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("viewHandler----------")
	fileId := r.FormValue("id")
	filePath := UPLOAD_DIR + fileId

	mimetype := mime.TypeByExtension(path.Ext(filePath))
	fmt.Println(mimetype)

	w.Header().Set("Content-Type", mimetype)
	http.ServeFile(w, r, filePath)
}

func isExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func showError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//func check(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

func renderHtml(w http.ResponseWriter, tmplName string, locals map[string]interface{}) (err error) {
	tmpl, err := template.ParseFiles(TEMPLATES_DIR + tmplName + ".html")
	if err != nil {
		return
	}
	err = tmpl.Execute(w, locals)
	return
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {

			if _, ok := recover().(error); ok {
				//http.Error(w , e.Error(),http.StatusInternalServerError)
				renderHtml(w, "err", nil)
			}
		}()
		fn(w, r)
	}
}
