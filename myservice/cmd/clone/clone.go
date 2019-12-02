package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	dirName         string
	codeVersion     string
	codeProjectPath string
)

var console *bufio.Reader

var oldUrl string
var oldName string
var oldDir string
var oldService string

var newUrl string
var newName string
var newDir string
var newService string

func main() {

	doPrompts()

	src := oldDir
	dst := newDir
	// make sure newDir doesn't exists
	oldDir = absDir(".")
	parent := parentDir(oldDir)
	newDir = path.Join(parent, newDir)
	//fmt.Printf("oldDir: %v, newDir: %v\n", oldDir, newDir)
	_, err := os.Stat(newDir)
	if err == nil {
		fmt.Printf("found existing dir: %v\n", newDir)
		os.Exit(1)
	}

	pkgFrom := oldUrl + "/" + src
	pkgTo := newUrl + "/" + dst

	f := []string{pkgFrom, oldName, oldService, src}
	t := []string{pkgTo, newName, newService, dst}

	err = copyDir(oldDir, newDir, f, t)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// rename proto dir in newDir/olf
	p := path.Join(newDir, src)
	q := path.Join(newDir, dst)
	_, err = os.Stat(p)
	if err == nil {
		//fmt.Printf("renaming %v to %v", p, q)
		os.Rename(p, q)
		// rename proto file
		os.Rename(path.Join(q, src+".proto"), path.Join(q, dst+".proto"))
		// remove old generated files
		p = q + "/" + src
		os.Remove(p + ".pb.go")
		os.Remove(p + ".pb.gw.go")
		os.Remove(p + ".swagger.json")
	}
}

func absDir(d string) string {
	wd, _ := os.Getwd()
	p := path.Join(wd, d)
	p = path.Clean(p)
	return p
}

func parentDir(p string) string {
	x := strings.Split(p, "/")
	if len(x) < 2 {
		return p
	}
	return strings.Join(x[:len(x)-1], "/")
}

func prompt(p string) string {
	fmt.Print(p + " > ")
	resp, err := console.ReadString('\n')
	if err != nil {
		panic(err)
	}
	resp = strings.Trim(resp, " \n\t")
	return resp
}

func doPrompts() {
	console = bufio.NewReader(os.Stdin)
	pwd, _ := os.Getwd()
	_, dirName = path.Split(pwd)

	fmt.Printf("clone will determine the current service name based on directory\n")
	fmt.Printf("then will ask for newName service name in the form  c and will copy all files to newservice \n and substitute strings and names accordingly\n")
	// ask for oldName/newName git urls
	oldUrl = prompt("enter oldName import/git repo url (<enter> for github.com/ivost)")
	if oldUrl == "" {
		oldUrl = "github.com/ivost"
	}
	if !strings.Contains(oldUrl, ".") || !strings.Contains(oldUrl, "/") {
		fmt.Printf("Invalid  url %v\n", oldUrl)
		os.Exit(1)
	}
	newUrl = prompt("enter newName git repo url or <enter> to keep " + oldUrl)
	if newUrl == "" {
		newUrl = oldUrl
	}

	oldName = prompt("enter old service name (<enter> for myservice.MyService)")
	if oldName == "" {
		oldName = "myservice.MyService"
	}
	newName = prompt("enter new service name (i.e. newservice.NewService)")
	//fmt.Printf("oldName: %v, newName: %v\n", oldName, newName)
	o1 := strings.Split(oldName, ".")
	if len(o1) != 2 {
		fmt.Printf("Invalid input (no dot) %v\n", oldName)
		os.Exit(1)
	}
	n1 := strings.Split(newName, ".")
	if len(n1) != 2 {
		fmt.Printf("Invalid input (no dot) %v\n", newName)
		os.Exit(1)
	}
	newDir = n1[0]
	newService = n1[1]
	oldDir = o1[0]
	oldService = o1[1]
	_ = oldService
	if oldDir != dirName {
		fmt.Printf("dir name %v should match %v\n", dirName, o1[0])
		os.Exit(1)
	}

	resp := prompt("enter y to clone " + dirName + " to " + newDir + " " + newService)
	//log.Printf("resp: %v", resp)
	if !strings.HasPrefix(resp, "y") {
		os.Exit(1)
	}

}

func copyDir(src string, dst string, from []string, to []string) error {
	//log.Printf("copyDir from: %v, to: %v", src, dst)
	if len(from) != len(to) {
		return fmt.Errorf("different sizes")
	}
	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = copyDir(srcPath, dstPath, from, to)
			if err != nil {
				continue
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				fmt.Printf("skip symlink %s\n", srcPath)
				continue
			}
			if len(from) == 0 && len(to) == 0 {
				err = copyFile(srcPath, dstPath)
			} else {
				err = copyFileSubst(srcPath, dstPath, from, to)
			}
		}
	}
	return err
}

func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

func copyFileSubst(src string, dst string, from []string, to []string) (err error) {
	//fmt.Printf("=== copyFileSubst %v %v %+v %+v\n", src, dst, from, to)
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	content, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	s := string(content)

	for i := 0; i < len(from); i++ {
		s = strings.Replace(s, from[i], to[i], -1)
	}

	err = ioutil.WriteFile(dst, []byte(s), os.ModePerm)
	return err
}
