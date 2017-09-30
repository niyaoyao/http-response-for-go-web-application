package main
import (
	"fmt"
	"log"
	"net/http"
	"html/template"
    "encoding/json"
    "path"
)

func headerOnly(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Servier", "Go Web Server from N.Y.")
    w.WriteHeader(200)
}

func plainText(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, welcome to N.Y. space!"))
}

func printHelloText(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Go, NY")
}

type User struct {
    Nick string         `json:"nick"`
    UserId string       `json:"userId"`
    Hobbies []string    `json:"hobbies"`
}


func myInfoJSON(w http.ResponseWriter, r *http.Request) {
    var me = User{
        Nick: "N.Y.",
        UserId: "0000000001",
        Hobbies: []string{"Programming", "Drawing"},
    }

    fmt.Println("nick:" + me.Nick)
    fmt.Println("id:" + me.UserId)

    b, err := json.Marshal(me)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(b)
}


func usersInfoJSON(w http.ResponseWriter, r *http.Request) {
    var me = User{
        Nick: "N.Y.",
        UserId: "0000000001",
        Hobbies: []string{"Programming", "Drawing"},
    }

    p := User{
        Nick: "niyaoyao",
        UserId: "0000000002",
        Hobbies: []string{"Programming", "Drawing"},
    }

    q := User{"nycode", "0000000003", []string{"Programming aa", "bbb"}}

    array := []User{me, p, q}
    b, err := json.Marshal(array)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(b)
}

func fileServer(w http.ResponseWriter, r *http.Request) {

    fp := path.Join(".", "daddy.mp4")
    http.ServeFile(w, r, fp)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
    templateFile, _ := template.ParseFiles("ny-home.html")
    
    profile := User{"niyaoyao", "0000000002", []string{"aaa", "nnn"}}
    if err := templateFile.Execute(w, profile); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func main() {
    fmt.Println("Service Started")

    http.HandleFunc("/", indexPage)
    http.HandleFunc("/header", headerOnly)
    http.HandleFunc("/hello", plainText)
    http.HandleFunc("/printHelloText",printHelloText)

    http.Handle("/server/", http.StripPrefix("/server/", http.FileServer(http.Dir("./"))))
    http.HandleFunc("/video", fileServer)

    http.HandleFunc("/me", myInfoJSON)
    http.HandleFunc("/users", usersInfoJSON)
    
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
