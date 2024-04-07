package reading_files

import (
	"reflect"
	"strings"
	"testing"
)

func TestPost(t *testing.T) {

	assertPost := func(t *testing.T, got Post, want Post) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\nwanted: %+v\ngot   : %+v", want, got)
		}
	}

	t.Run("create a new post with all fields", func(t *testing.T) {
		data := `Title: Post 1
Description: description 1
Tags: tdd, go
---
Hello
World!
`
		reader := strings.NewReader(data)

		post, _ := newPost(reader)

		assertPost(t, post, Post{
			Title:       "Post 1",
			Description: "description 1",
			Tags:        []string{"tdd", "go"},
			Body:        "Hello\nWorld!",
		})
	})

	t.Run("create a new post without tags", func(t *testing.T) {
		data := "Title: Post 1\nDescription: description 1\n---\nHello"
		reader := strings.NewReader(data)

		post, _ := newPost(reader)

		assertPost(t, post, Post{Title: "Post 1", Description: "description 1", Tags: []string{}, Body: "Hello"})
	})

	t.Run("create a new post without body", func(t *testing.T) {
		data := "Title: Post 1\nDescription: description 1"
		reader := strings.NewReader(data)

		post, _ := newPost(reader)

		assertPost(t, post, Post{Title: "Post 1", Description: "description 1", Tags: []string{}, Body: ""})
	})

	t.Run("create a post with non-body data in any order", func(t *testing.T) {
		data := "Description: description 1\nTags: tdd, go\nTitle: Post 1"
		reader := strings.NewReader(data)

		post, _ := newPost(reader)

		assertPost(t, post, Post{Title: "Post 1", Description: "description 1", Tags: []string{"tdd", "go"}})
	})

}
