package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var titleFlag string
var linkFlag string

func init() {
	flag.StringVarP(&titleFlag, "title", "", "untitled", "entry title (default: untitled)")
	flag.StringVarP(&linkFlag, "link", "", "", "URL")
}

func main() {
	flag.Parse()

	if linkFlag == "" {
		log.Fatal("must provide --link=<STRING> option")
	}

	now := time.Now()

	srcPath := filepath.Join("./source", strconv.FormatInt(now.Unix(), 10))
	if err := os.MkdirAll(srcPath, 0755); err != nil {
		log.WithError(err).Fatal("failed creating entry source directory")
	}

	var buffer strings.Builder
	buffer.WriteString("---\n")
	buffer.WriteString(fmt.Sprintf("title: %s\n", titleFlag))
	buffer.WriteString(fmt.Sprintf("date: %s\n", now.Format("2006-01-02T15:04:05-07:00")))
	buffer.WriteString("tags:\n")
	buffer.WriteString("---\n")
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("[%s](%s)\n", titleFlag, linkFlag))

	if err := ioutil.WriteFile(srcPath+"/entry.md", []byte(buffer.String()), 0644); err != nil {
		log.WithError(err).Fatal("failed creating entry markdown file")
	}
}
