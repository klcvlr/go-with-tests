package templating

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"go-with-test/reading_files"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	var aPost = reading_files.Post{
		Title:       "hello world",
		Description: "This is the description",
		Body:        "This is a post",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("converts a post into HTML", func(t *testing.T) {
		var buff bytes.Buffer
		renderer, _ := NewPostRenderer()

		err := renderer.Render(&buff, aPost)
		got := buff.String()

		if err != nil {
			t.Fatal("got an error but didn't want one", err)
		}
		approvals.VerifyString(t, got)
	})
}

func BenchmarkRender(b *testing.B) {
	var aPost = reading_files.Post{
		Title:       "hello world",
		Description: "This is the description",
		Body:        "This is a post",
		Tags:        []string{"go", "tdd"},
	}
	b.ResetTimer()
	renderer, err := NewPostRenderer()
	if err != nil {
		b.Fatal("got an error but didn't want one", err)
	}
	for i := 0; i < b.N; i++ {
		err = renderer.Render(io.Discard, aPost)
		if err != nil {
			b.Fatal("got an error but didn't want one", err)
		}
	}
}
