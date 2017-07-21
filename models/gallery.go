package models

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type Image struct {
	Name       string
	Caption    string
	Categories string
}

func GetImages(max int) []Image {
	files := make([]Image, 0)
	entries, err := ioutil.ReadDir("static/img/fulls")
	if err != nil {
		beego.Critical("can't read images directory")
		return files
	}
	for i, f := range entries {
		name := f.Name()
		caption := "no caption"
		categories := ""
		c, err := ioutil.ReadFile("static/img/captions/" + f.Name() + ".txt")
		if err == nil {
			line := string(c) // categories|caption need to split
			elements := strings.Split(line, "|")
			categories = elements[0]
			caption = strings.TrimSpace(elements[1])
		}
		files = append(files, Image{Name: name, Caption: caption, Categories: categories})
		if max > 0 && i >= max-1 {
			return files
		}
	}
	return files
}
