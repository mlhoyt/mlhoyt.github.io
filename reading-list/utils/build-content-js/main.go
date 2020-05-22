package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/russross/blackfriday/v2"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var sourceFlag string
var contentHTMLFlag string
var contentJSFlag string

func init() {
	flag.StringVarP(&sourceFlag, "source", "", "", "content source directory")
	flag.StringVarP(&contentHTMLFlag, "content-html", "", "", "generated content HTML output directory")
	flag.StringVarP(&contentJSFlag, "content-js", "", "", "generated content JS file")
}

func main() {
	flag.Parse()

	entries, err := loadContentSource(sourceFlag)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{"directory": sourceFlag}).Fatal("failed loading content source")
	}

	if err := generateContentHTML(contentHTMLFlag, entries); err != nil {
		log.WithError(err).WithFields(log.Fields{"directory": contentHTMLFlag}).Fatal("failed generated content HTML")
	}

	if err := generateContentJS(contentJSFlag, entries); err != nil {
		log.WithError(err).WithFields(log.Fields{"file": contentJSFlag}).Fatal("failed generated content JS")
	}
}

func loadContentSource(dirPath string) ([]entry, error) {
	entries := []entry{}

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Base(path) != "entry.md" {
			return nil
		}

		postEntry, err := NewEntryFromFile(path)
		if err != nil {
			return errors.New(fmt.Sprintf("failed reading entry file: file=%s : %v", path, err.Error()))
		}
		log.WithFields(log.Fields{"file": postEntry}).Debug("successfully read entry file")

		entries = append(entries, *postEntry)

		return nil
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed walking content source directory: dir=%s : %v", dirPath, err.Error()))
	}

	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].date.Unix() > entries[j].date.Unix()
	})

	return entries, nil
}

func generateContentHTML(dirPath string, entries []entry) error {
	if err := os.RemoveAll("./" + dirPath); err != nil {
		return err
	}

	if err := os.MkdirAll("./"+dirPath, 0755); err != nil {
		return err
	}

	templatePathGlob := filepath.Join("./templates", "*.tmpl")
	tmpl, err := template.New("Entry HTML").ParseGlob(templatePathGlob)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryDir := "./" + dirPath + "/" + strconv.FormatInt(entry.date.Unix(), 10)

		if err := os.MkdirAll(entryDir, 0755); err != nil {
			return err
		}

		entryMarkdown, err := entry.ToMarkdown()
		if err != nil {
			return err
		}

		entryHTML := blackfriday.Run(entryMarkdown)

		var buffer bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buffer, "entryHTML", strings.Split(string(entryHTML), "\n")); err != nil {
			return err
		}

		if err := ioutil.WriteFile(entryDir+"/entry.html", buffer.Bytes(), 0644); err != nil {
			return err
		}
	}

	return nil
}

func generateContentJS(filePath string, entries []entry) error {
	fileHandle, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	fmt.Fprintf(fileHandle, "content = [\n")
	for _, entry := range entries {
		tags := []string{}
		for _, v := range entry.tags {
			tags = append(tags, fmt.Sprintf("%q", v))
		}

		fmt.Fprintf(fileHandle, "  new ContentEntry(%d, \"%s\", [%s]),\n", entry.date.Unix(), entry.title, strings.Join(tags, ", "))
	}
	fmt.Fprintf(fileHandle, "]\n")

	return nil
}
