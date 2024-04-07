package reading_files

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (f StubFailingFS) Open(_ string) (fs.File, error) {
	return nil, errors.New("that's what I do")
}

func TestReadingFiles(t *testing.T) {
	assertPost := func(t *testing.T, got Post, want Post) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\nwanted: %+v\ngot   : %+v", want, got)
		}
	}

	t.Run("get posts from FS", func(t *testing.T) {
		firstBody := `Title: Post 1
Description: description 1`
		secondBody := `Title: Post 2
Description: description 2`
		fileSystem := fstest.MapFS{
			"hello-world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := NewPostsFromFS(fileSystem)
		if err != nil {
			t.Fatal("err")
		}
		assertPost(t, posts[0], Post{Title: "Post 1", Description: "description 1", Tags: []string{}})
		assertPost(t, posts[1], Post{Title: "Post 2", Description: "description 2", Tags: []string{}})
	})

	t.Run("should fail if the file is not a .md file", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"file": {Data: []byte("")},
		}

		_, err := NewPostsFromFS(fileSystem)

		if err == nil {
			t.Fatal("expected an error but got none", err)
		}
		if !errors.Is(err, ErrUnsupportedFileType) {
			t.Errorf("expected an ErrUnsupportedFileType error")
		}
	})

	t.Run("error when reading post", func(t *testing.T) {
		fileSystem := StubFailingFS{}

		_, err := NewPostsFromFS(fileSystem)

		if err == nil {
			t.Fatal("expected an error but didn't get one", err)
		}
	})
}
