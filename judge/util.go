package judge

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func create(path string) (*os.File, error) {
	base := filepath.Dir(path)
	err := os.MkdirAll(base, 0755)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
}

func withTemp(resources []string, f func() error) error {
	// directory setting
	prev, err := filepath.Abs(".")
	if err != nil {
		logrus.WithError(err).Errorf("Failed to get current directory.")
		return err
	}
	dir, err := ioutil.TempDir("", "spica")
	if err != nil {
		logrus.WithError(err).Errorf("Failed to create temporally directory.")
		return err
	}
	for _, v := range resources {
		copy(v, filepath.Join(dir, v))
	}
	defer os.Chdir(prev)
	defer os.RemoveAll(dir)

	err = os.Chdir(dir)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to change working directory.")
		return err
	}
	return f()
}

func withSrc(abs, rel string, f func()) error {
	err := copy(abs, rel)
	if err != nil {
		logrus.WithError(err).Error("Failed to copying src file.")
		return err
	}
	err = os.Setenv("SRC", rel)
	if err != nil {
		logrus.WithError(err).Error("Failed to copying src file.")
		return err
	}
	f()
	return nil
}
func copy(src string, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		logrus.WithField("src", src).WithError(err).Error("Failed to open file")
		return err
	}
	file, err := create(dst)
	if err != nil {
		logrus.WithField("dst", dst).WithError(err).Error("Failed to create file")
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		logrus.WithField("dst", dst).WithError(err).Error("Failed to write to file")
	}
	return err
}

func trimedStr(str string, size int) string {
	if len(str) <= size {
		return str
	}
	return str[0:size] + "..."
}
