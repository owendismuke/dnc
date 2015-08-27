package dnc

import (
  "fmt"
  "net/http"
  "html/template"
  "io/ioutil"
  "strings"
  "log"
  "bytes"
)

//Main Function - Listens for Requests
func init() {
  http.HandleFunc("/public/", publicHandler)//http.StripPrefix("/public/", http.FileServer(http.Dir("../public/"))) )
  
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/test", testHandler)
  http.HandleFunc("/view", viewHandler)
  // http.ListenAndServe(":8080", nil)
}

//Memory object to dump data into
var buffer bytes.Buffer

func testHandler(res http.ResponseWriter, req *http.Request) {
  //Handling the request
  
  //Handling the response
  //Writes string to response
  switch req.Method {
    case "POST":  log.Printf("POST from %s", req.RemoteAddr) //Load bytes into buffer as they come in
    default:  fmt.Fprintf(res, string("Response from DNC Web Client\n") )
  }
}


func indexHandler(w http.ResponseWriter, req *http.Request) {
  //Serve /templates/index.html
  log.Println(req.URL.Path)
  log.Println("index Handler called - index.html template should be served")
  templates, err := template.ParseFiles("templates/index.html")
  if err != nil {
    log.Println(err)
    return
  }
  templates.Execute(w, nil)
}

func publicHandler(w http.ResponseWriter, req *http.Request) {
  path := req.URL.Path[1:]
  log.Println("public file requested")
  log.Println(path)
  data, err := ioutil.ReadFile(string(path))

  if err == nil {
    var contentType string
    
    if strings.HasSuffix(path, ".css") {
      contentType = "text/css"
    } else if strings.HasSuffix(path, ".png") {
      contentType = "image/png"
    } else if strings.HasSuffix(path, ".jpg") {
      contentType = "image/jpg"
    } else if strings.HasSuffix(path, ".svg") {
      contentType = "image/svg+xml"
    } else if strings.HasSuffix(path, ".js") {
      contentType = "application/javascript"
    } else {
      contentType = "text/plain"
    }

    w.Header().Add("Content Type", contentType)
    w.Write(data)
  } else {
    w.WriteHeader(404)
    w.Write([]byte("404 no bueno " + http.StatusText(404)))
  }
}

func viewHandler(res http.ResponseWriter, req *http.Request) {
  //Handling the request
    // Basic output
  //Handling the response
  res.Header().Add("Content Type", "text/html")
  // tmpl, err := template.New("index").Parse(doc)
  // if err == nil {
  //   tmpl.Execute(res, "Just the home page")
  // }
  t, _ := template.ParseFiles("templates/view.html")
  t.Execute(res, req)
}