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

const path = "log"
const config = "config"

type noName struct {
	Date    string
	Dirs    []string
	Root    string
	Command string
	Error   string
	Warning string
}

func newStr() noName {
	p := noName{Date: "", Dirs: []string{}}
	dat, err := os.ReadFile(config)
	if err != nil {
		p.Error = "Error trying to read config: " + err.Error()
	} else {
    str := string(dat)
    if len(str)==0{
      p.Error="You should add a dir where the files will be stored"
    }else{
      p.Root=str
    }
	}
	p.ReadLines()
	return p
}

func main() {
	server := newStr()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			tpl, err := template.ParseFiles("index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err != nil {
				return
			}

			err = tpl.Execute(w, server)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "POST":
			server.Error = ""
      server.Warning=""
			fmt.Println("Mensaje Recibido")
			message := r.FormValue("message")
			root := r.FormValue("root")
			if len(server.Root) == 0 && len(root) == 0 {
				server.Error = "Please provide a root to where the files will be sync"
			} else if len(root) != 0 {
				if _, err := os.Stat(root); os.IsNotExist(err) {
					server.Error = "Path for root does not exist"
				} else {
					server.Root = root
				}
			}
			if len(message) != 0 {
				server.Dirs = append(server.Dirs, message)
        server.Warning=fmt.Sprintf("You should check that the directory exists with: [ -d %s ] && echo 'the directory exists.'", message)
			}
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

func (this *noName) Save() error {
	// get file size
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	// write to the file
	for _, newdata := range this.Dirs {
		_, err = file.WriteString(newdata + "\n")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	file2, err := os.OpenFile(config, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = file2.WriteString(this.Root)
	defer file2.Close()
	return nil
}

func (this *noName) ReadLines() error {
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
