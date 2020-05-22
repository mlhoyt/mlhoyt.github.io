package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

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
				tags := strings.TrimSpace(matches[1])
				if tags != "" {
					postEntry.tags = strings.Split(tags, " ")
				}
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

func (u entry) ToMarkdown() ([]byte, error) {
	var buffer bytes.Buffer
	var err error

	_, err = buffer.WriteString(fmt.Sprintf("# %s\n", u.title))
	if err != nil {
		return nil, err
	}

	_, err = buffer.WriteString(fmt.Sprintf("<p class=\"meta\">%s</p>\n", u.date.Format("2006-01-02")))
	if err != nil {
		return nil, err
	}

	_, err = buffer.WriteString(fmt.Sprintf("<p class=\"meta\">\n"))
	if err != nil {
		return nil, err
	}

	for _, tag := range u.tags {
		_, err = buffer.WriteString(fmt.Sprintf("<span class=\"tag\">%s</span>\n", tag))
		if err != nil {
			return nil, err
		}
	}

	_, err = buffer.WriteString(fmt.Sprintf("</p>\n"))
	if err != nil {
		return nil, err
	}

	_, err = buffer.WriteString(fmt.Sprintf("\n"))
	if err != nil {
		return nil, err
	}

	for _, line := range u.content {
		_, err = buffer.WriteString(fmt.Sprintf("%s\n", line))
		if err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}
