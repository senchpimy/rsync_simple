package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const path = "log";

type noName struct{
  Date string
  Dirs []string
}

func newStr() noName{
  p:=noName{Date:"", Dirs: []string{}}
  p.ReadLines()
  return p
}

func main() {
  server:=newStr()
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  	switch r.Method{
  		case "GET":
  			tpl, err := template.ParseFiles("index.html")
  			if err != nil {
  				http.Error(w, err.Error(), http.StatusInternalServerError)
  				return
  			}
  			if err != nil {return}
  	
  			err = tpl.Execute(w, server)
  			if err != nil {
  				http.Error(w, err.Error(), http.StatusInternalServerError)
  				return
  			}
  		case "POST":
  			fmt.Println("Mensaje Recibido")
  			message:=r.FormValue("message")
        server.Dirs = append(server.Dirs, message)
  		  server.Save()
  			http.Redirect(w, r, "/", http.StatusSeeOther)
  	}
	})

        log.Print("Listening on :3001..")
        err := http.ListenAndServe(":3001", nil)
        if err != nil {
                log.Fatal(err)
        }
}

func (this *noName)Save() error {
  // get file size
  file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
  if err != nil {
          fmt.Println(err)
          return err
  }
  defer file.Close()

  // write to the file
  for _,newdata:= range this.Dirs{
    _, err = file.WriteString(newdata+"\n")
    if err != nil {
            fmt.Println(err)
            return err
    }
  }
  return nil
}

func (this *noName)ReadLines() error {
  info, err := os.Stat(path)
  if err != nil {
          return err
  }
  this.Date = info.ModTime().Format(time.UnixDate)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	//var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		this.Dirs = append(this.Dirs, scanner.Text())
	}
	return scanner.Err()
}
