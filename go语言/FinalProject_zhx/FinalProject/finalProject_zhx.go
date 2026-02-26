package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

const (
	LISTDIR       = 0x0001
	PUBLIC_DIR    = "./public/"
	UPLOAD_DIR    = "./upload/"
	TEMPLATES_DIR = "./views/"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/public", safeHandler(publicHandler))
	r.HandleFunc("/list", safeHandler(listHandler))
	r.HandleFunc("/upload", safeHandler(uploadHandler))
	r.HandleFunc("/view", safeHandler(viewHandler))
	r.HandleFunc("/score", safeHandler(scoreHandleFunc))
	r.HandleFunc("/login", safeHandler(loginHandlerFunc))

	//restful api /article
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	r.HandleFunc("/articles", returnAllArticles)
	r.HandleFunc("/article/{id}", returnSingleArticle)
	r.HandleFunc("/article", createNewArticle).Methods("POST")
	r.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	r.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

	r.HandleFunc("/citys", getCitys)

	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/world?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = http.ListenAndServe(":8082", r)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	file := PUBLIC_DIR + r.URL.Path[len(PUBLIC_DIR)-1:]
	if (0 & LISTDIR) == 0 {
		if exists := isExists(file); !exists {
			http.NotFound(w, r)
			return
		}
	}
	http.ServeFile(w, r, file)
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

func loginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginHandler----------")
	if r.Method == "GET" {
		fmt.Println("get login: login page")
		//返回登陆界面
		err := renderHtml(w, "login", nil)
		showError(err, w)
		return

	}

	if r.Method == "POST" {
		fmt.Println("post login:")

		username := r.FormValue("username")
		pwd := r.FormValue("password")
		if len(username) < 4 || len(username) > 10 {
			w.Write([]byte("用户名不符合规范"))
			return
		}
		if len(pwd) < 4 || len(pwd) > 10 {
			w.Write([]byte("密码不符合规范"))
			return
		}
		//josn msg
		m := map[string][]string{
			"api":     {"register"},
			"message": {"user register success"},
		}
		if data, err := json.MarshalIndent(m, "", " "); err == nil {
			fmt.Fprintf(w, "%s\n", data)
		}
		return
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

type User struct {
	Name   string
	Age    int
	Gender string
	Score  map[string]float64
}

func scoreHandleFunc(w http.ResponseWriter, r *http.Request) {
	// 2. 解析模板
	t, err := template.ParseFiles(TEMPLATES_DIR + "score.html")
	if err != nil {
		fmt.Println("template parsefile failed, err:", err)
		return
	}
	// 3.渲染模板
	user := User{
		Name:   "zhx",
		Age:    20,
		Gender: "female",
		Score: map[string]float64{
			"yuwen":  98,
			"shuxue": 100,
			"yingyu": 94,
		},
	}
	t.Execute(w, user)
}
