package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 定义模板函数
var templates = template.Must(template.ParseGlob("templates/*.html"))

// 主页数据
type HomePageData struct {
	CompanyName string
	Slogan      string
	Products    []string
}

// 初始化产品数据
var products = []string{
	"Product 1 - Lorem ipsum dolor sit amet",
	"Product 2 - Consectetur adipiscing elit",
	"Product 3 - Sed do eiusmod tempor incididunt",
}

func main() {
	http.HandleFunc("/", homeHandler)                                                          // 主页
	http.HandleFunc("/about", aboutHandler)                                                    // 关于我们
	http.HandleFunc("/products", productsHandler)                                              // 产品页面
	http.HandleFunc("/contact", contactHandler)                                                // 联系我们
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 静态文件

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// 主页处理函数
func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := HomePageData{
		CompanyName: "Awesome Company",
		Slogan:      "Empowering innovation through technology",
		Products:    products,
	}
	templates.ExecuteTemplate(w, "home.html", data)
}

// 关于我们处理函数
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "about.html", nil)
}

// 产品页面处理函数
func productsHandler(w http.ResponseWriter, r *http.Request) {
	data := HomePageData{
		Products: products,
	}
	templates.ExecuteTemplate(w, "products.html", data)
}

// 联系我们处理函数
func contactHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contact.html", nil)
}
