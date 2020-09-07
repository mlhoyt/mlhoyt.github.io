package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		postEntry, err := NewEntryFromFile(path)
		if err != nil {
			log.Println(fmt.Sprintf("[ERROR] failed reading posts file -- error:%v file:%s", err, path))
			return err
		}
		log.Println(fmt.Sprintf("[INFO] successfully read posts file -- data:%+v", postEntry))

		if err := postEntry.WriteToFile(os.Args[2]); err != nil {
			log.Println(fmt.Sprintf("[ERROR] failed writing content file -- error:%v", err))
			return err
		}

		return nil
	})
	if err != nil {
		log.Println(fmt.Sprintf("[ERROR] failed walking posts dir -- error:%v dir:%s", err, os.Args[1]))
	}
}

type entry struct {
	title   string
	date    *time.Time
	tags    []string
	content []string
}

type scanState int

const (
	scanStateInit scanState = 1 + iota
	scanStateHeader
	scanStateContent
)

func NewEntryFromFile(path string) (*entry, error) {
	fileHandle, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	scannerState := scanStateInit
	postEntry := entry{}

	reHeaderSeparator := regexp.MustCompile(`^---\s*$`)
	reHeaderTitle := regexp.MustCompile(`^title:\s*\"([^\"]+)\"\s*$`)
	reHeaderDate := regexp.MustCompile(`^date:\s*(\d{4}-\d{2}-\d{2})\s+(\d{2}:\d{2}:\d{2})\s+([\-\+]\d{2}):?(\d{2})\s*$`)
	reHeaderTags := regexp.MustCompile(`^categories:\s*(.*)\s*$`)

	for fileScanner.Scan() {
		switch scannerState {
		case scanStateInit:
			if reHeaderSeparator.MatchString(fileScanner.Text()) {
				scannerState = scanStateHeader
			}
		case scanStateHeader:
			switch {
			case reHeaderSeparator.MatchString(fileScanner.Text()):
				scannerState = scanStateContent
			case reHeaderTitle.MatchString(fileScanner.Text()):
				matches := reHeaderTitle.FindStringSubmatch(fileScanner.Text())
				postEntry.title = matches[1]
			case reHeaderDate.MatchString(fileScanner.Text()):
				matches := reHeaderDate.FindStringSubmatch(fileScanner.Text())
				t, err := time.Parse("2006-01-02 15:04:05 -07:00", matches[1]+" "+matches[2]+" "+matches[3]+":"+matches[4])
				if err != nil {
					return nil, err
				}
				postEntry.date = &t
			case reHeaderTags.MatchString(fileScanner.Text()):
				matches := reHeaderTags.FindStringSubmatch(fileScanner.Text())
				postEntry.tags = strings.Split(matches[1], " ")
			}
		case scanStateContent:
			postEntry.content = append(postEntry.content, fileScanner.Text())
		default:
			return nil, errors.New("unexpected scanner state")
		}
	}

	if err := postEntry.Validate(); err != nil {
		return nil, errors.New("[ERROR] " + err.Error() + " path:" + path)
	}

	return &postEntry, nil
}

func (u entry) Validate() error {
	invalidFields := []string{}

	if u.title == "" {
		invalidFields = append(invalidFields, "title")
	}
	if u.date == nil {
		invalidFields = append(invalidFields, "date")
	}
	if u.content == nil {
		invalidFields = append(invalidFields, "content")
	}

	if len(invalidFields) > 0 {
		return errors.New("entry object contains invalid fields: (" + strings.Join(invalidFields, ", ") + ")")
	}

	return nil
}

func (u entry) WriteToFile(path string) error {
	filename := path + "/" + strconv.FormatInt(u.date.Unix(), 10) + "/entry.md"
	log.Println(fmt.Sprintf("[DEBUG]: writing entry to file -- path:" + filename))

	if _, err := os.Stat(filename); err == nil {
		return errors.New("cannot write entry file because it already exists -- file:" + filename)
	}

	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return err
	}

	fileHandle, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	fmt.Fprintf(fileHandle, "---\n")
	fmt.Fprintf(fileHandle, "title: %s\n", u.title)
	fmt.Fprintf(fileHandle, "date: %s\n", u.date.Format(time.RFC3339))
	fmt.Fprintf(fileHandle, "tags: %s\n", strings.Join(u.tags, " "))
	fmt.Fprintf(fileHandle, "---\n")
	for _, line := range u.content {
		fmt.Fprintf(fileHandle, "%s\n", line)
	}

	return nil
}

// ---
// layout: post
// title: "theagileadmin: CNCF and K8s 101"
// date: 2018-01-26 04:43:27 -0800
// categories: devops k8s
// ---
// [http://theagileadmin.com/2018/01/26/cncf-and-k8s-101s/](http://theagileadmin.com/2018/01/26/cncf-and-k8s-101s/)
