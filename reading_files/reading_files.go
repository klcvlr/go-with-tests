package reading_files

import (
	"errors"
	"io/fs"
	"strings"
)

var ErrUnsupportedFileType = errors.New("unsupported file type")

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return []Post{}, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	if !strings.HasSuffix(fileName, ".md") {
		return Post{}, ErrUnsupportedFileType
	}
	file, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer func(file fs.File) { _ = file.Close() }(file)
	return newPost(file)
}
