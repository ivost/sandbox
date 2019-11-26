package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	podName         string
	codeVersion     string
	codeProjectPath string
)

func DoClone() {
	console := bufio.NewReader(os.Stdin)
	pwd, _ := os.Getwd()
	_, podName = path.Split(pwd)

	log.Printf("pod name: %v", podName)
	if podName == "boss_pod_template" {
		log.Panic("run clone in new empty working dir/git repo")
	}
	// cd /tmp && git clone https://github.com/braincorp/boss_pod_template
	tmpPath := "/tmp/boss_pod_template"
	_, err := os.Stat(tmpPath)
	if err != nil {
		log.Panicf("template not found in %v (cd /tmp && git clone https://github.com/braincorp/boss_pod_template)", tmpPath)
	}
	log.Printf("enter y to continue")
	resp, err := console.ReadString('\n')
	if err != nil {
		panic(err)
	}
	resp = strings.Trim(resp, " \n\t")
	//log.Printf("resp: %v", resp)
	if !strings.HasPrefix(resp, "y") {
		os.Exit(1)
	}
	pkgFrom := "github.com/braincorp/boss_pod_template"
	pkgTo := fmt.Sprintf("github.com/braincorp/%s", podName)

	if err := replaceImports(tmpPath, pkgFrom, pkgTo); err != nil {
		log.Fatalf("Error parsing imports: %s", err)
		os.Exit(1)
	}

	from := "boss_pod_template"
	to := podName
	dirs := []string{"pkg", "cmd"}
	for _, dir := range dirs {
		//log.Printf("dir %v", dir)
		err = os.MkdirAll(path.Join(codeProjectPath, dir), os.ModePerm)
		if err != nil {
			log.Fatalf("Error: %s", err)
			os.Exit(1)
		}
		err = copyDir(path.Join(tmpPath, dir), path.Join(codeProjectPath, dir), from, to)
		if err != nil {
			log.Fatalf("Error: %s", err)
			os.Exit(1)
		}
	}

	makeFiles := []string{"Makefile", "go.mod", "go.sum", "LICENSE", "README.md"}
	for _, file := range makeFiles {
		fileContent, err := ioutil.ReadFile(path.Join(tmpPath, file))
		if err != nil {
			log.Printf("File: %v, Error: %s", file, err)
			//os.Exit(1)
			continue
		}
		destFile := file
		newContent := strings.Replace(string(fileContent), from, to, -1)
		err = ioutil.WriteFile(path.Join(codeProjectPath, destFile), []byte(newContent), os.ModePerm)
		if err != nil {
			log.Printf("Error: %s", err)
			//os.Exit(1)
			continue
		}
	}
}

func gitPush() error {
	cmdPush := fmt.Sprintf("git add . && git commit -m \"sync %s\" && git push", codeVersion)
	cmd := exec.Command("sh", "-c", cmdPush)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func replaceImports(projectPath string, pkgFrom string, pkgTo string) error {
	fmt.Printf("replaceImports %v from: %v, to: %v\n", projectPath, pkgFrom, pkgTo)
	regexImport, err := regexp.Compile(`(?s)(import(.*?)\)|import.*$)`)
	if err != nil {
		return err
	}

	regexImportedPackage, err := regexp.Compile(`"(.*?)"`)
	if err != nil {
		return err
	}

	found := []string{}

	err = filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" {
			//log.Printf("READ FILE %v", path)
			bts, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			content := string(bts)
			matches := regexImport.FindAllString(content, -1)
			isExists := false

		isReplaceable:
			for _, each := range matches {
				for _, eachLine := range strings.Split(each, "\n") {
					matchesInline := regexImportedPackage.FindAllString(eachLine, -1)
					if err != nil {
						return err
					}

					for _, eachSubline := range matchesInline {
						if strings.Contains(eachSubline, pkgFrom) {
							isExists = true
							break isReplaceable
						}
					}
				}
			}

			if isExists {
				content = strings.Replace(content, `"`+pkgFrom+`"`, `"`+pkgTo+`"`, -1)
				content = strings.Replace(content, `"`+pkgFrom+`/`, `"`+pkgTo+`/`, -1)
				found = append(found, path)
			}
			//log.Printf("WRITE FILE %v", path)

			err = ioutil.WriteFile(path, []byte(content), info.Mode())
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("ERROR", err.Error())
	}

	if len(found) == 0 {
		log.Println("Nothing replaced")
	} else {
		log.Printf("Go imports total %d file replaced\n", len(found))
	}

	return nil
}

func copyDir(src, dst, from, to string) error {
	//log.Printf("copyDir from: %v, to: %v", src, dst)
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
				log.Printf("skip symlink %s", srcPath)
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

func copyFileSubst(src, dst, from, to string) (err error) {
	//log.Printf("copyFileSubst %v %v %v %v", src, dst, from, to)
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
	s2 := strings.Replace(s, from, to, -1)
	//if strings.Compare(s, s2) != 0 {
	//	log.Printf(s2)
	//}
	err = ioutil.WriteFile(dst, []byte(s2), os.ModePerm)
	return err
}
