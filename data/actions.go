package data

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"io"
	"log"
	"os"
	"os/exec"
	"text/template"
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

func ExportServiceFiles(file YamlFile)  {
	box := packr.NewBox("./templates")
	svcFile, err := box.FindString("systemd.txt")
	if err == nil {
		tmpl, _ := template.New("systemD").Parse(svcFile)

		_ = tmpl.Execute(os.Stdout, file)
	} else {
		log.Fatal(err)
	}

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
