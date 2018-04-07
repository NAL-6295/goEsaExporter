package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var endpoint = "https://api.esa.io/v1"
var token = flag.String("token", "", "api  token")
var fromteam = flag.String("team", "", "teamname ***.esa.io")
var rootpath = flag.String("root", "", "rootpath")
var mode = flag.String("mode", "json", "json = jsonfile ,md = markdown file ")
var basepath = ".esa.io"

var ext = ""

//Exists パス存在確認
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func main() {
	flag.Parse()
	fmt.Println("test")

	switch *mode {
	case "md":
		ext = ".md"
	case "json":
		ext = ".json"
	default:
		panic("mode is json|md")
	}

	exportpath := *rootpath + *fromteam + basepath
	if !Exists(exportpath) {
		err := os.Mkdir(exportpath, 0644)
		if err != nil {
			panic(err)
		}
	}

	page := 1
	nextPage := 0

	for {

		nextPage = requestPage(page, exportpath)
		if nextPage == 0 {
			return
		}
		//12秒に1回にすることで、API制限を回避
		time.Sleep(12 * time.Second)
		page++
	}

}

func requestPage(page int, exportpath string) int {
	authorizationValue := "Bearer " + *token
	var url = endpoint + "/teams/" + *fromteam + "/posts?page=" + strconv.Itoa(page)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", authorizationValue)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var posts Posts
	json.NewDecoder(resp.Body).Decode(&posts)

	for _, post := range posts.Posts {
		ToLocal(post, exportpath)
	}
	fmt.Println("page:" + strconv.Itoa(page))

	return posts.NextPage
}

//ToLocal Postをローカルに保存
func ToLocal(post Post, exportpath string) {

	stringReader := strings.NewReader(post.BodyHTML)
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		fmt.Print("url scarapping failed")
		return
	}

	var postFileName = strconv.Itoa(post.Number)
	var postImagePath = filepath.Join(exportpath, postFileName)
	var replaceImagePath = filepath.Join(".", postFileName)
	if !Exists(postImagePath) {
		err := os.Mkdir(postImagePath, 0644)
		if err != nil {
			panic(err)
		}
	}

	var md = post.BodyMd
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("src")

		var savedFileName = DownloadImage(url, postImagePath, replaceImagePath)
		md = strings.Replace(md, url, savedFileName, 1)
	})

	file, err := os.OpenFile(postImagePath+ext, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	switch *mode {
	case "json":
		json.NewEncoder(file).Encode(&post)
	case "md":
		fmt.Fprintln(file, md)
	}
}

//DownloadImage 画像のダウンロード
func DownloadImage(url string, exportpath string, replacePath string) string {
	response, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	var _, fileName = path.Split(url)
	fileName = strings.Replace(fileName, ",", "_", -1)
	var fullName = filepath.Join(exportpath, fileName)
	fmt.Println(fullName)
	file, err := os.OpenFile(fullName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
	return filepath.Join(replacePath, fileName)
}
