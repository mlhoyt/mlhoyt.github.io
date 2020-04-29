package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

func main() {
	allEntries, err := scanDir(os.Args[1])
	if err != nil {
		log.Println(fmt.Sprintf("[ERROR] failed walking content dir -- error:%v dir:%s", err, os.Args[1]))
		os.Exit(1)
	}

	sort.SliceStable(allEntries, func(i, j int) bool {
		return allEntries[i].date.Unix() > allEntries[j].date.Unix()
	})

	if err := WriteToContentJS(allEntries); err != nil {
		log.Println(fmt.Sprintf("[ERROR] failed writing content.js -- error:%v", err))
		os.Exit(1)
	}
}

func scanDir(dirPath string) ([]entry, error) {
	allEntries := []entry{}

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Base(path) != "entry.md" {
			return nil
		}

		postEntry, err := NewEntryFromFile(path)
		if err != nil {
			log.Println(fmt.Sprintf("[ERROR] failed reading entry file -- error:%v file:%s", err, path))
			return err
		}
		log.Println(fmt.Sprintf("[INFO] successfully read entry file -- data:%+v", postEntry))

		allEntries = append(allEntries, *postEntry)

		return nil
	})
	if err != nil {
		log.Println(fmt.Sprintf("[ERROR] failed walking content dir -- error:%v dir:%s", err, os.Args[1]))
	}

	return allEntries, err
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
	reHeaderTitle := regexp.MustCompile(`^title:\s*(.*)\s*$`)
	reHeaderDate := regexp.MustCompile(`^date:\s*(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[\-\+]\d{2}:\d{2})\s*$`)
	reHeaderTags := regexp.MustCompile(`^tags:\s*(.*)\s*$`)

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
				t, err := time.Parse("2006-01-02T15:04:05-07:00", matches[1])
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

func WriteToContentJS(entries []entry) error {
	filename := "content.js"
	log.Println(fmt.Sprintf("[DEBUG]: writing file -- path:" + filename))

	fileHandle, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	fmt.Fprintf(fileHandle, "content = [\n")
	for _, entry := range entries {
		fmt.Fprintf(fileHandle, "  new ContentEntry(%d, \"%s\"),\n", entry.date.Unix(), entry.title)
	}
	fmt.Fprintf(fileHandle, "]\n")

	// let tags = {
	//   Keys: function() {
	//     return [
	//       <TAG>,
	//       ...
	//     ]
	//   },
	//   <TAG>: [
	//     <ENTRY-ID>,
	//     ...
	//   ],
	// }

	return nil
}
