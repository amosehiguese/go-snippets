package snippet

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	ErrFileNotFound 	= errors.New("file not found")
	ErrCannotLoadFile =	errors.New("unable to load file")
)

type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}

type LocalFile struct {
	Base string
}

func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	var oerr error
	o, err := os.Open(p)
	if err != nil && os.IsNotExist(err) {
		log.Printf("Unable to find %s", path)
		oerr = ErrFileNotFound
	} else if err != nil {
		log.Printf("Error loading file %s, err: %s", path, err)
		oerr = ErrCannotLoadFile
	}
	return o, oerr
}

func (l LocalFile) Save(path string, body io.ReadSeeker) error {
	p := filepath.Join(l.Base, path)
	d := filepath.Dir(p)
	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	return err
}

func CloudLockMain() {
	content := `Lorem ipsum dolor sit amet, consectetur` +
		`adipiscing elit. Donec a diam lectus.Sed sit` +
		`amet ipsum mauris. Maecenascongue ligula ac` +
		`quam viverra nec consectetur ante hendrerit.`

		body := bytes.NewReader([]byte(content))

		store, err := fileStore()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Storing content...")
		err = store.Save("foo/bar", body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Retrieving content...")
		c, err := store.Load("foo/bar")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		o, err := io.ReadAll(c)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(o))

}

func fileStore() (File, error) {
	return &LocalFile{Base: "."}, nil
}
