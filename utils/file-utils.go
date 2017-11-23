package utils

import (
	"os"
	"path/filepath"
	"io"
	"archive/zip"
	log "github.com/sirupsen/logrus"
	"strings"
	"io/ioutil"
	"crypto/md5"
	"github.com/pkg/errors"
	"encoding/hex"
	"net/http"
)

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Debug("Creating directory: ", dir)
		err = os.MkdirAll(dir, 0755)
		return err
	}
	return nil
}

func Unzip(src, folderName, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		if folderName == "" || strings.HasPrefix(f.Name, folderName) {

			filenames = append(filenames, fpath)

			if f.FileInfo().IsDir() {

				// Make Folder
				os.MkdirAll(fpath, os.ModePerm)

			} else {

				// Make File
				var fdir string
				if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
					fdir = fpath[:lastIndex]
				}

				err = os.MkdirAll(fdir, os.ModePerm)
				if err != nil {
					log.Fatal(err)
					return filenames, err
				}
				f, err := os.OpenFile(
					fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					return filenames, err
				}
				defer f.Close()

				_, err = io.Copy(f, rc)
				if err != nil {
					return filenames, err
				}

			}
		}
	}
	return filenames, nil
}

// Copies file source to destination dest.
func CopyFile(source string, dest string) (err error) {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}

	}

	return
}

func FileExists(fileLocation string) (exists bool) {
	exists = false
	if _, err := os.Stat(fileLocation); !os.IsNotExist(err) {
		exists = true
	}
	return
}

// Recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return &CustomError{"Source is not a directory"}
	}

	// ensure dest dir does not already exist

	_, err = os.Open(dest)
	if !os.IsNotExist(err) {
		return &CustomError{"Destination already exists"}
	}

	// create dest dir

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(source)

	for _, entry := range entries {

		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = CopyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// A struct for returning custom error messages
type CustomError struct {
	What string
}

// Returns the error message defined in What as a string
func (e *CustomError) Error() string {
	return e.What
}

func GenerateChecksum(fileName string) (string, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return "", errors.Wrap(err, "Unable to generate checksum, open file.")
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", errors.Wrap(err, "Unable to generate checksum, copy?")
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
