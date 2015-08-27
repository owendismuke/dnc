package main

import (
  "fmt"
  "net/http"
  "html/template"
  "io/ioutil"
  "strings"
  "log"

)

type MyHandler struct {
}

//Main Function - Listens for Requests
func main() {
  http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("../public/"))) )
  
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/test", testHandler)
  http.ListenAndServe(":8080", nil)
}

func testHandler(res http.ResponseWriter, req *http.Request) {
  //Handling the request
  
  //Handling the response
  //Writes string to response
  fmt.Fprintf(res, string("Response from DNC Web Client\n") )
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
  //Serve /templates/index.html
  log.Println(req.URL.Path)
  templates, err := template.ParseFiles("../templates/index.html")
  if err != nil {
    log.Println(err)
    return
  }
  templates.Execute(w, nil)
}

func publicHandler(w http.ResponseWriter, req *http.Request) {
  path := req.URL.Path
  log.Println(path)
  data, err := ioutil.ReadFile(string(""+path))

  if err ==nil {
    var contentType string
    
    if strings.HasSuffix(path, ".css") {
      contentType = "text/css"
    } else if strings.HasSuffix(path, ".js") {
      contentType = "application/javascript"
    } else if strings.HasSuffix(path, ".html") {
      contentType = "text/html"
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
  t, _ := template.ParseFiles("../templates/index.html")
  t.Execute(res, req)
}