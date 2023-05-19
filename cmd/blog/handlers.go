package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Header             []headerdata
	PostsHeader        []postsheaderdata
	Menu               []menudata
	FeaturedPostsTitle string
	FeaturedPosts      []featuredPosts
	MostRecentTitle    string
	MostRecent         []mostRecent
	Footer             []footerdata
}

type postPage struct {
	HeaderPost []headerpostdata
	Post       []postdata
	Footer     []footerdata
}

type headerdata struct {
	BackroundHeader string
	HeaderTitle     []headertitledata
}

type headertitledata struct {
	Escape string
	Nav    []navdata
}

type headerpostdata struct {
	Escape string
	Nav    []navdata
}

type postdata struct {
	Title    string `db:"title"`
	Subtitle string `db:"subtitle"`
	Image    string `db:"image_url"`
	Content  string `db:"content"`
}

type textdata struct {
	First  string
	Second string
	Third  string
	Fourth string
}

type navdata struct {
	First  string
	Second string
	Third  string
	Fourth string
}

type postsheaderdata struct {
	Title    string
	Subtitle string
	Button   string
}

type menudata struct {
	BackroundMenu string
	MenuTitle     []menutitledata
}

type menutitledata struct {
	First  string
	Second string
	Third  string
	Fourth string
	Fiveth string
	Sixth  string
}

type featuredPosts struct {
	PostID      string `db:"post_id"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	Authorurl   string `db:"author_url"`
	Publishdate string `db:"publish_date"`
	Imageurl    string `db:"image_url"`
	PostURL     string
}

type mostRecent struct {
	PostID      string `db:"post_id"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	Authorurl   string `db:"author_url"`
	Publishdate string `db:"publish_date"`
	Imageurl    string `db:"image_url"`
	PostURL     string
}

type footerdata struct {
	Background string
	Title      string
	Button     string
	Rectangl   string
	Bottom     []bottomdata
}

type loginpage struct {
	Header []headerlogindata
	Main   []mainlogindata
}

type headerlogindata struct {
	Logo  string
	Title string
}

type mainlogindata struct {
	Title  string
	Email  string
	Pass   string
	Button string
}

type adminpage struct {
	Header   []headeradmindata
	MainTop  []maintopdata
	MainInfo []maininfodata
	Content  []contentdata
}

type headeradmindata struct {
	Logo      string
	Avatar    string
	ImageExit string
}

type maintopdata struct {
	Title    string
	Subtitle string
	Button   string
}

type maininfodata struct {
	Title   string
	Fields  []fieldsdata
	Preview []previewdata
}

type fieldsdata struct {
	Title          string
	Description    string
	AuthorName     string
	AuthorPhoto    string
	AuthorPhotoURL string
	Upload         string
	Date           string
	TitleImage     string
	BigImageURL    string
	SmallImageURL  string
	BigNote        string
	SmallNote      string
}

type previewdata struct {
	Article  []articledata
	PostCard []postcarddata
}

type articledata struct {
	Label    string
	FrameURL string
	Title    string
	Subtitle string
	Imageurl string
}

type postcarddata struct {
	Label          string
	FrameURL       string
	Imageurl       string
	Title          string
	Subtitle       string
	AuthorPhotoURL string
	AuthorName     string
	Data           string
}

type contentdata struct {
	Title   string
	Comment string
}

