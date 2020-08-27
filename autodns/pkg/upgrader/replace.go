package upgrader

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/lolPants/autodns/autodns/pkg/logger"
)

func executablePath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		logger.Stderr.Printf(1, "failed to get current executable path, error: `%s`\n", err.Error())
		return "", err
	}

	resolved, err := filepath.EvalSymlinks(exe)
	if err != nil {
		logger.Stderr.Printf(1, "resolve symlinks in executable path, error: `%s`\n", err.Error())
		return "", err
	}

	return resolved, nil
}

func replaceLinux(data io.Reader, path string) error {
	logger.Stdout.Printf(2, "deleting current binary `%s`\n", path)
	err := os.Remove(path)
	if err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}

	logger.Stdout.Printf(2, "writing new binary to `%s`\n", path)
	defer out.Close()
	io.Copy(out, data)

	logger.Stdout.Println(2, "changing binary permissions to 755")
	err = os.Chmod(path, 0755)
	if err != nil {
		return err
	}

	return nil
}

func replaceWindows(data io.Reader, path string) error {
	oldPath := path + "-old"
	err := os.Rename(path, oldPath)
	logger.Stdout.Printf(2, "moving binary `%s` to `%s`\n", path, oldPath)

	out, err := os.Create(path)
	if err != nil {
		return err
	}

	logger.Stdout.Printf(2, "writing new binary to `%s`\n", path)
	defer out.Close()
	io.Copy(out, data)

	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// Replace replaces the current executable with an `io.Reader`
func Replace(data io.Reader) error {
	path, err := executablePath()
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "linux":
		return replaceLinux(data, path)
	case "windows":
		return replaceWindows(data, path)
	default:
		return errors.New("unsupported os")
	}
}

// Cleanup cleans up old artifacts left over from the upgrade process
func Cleanup() error {
	path, err := executablePath()
	if err != nil {
		return err
	}

	oldPath := path + "-old"
	if fileExists(oldPath) {
		logger.Stdout.Println(2, "cleaning up old version at `"+oldPath+"`")
		return os.Remove(oldPath)
	}

	return nil
}
