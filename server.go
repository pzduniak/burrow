package burrow

import (
	"github.com/pzduniak/graval"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type PomFTP struct {
	HomePath string
	Login    func(string, string) bool
}

func (driver *PomFTP) Authenticate(user string, pass string) bool {
	return driver.Login(user, pass)
}

func (driver *PomFTP) Bytes(path string) (bytes int) {
	path = filepath.Join(driver.HomePath, path)
	file, err := os.Stat(path)

	if err != nil {
		log.Printf("No such file or directory: %s", path)
		return -1
	}

	return int(file.Size())
}

func (driver *PomFTP) ModifiedTime(path string) (time.Time, error) {
	path = filepath.Join(driver.HomePath, path)
	file, err := os.Stat(path)

	if err != nil {
		log.Printf("No such file or directory: %s", path)
		return time.Now(), err
	}

	return file.ModTime(), nil
}

func (driver *PomFTP) ChangeDir(path string) bool {
	path = filepath.Join(driver.HomePath, path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("No such file or directory: %s", path)
		return false
	}

	return true
}

func (driver *PomFTP) DirContents(path string) []os.FileInfo {
	path = filepath.Join(driver.HomePath, path)
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Print("Contents listing error: %s", err.Error())
		return []os.FileInfo{}
	}

	return files
}

func (driver *PomFTP) DeleteDir(path string) bool {
	path = filepath.Join(driver.HomePath, path)

	file, err := os.Stat(path)
	if err != nil {
		log.Printf("Delete directory error: %s", err.Error())
		return false
	}

	if !file.IsDir() {
		return false
	}

	err = os.Remove(path)

	if err != nil {
		log.Printf("Delete directory error: %s", err.Error())
		return false
	}

	return true
}

func (driver *PomFTP) DeleteFile(path string) bool {
	path = filepath.Join(driver.HomePath, path)

	file, err := os.Stat(path)
	if err != nil {
		log.Printf("Delete file error: %s", err.Error())
		return false
	}

	if file.IsDir() {
		return false
	}

	err = os.Remove(path)

	if err != nil {
		log.Printf("Delete file error: %s", err.Error())
		return false
	}

	return true
}

func (driver *PomFTP) Rename(fromPath string, toPath string) bool {
	fromPath = filepath.Join(driver.HomePath, fromPath)
	toPath = filepath.Join(driver.HomePath, toPath)

	err := os.Rename(fromPath, toPath)
	if err != nil {
		return false
	}

	return true
}

func (driver *PomFTP) MakeDir(path string) bool {
	path = filepath.Join(driver.HomePath, path)
	err := os.Mkdir(path, 0644)
	if err != nil {
		log.Printf("mkdir error: %s", err.Error())
		return false
	}
	return true
}

func (driver *PomFTP) GetFile(path string, w io.Writer) bool {
	path = filepath.Join(driver.HomePath, path)

	fi, err := os.Open(path)
	if err != nil {
		log.Printf("Error while reading data: %s", err.Error())
		return false
	}

	defer func() {
		if err := fi.Close(); err != nil {
			log.Printf("Error while closing file: %s", err.Error())
		}
	}()

	err = CopyData(fi, w)
	if err != nil {
		log.Printf("Error while writing data to buffer: %s", err.Error())
		return false
	}

	return true
}

func (driver *PomFTP) PutFile(path string, data io.Reader) bool {
	path = filepath.Join(driver.HomePath, path)

	fi, err := os.Create(path)
	if err != nil {
		log.Printf("Error while writing data: %s", err.Error())
		return false
	}

	defer func() {
		if err := fi.Close(); err != nil {
			log.Printf("Error while closing file: %s", err.Error())
		}
	}()

	err = CopyData(data, fi)
	if err != nil {
		log.Printf("Error while writing data: %s", err.Error())
		return false
	}

	return true
}

type PomFTPFactory struct {
	HomePath string
	Login    func(string, string) bool
}

func (factory *PomFTPFactory) NewDriver() (graval.FTPDriver, error) {
	return &PomFTP{factory.HomePath, factory.Login}, nil
}
