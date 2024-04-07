package reading_files

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
)

func newPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)
	postData := map[string]string{}

Loop:
	for {
		scanner.Scan()
		text := scanner.Text()
		switch {
		case strings.HasPrefix(text, titlePrefix):
			postData["title"] = strings.TrimPrefix(text, titlePrefix)
		case strings.HasPrefix(text, descriptionPrefix):
			postData["description"] = strings.TrimPrefix(text, descriptionPrefix)
		case strings.HasPrefix(text, tagsPrefix):
			postData["tags"] = strings.TrimPrefix(text, tagsPrefix)
		default:
			break Loop
		}
	}

	var body bytes.Buffer
	for scanner.Scan() {
		_, err := fmt.Fprintln(&body, scanner.Text())
		if err != nil {
			return Post{}, err
		}
		postData["body"] = strings.TrimSuffix(body.String(), "\n")
	}

	return Post{
		Title:       postData["title"],
		Description: postData["description"],
		Tags:        getTags(postData["tags"]),
		Body:        postData["body"],
	}, nil
}

func getTags(input string) []string {
	trimmed := strings.ReplaceAll(input, ",", "")
	trimmed = strings.TrimSpace(input)
	if trimmed == "" {
		return []string{}
	}
	return strings.Split(input, ", ")
}
