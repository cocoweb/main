/* 
   File      : httpShareWithTrace.go 
   Author    : Mike 
   E-Mail    : Mike_Zhang@live.com 
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getFilelist(path string) string {
	m_files, err := ioutil.ReadDir(path)
	if err != nil {
		//     println( "Get filelist error !" ) 
		return ""
	}
	var strRet string
	for _, f := range m_files {
		//    println(f.Name(),f.IsDir()) 
		if path == "./" {
			strRet += "<p><a href=\"" + path + "" + f.Name() + " \">" + f.Name() + "</a></p>"
		} else {
			strRet += "<p><a href=\"" + path[1:] + "/" + f.Name() + " \">" + f.Name() + "</a></p>"
		}
	}
	return strRet
}

func Handler(w http.ResponseWriter, r *http.Request) {
	println("Request ", r.URL.Path, " from ", r.RemoteAddr)
	//   path := r.URL.Path[1:] 
	path := "." + r.URL.Path
	if path == "./favicon.ico" {
		http.NotFound(w, r)
		return
	}
	if path == "./" || getFilelist(path) != "" {
		fmt.Fprintf(w, "%s", getFilelist(path))
		return
	}
	fin, err := os.Open(path)
	defer fin.Close()
	if err != nil {
		fmt.Fprintf(w, "404 : Not found")
		return
	}
	readLen := 1024 * 1024
	buf := make([]byte, readLen)
	startPos := 0
	println("Transfer file ", path, " ... ")
	for {
		n, err := fin.ReadAt(buf, int64(startPos))
		fmt.Fprintf(w, "%s", buf[:n])
		if 0 == n || err != nil {
			break
		}
		startPos += readLen
	}
}
func main() {
	port := "8080" //Default port  
	if len(os.Args) > 1 {
		port = strings.Join(os.Args[1:2], "")
	}
	http.HandleFunc("/", Handler)
	s := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    1 * time.Hour,
		WriteTimeout:   1 * time.Hour,
		MaxHeaderBytes: (1 << 31) - 1, //Max file size is 2048M 
	}
	println("Listening on port ", port, "...")
	log.Fatal(s.ListenAndServe())
}
