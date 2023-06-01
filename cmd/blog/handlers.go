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

	"encoding/base64"
	"encoding/json"
	"io"
	"os"
	"strings"
)

type indexPage struct {
	Header             []headerdata
	Menu               []menudata
	FeaturedPostsTitle string
	FeaturedPosts      []*featuredPosts
	MostRecentTitle    string
	MostRecent         []*mostRecent
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
	PostsHeader     []postsheaderdata
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

type bottomdata struct {
	Escape string
	Nav    []navdata
}

type loginpage struct {
	Background string
	Header     []headerlogindata
	Main       []mainlogindata
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

type AdminPage struct {
	AdminHeader []adminheaderdata
	MainInfo    []maininfodata
	FullPage    []fullpagedata
}

type adminheaderdata struct {
	ImageLogo1    string
	ImageLogo2    string
	FirstCharName string
	ImageExit     string
}

type maininfodata struct {
	Title    string
	Subtitle string
	Button   string
}

type UserRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
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
			http.Error(w, "Invalid post id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Post not found", 404)
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
		Background: "../static/images/login_background.png",
		Header:     headerlogin(),
		Main:       mainlogin(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/admin.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := AdminPage{
		AdminHeader: adminheader(),
		MainInfo:    maininfo(),
		FullPage:    fullpage(),
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func adminheader() []adminheaderdata {
	return []adminheaderdata{
		{
			ImageLogo1:    "../static/images/escapeforadmin.svg",
			ImageLogo2:    "../static/images/authorforadmin.svg",
			FirstCharName: "K",
			ImageExit:     "../static/images/log-out.svg",
		},
	}
}

func maininfo() []maininfodata {
	return []maininfodata{
		{
			Title:    "New Post",
			Subtitle: "Fill out the form bellow and publish your article",
			Button:   "Publish",
		},
	}
}

type fullpagedata struct {
	MainTitle       string
	Title           string
	Description     string
	AuthorName      string
	TextAuthorphoto string
	Authorphoto     string
	TextDate        string
	ImageCamera     string
	Upload          string
	TextImage       string

	PreviewArticle1 string
	PreviewImage1   string

	PreviewArticle2 string
	PreviewImage2   string

	TitleInPreview           string
	SubtitleInPreview        string
	PreviewImageChange       string
	PreviewImageChangeAuthor string
	RandomName               string
	RandomDate               string

	FooterTitle    string
	FooterSubtitle string

	ImageMegaCamera string
	ImageMegaTrash  string
}

func fullpage() []fullpagedata {
	return []fullpagedata{
		{
			MainTitle:                "Main Information",
			Title:                    "Title",
			Description:              "Short description",
			AuthorName:               "Author name",
			TextAuthorphoto:          "Author Photo",
			Authorphoto:              "../static/images/photo_icon.svg",
			TextDate:                 "Publish Date",
			ImageCamera:              "../static/images/photo_icon.svg",
			Upload:                   "Upload",
			TextImage:                "Hero Image",
			PreviewArticle1:          "Article preview",
			PreviewImage1:            "../static/images/aritcle_frame.png",
			PreviewArticle2:          "Post card preview",
			PreviewImage2:            "../static/images/post_card_frame.png",
			TitleInPreview:           "New Post",
			SubtitleInPreview:        "Please, enter any description",
			PreviewImageChangeAuthor: "../static/images/blankLogo.png",
			PreviewImageChange:       "../static/images/kek.jpg",
			RandomName:               "Enter author name",
			RandomDate:               "4/19/2023",
			FooterTitle:              "Content",
			FooterSubtitle:           "Post content (plain text)",
			ImageMegaCamera:          "../static/images/camera.png",
			ImageMegaTrash:           "../static/images/trash-2.png",
		},
	}
}

func header() []headerdata {
	return []headerdata{
		{
			BackroundHeader: "../static/images/head.png",
			HeaderTitle:     headertitle(),
			PostsHeader:     postsheader(),
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
	// В SQL-запросе добавились параметры, как в шаблоне. ? означает параметр, который мы передаем в запрос ниже

	var post []postdata

	// Обязательно нужно передать в параметрах orderID
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

func featuredposts(db *sqlx.DB) ([]*featuredPosts, error) {
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
	var featuredPosts []*featuredPosts // Заранее объявляем массив с результирующей информацией

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

func mostrecent(db *sqlx.DB) ([]*mostRecent, error) {
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
	var mostrecent []*mostRecent // Заранее объявляем массив с результирующей информацией

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
			Logo:  "../static/images/Logo Inversed.svg",
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

type createPostRequest struct {
	Title           string `json:"title_g"`
	Description     string `json:"subtitle_g"`
	AuthorName      string `json:"author_name_g"`
	AuthorPhoto     string `json:"author_url_name"`
	AuthorPhotoName string `json:"author_url_name_base64"`
	Date            string `json:"date_g"`
	BigImage        string `json:"big_image_name"`
	BigImageName    string `json:"big_image_name_base64"`
	SmallImage      string `json:"small_image_name"`
	SmallImageName  string `json:"small_image_name_base64"`
	ContentPost     string `json:"text_area_content_g"`
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "1Error", 500)
			log.Println(err.Error())
			return
		}

		var req createPostRequest

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "2Error", 500)
			log.Println(err.Error())
			return
		}

		b64Author := req.AuthorPhotoName[strings.IndexByte(req.AuthorPhotoName, ',')+1:]
		authorImg, err := base64.StdEncoding.DecodeString(b64Author)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileAuthor, err := os.Create("static/images/" + req.AuthorPhoto)
		if err != nil {
			http.Error(w, "file", 500)
			fmt.Println(err.Error())
			return
		}

		_, err = fileAuthor.Write(authorImg)
		if err != nil {
			http.Error(w, "write", 500)
			log.Println(err.Error())
			return
		}

		b64Big := req.BigImageName[strings.IndexByte(req.BigImageName, ',')+1:]
		bigImg, err := base64.StdEncoding.DecodeString(b64Big)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileBig, err := os.Create("static/images/" + req.BigImage)
		if err != nil {
			http.Error(w, "file", 500)
			log.Println(err.Error())
			return
		}

		_, err = fileBig.Write(bigImg)
		if err != nil {
			http.Error(w, "write", 500)
			log.Println(err.Error())
			return
		}

		b64Small := req.SmallImageName[strings.IndexByte(req.SmallImageName, ',')+1:]
		smallImg, err := base64.StdEncoding.DecodeString(b64Small)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileSmall, err := os.Create("static/images/" + req.SmallImage)
		if err != nil {
			http.Error(w, "file", 500)
			log.Println(err.Error())
			return
		}

		_, err = fileSmall.Write(smallImg)
		if err != nil {
			http.Error(w, "write", 500)
			log.Println(err.Error())
			return
		}

		err = saveOrder(db, req)
		if err != nil {
			http.Error(w, "bd", 500)
			log.Println(err.Error())
			return
		}
		return
	}
}

func saveOrder(db *sqlx.DB, req createPostRequest) error {
	const query = `
		INSERT INTO
			post
		(
			title,
			subtitle,
			author,
			author_url,
			publish_date,
			image_url,
			content,
			featured
		)
		VALUES
		(
			?,
			?,
			?,
			CONCAT('../static/images/', ?),
			?,
			CONCAT('../static/images/', ?),
			?,
			?
		)
	`

	_, err := db.Exec(query, req.Title, req.Description, req.AuthorName, req.AuthorPhoto, req.Date, req.SmallImage, req.ContentPost, 0)
	return err
}
