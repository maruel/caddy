package templates

import (
	"errors"
	"io/ioutil"
	"sync"
	"text/template"
)

type item struct {
	path string
	t    *template.Template
}

type cacheStore struct {
	items map[string]item
	mu    sync.Mutex
}

var cache = cacheStore{items: map[string]*template.Template{}}

func Exec(name string) error {
}

// Add reads in the template with the filename provided. If the file does not exist or is not parsable, it will return an error.
func Add(name, filename string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if i, ok := cache.items[name]; ok {
		if i.path != filename {
			return errors.New("adding same template name with different paths")
		}
		return nil
	}
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	t, err := template.New(name).Parse(string(buf))
	if err != nil {
		return err
	}
	cache.items[name] = item{filepath, t}
	return nil
}
