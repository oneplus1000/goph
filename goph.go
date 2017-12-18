package goph

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cgi"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

//Serv php cgi server
type Serv struct {
	WwwRoot string //root of www
	PhpBin  string //path of php-cgi bin
}

//Start start server
func (s *Serv) Start() error {
	http.HandleFunc("/", s.handleFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return errors.Wrap(err, "http.ListenAndServe(\":8080\", nil) fail")
	}
	return nil
}

//Close close server
func (s *Serv) Close() {

}

func (s *Serv) handleFunc(w http.ResponseWriter, r *http.Request) {

	f := filepath.Join(s.WwwRoot, r.URL.Path)
	ext := strings.ToLower(filepath.Ext(f))
	if ext == ".php" {
		//php file
		handler := new(cgi.Handler)
		handler.Path = s.PhpBin
		handler.Env = append(handler.Env, "REQUEST_METHOD="+r.Method)
		handler.Env = append(handler.Env, "REDIRECT_STATUS=CGI")
		handler.Env = append(handler.Env, "SCRIPT_FILENAME="+f)
		handler.ServeHTTP(w, r)
	} else {
		//static file
		data, err := ioutil.ReadFile(f)
		if err != nil {
			log.Printf("file %s not found", f)
			return
		}
		buff := bytes.NewReader(data)
		_, err = io.Copy(w, buff)
		if err != nil {
			log.Printf("io.Copy(w, buff) fail (file:%s)", f)
			return
		}
	}
}
