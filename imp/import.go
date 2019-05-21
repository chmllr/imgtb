package imp

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/chmllr/imgtb/index"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

var (
	jpegRegexp = regexp.MustCompile(`(?i)\.jpe?g`)
	mp4Regexp  = regexp.MustCompile(`(?i)\.mp4`)
)

// Import puts every file into the lib with YYYY/MM/DD folder structure
func Import(lib, sourceFolder string) (refs []index.LibRef, err error) {
	exif.RegisterParsers(mknote.All...)

	files, err := ioutil.ReadDir(sourceFolder)
	if err != nil {
		return nil, fmt.Errorf("could't read folder %s: %v", sourceFolder, err)
	}

	log.Println("files found:", len(files))
	folders := map[string][]string{}
	skipped := 0
	errors := 0
	for _, f := range files {
		var err error
		var t time.Time
		fullName := filepath.Join(sourceFolder, f.Name())
		if jpegRegexp.MatchString(fullName) {
			t, err = imgDateTime(fullName)
		} else if mp4Regexp.MatchString(fullName) {
			t, err = mp4DateTime(fullName)
		} else {
			skipped++
			continue
		}
		if err != nil {
			log.Printf("skipping file %s due to errors: %v\n", f.Name(), err)
			errors++
			continue
		}
		folder := t.Format("2006/01/02")
		folders[folder] = append(folders[folder], f.Name())
	}
	log.Println("new folders required:", len(folders))
	imported := 0
	for folder, files := range folders {
		destinationPath := filepath.Join(lib, folder)
		if err := os.MkdirAll(destinationPath, os.ModePerm); err != nil && !os.IsExist(err) {
			log.Printf("couldn't create folder %q: %v\n", destinationPath, err)
			continue
		}
		for _, fileName := range files {
			from := filepath.Join(sourceFolder, fileName)
			to := filepath.Join(destinationPath, fileName)
			err := moveFile(from, to)
			if err != nil {
				log.Printf("couldn't move file from %q to %q: %v\n", from, to, err)
				continue
			}
			info, err := os.Stat(to)
			if err != nil {
				return nil, fmt.Errorf("couldn't open file %s: %v", to, err)
			}
			ref, err := index.NewLibRef(to, info.Size())
			if err != nil {
				return nil, fmt.Errorf("couldn't create libref for %v: %v", info, err)
			}
			refs = append(refs, ref)
			imported++
		}
	}
	log.Printf("files succesfully imported: %d/%d (%d skipped, %d failed)", imported, len(files), skipped, errors)
	return
}

func mp4DateTime(path string) (time.Time, error) {
	_, fileName := filepath.Split(path)
	fNameParts := strings.Split(fileName, "_")
	if len(fNameParts) != 2 {
		return time.Time{}, fmt.Errorf("unexpected filename %q", fileName)
	}
	return time.Parse("20060102", fNameParts[0])
}

func imgDateTime(path string) (time.Time, error) {
	f, err := os.Open(path)
	if err != nil {
		return time.Time{}, err
	}

	x, err := exif.Decode(f)
	if err != nil {
		return time.Time{}, err
	}

	return x.DateTime()
}

func moveFile(from, to string) error {
	stat, err := os.Stat(to)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	destinationNonEmpty := stat != nil

	if destinationNonEmpty {
		data, err := ioutil.ReadFile(from)
		if err != nil {
			return err
		}

		destData, err := ioutil.ReadFile(to)
		if err != nil {
			return fmt.Errorf("file %q exists and couldn't be read: %v", to, err)
		}
		for i := range destData {
			if i >= len(data) || data[i] != destData[i] {
				return fmt.Errorf("file %q already exists and differs from file %q", to, from)
			}
		}
		return fmt.Errorf("file %q is already imported, skipping...", from)
	}

	return os.Rename(from, to)
}
