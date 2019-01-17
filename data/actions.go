package data

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func cp(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}

func PathExists(dirName string) (b bool) {
	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		b = true
	} else {
		b = false
	}
	return
}

func CopyLocalFiles(file YamlFile, srcFolder string)  {
	dstFolder := file.HomeFolder
	if !PathExists(dstFolder) {
		_ = os.MkdirAll(dstFolder, 0755)
	}
	copyCommand := exec.Command("/bin/cp", "-Rv", srcFolder, dstFolder)
	out, err := copyCommand.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(out))

}