type bottomdata struct {
	Escape string
	Nav    []navdata
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	return func(w http.ResponseWriter, r *http.Request) {
		featuredposts, err := featuredposts(db)
		if err != nil {
			http.Error(w, "Error1", 500)
			log.Println(err)
			return
		}

		mostrecent, err := mostrecent(db)
		if err != nil {
			http.Error(w, "Error2", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошибки в консоль
			return                                      // Выполнение ф-ии
		}

		data := indexPage{
			Header:             header(),
			PostsHeader:        postsheader(),
			Menu:               menu(),
			FeaturedPostsTitle: "Featured Posts",
			FeaturedPosts:      featuredposts,
			MostRecentTitle:    "Most Recent",
			MostRecent:         mostrecent,
			Footer:             footer(),
		}

		err = ts.Execute(w, data) // Запускаем шаблонизатор для вывода шаблона в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"]

		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid order id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Order not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html") // Второстепенная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
			return                                      // Bыполнение ф-ии
		}

		data := postPage{
			HeaderPost: headerpost(),
			Post:       post,
			Footer:     footer(),
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/login.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := loginpage{
		Header: headerlogin(),
		Main:   mainlogin(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func header() []headerdata {
	return []headerdata{
		{
			BackroundHeader: "../static/images/head.png",
			HeaderTitle:     headertitle(),
		},
	}
}

func headertitle() []headertitledata {
	return []headertitledata{
		{
			Escape: "../static/images/Escape1.svg",
			Nav:    nav(),
		},
	}
}

func headerpost() []headerpostdata {
	return []headerpostdata{
		{
			Escape: "../static/images/Escape2.svg",
			Nav:    nav(),
		},
	}
}

func postByID(db *sqlx.DB, postID int) ([]postdata, error) {
	const query = `
		SELECT
		  title,
		  subtitle,
		  image_url,
		  content
		FROM
		  post
	    WHERE
		  post_id = ?
	`

	var post []postdata

	err := db.Select(&post, query, postID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func nav() []navdata {
	return []navdata{
		{
			First:  "HOME",
			Second: "CATEGORIES",
			Third:  "ABOUT",
			Fourth: "CONTACT",
		},
	}
}

func postsheader() []postsheaderdata {
	return []postsheaderdata{
		{
			Title:    "Let's do it together",
			Subtitle: "We travel the world in search of stories. Come along the ride",
			Button:   "View Latest Posts",
		},
	}
}

func menu() []menudata {
	return []menudata{
		{
			BackroundMenu: "../static/images/Rect.png",
			MenuTitle:     menutitle(),
		},
	}
}

func menutitle() []menutitledata {
	return []menutitledata{
		{
			First:  "Nature",
			Second: "Photography",
			Third:  "Relaxation",
			Fourth: "Vacation",
			Fiveth: "Travel",
			Sixth:  "Adventure",
		},
	}
}

func featuredposts(db *sqlx.DB) ([]featuredPosts, error) {
	const query = `
		SELECT
		  post_id,
		  title,
		  subtitle,
		  author,
		  author_url,
		  publish_date,
		  image_url
		FROM
		  post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts
	var featuredPosts []featuredPosts // Заранее объявляем массив с результирующей информацией

	err := db.Select(&featuredPosts, query) // Делаем запрос в базу данных
	if err != nil {                         // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	for _, post := range featuredPosts {
		post.PostURL = "/post/" + post.PostID
	}

	fmt.Println(featuredPosts)

	return featuredPosts, nil
}

func mostrecent(db *sqlx.DB) ([]mostRecent, error) {
	const query = `
		SELECT
		  post_id,
		  title,
		  subtitle,
		  author,
		  author_url,
		  publish_date,
		  image_url
		FROM
		  post
		WHERE featured = 0
	` // Составляем SQL-запрос для получения записей для секции most-posts
	var mostrecent []mostRecent // Заранее объявляем массив с результирующей информацией

	err := db.Select(&mostrecent, query) // Делаем запрос в базу данных
	if err != nil {                      // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	for _, post := range mostrecent {
		post.PostURL = "/post/" + post.PostID
	}

	fmt.Println(mostrecent)

	return mostrecent, nil

}

func footer() []footerdata {
	return []footerdata{
		{
			Background: "../static/images/footer.png",
			Title:      "Stay in Touch",
			Button:     "Sumbit",
			Rectangl:   "../static/images/Rectangl.png",
			Bottom:     bottom(),
		},
	}
}

func bottom() []bottomdata {
	return []bottomdata{
		{
			Escape: "../static/images/Escape1.svg",
			Nav:    nav(),
		},
	}
}

func headerlogin() []headerlogindata {
	return []headerlogindata{
		{
			Logo:  "../static/svg_files/Logo Inversed.svg",
			Title: "Log in to start creating",
		},
	}
}

func mainlogin() []mainlogindata {
	return []mainlogindata{
		{
			Title:  "Log In",
			Email:  "Email",
			Pass:   "Password",
			Button: "Log In",
		},
	}
}
