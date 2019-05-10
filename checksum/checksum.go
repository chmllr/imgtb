package checksum

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

var (
	jpegRegexp = regexp.MustCompile(`(?i)\.jpe?g`)
	mp4Regexp  = regexp.MustCompile(`(?i)\.mp4`)
)

// Reports walks through the folder structure and returns a mapping
// file path -> md5 hash
func Report(lib string) (res []struct{ Path, Hash string }, err error) {
	err = filepath.Walk(lib, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() || !jpegRegexp.MatchString(path) && !mp4Regexp.MatchString(path) {
			return nil
		}

		h, err := hash(path)
		if err != nil {
			return err
		}
		res = append(res, struct{ Path, Hash string }{path, h})
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(res, func(i, j int) bool { return res[i].Path < res[j].Path })
	return
}

func hash(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	return fmt.Sprintf("%x", md5.Sum(data)), err
}
