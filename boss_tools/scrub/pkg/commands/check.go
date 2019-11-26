package commands

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"scrub/pkg/logutils"
	"strings"
	"sync/atomic"

	"github.com/spf13/cobra"
)

func (e *Executor) initCheck() {
	cmd := "check"
	e.checkCmd = &cobra.Command{
		Use:   cmd,
		Short: cmd,
		Run:   e.executeCheck,
	}
	e.rootCmd.AddCommand(e.checkCmd)
	e.checkCmd.SetOut(logutils.StdOut) // use custom output to properly color it in Windows terminals

}

func (e *Executor) executeCheck(_ *cobra.Command, _ []string) {
	e.Check()
}

func (e *Executor) Check() {
	c := e.cfg.Input

	e.SanityCheck()

	_ = e.Traverse(c.RootDir, CheckVisitor)
	e.log.Infof("FOUND %v file(s) without header",
		e.NumFoundFiles)
	e.log.Infof("checked %v dirs total (%v skipped), %v total files, %v skipped",
		e.NumTotalDirs, e.NumSkippedDirs, e.NumTotalFiles, e.NumSkippedFiles)
	// set exit code to be number of files found
	e.exitCode = int(e.NumFoundFiles)
}

func CheckVisitor(ex *Executor, path string, info os.FileInfo) {
	ex.CheckVisitor(path, info)
}

func (e *Executor) CheckVisitor(path string, info os.FileInfo) {
	//e.log.Infof("++++ File without header: %v", path)
}

func (e *Executor) Traverse(dir string, visitor VisitorFunc) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if dir == path || info == nil || err != nil {
			return nil
		}
		e.IncrementProgress()
		// check for exclusion dirs - .git, vendor etc.
		if info.IsDir() {
			atomic.AddInt64(&e.NumTotalDirs, 1)
			for _, rex := range e.cfg.Input.SkipDirsRex {
				if rex.MatchString(path) {
					//log.Printf("--- skip dir: %v", path)
					atomic.AddInt64(&e.NumSkippedDirs, 1)
					return filepath.SkipDir
				}
			}
			return nil
		}
		// skip very small files
		if info.Size() < MinSize {
			//log.Printf("small size %v", path)
			atomic.AddInt64(&e.NumSkippedFiles, 1)
			return nil
		}
		//log.Printf("reg.file %v", path)
		for _, rex := range e.cfg.Input.SkipFilesRex {
			if rex.MatchString(path) {
				atomic.AddInt64(&e.NumSkippedFiles, 1)
				return nil
			}
		}
		// not in skip list, check inclusion
		for _, rex := range e.cfg.Input.FilesRex {
			//e.log.Debugf("INCLUDE: %v", e.cfg.Input.Files[i])
			if !rex.MatchString(path) {
				// skip if not in files
				//e.log.Debugf("SKIP %v", path)
				atomic.AddInt64(&e.NumSkippedFiles, 1)
				return nil
			}
		}
		// regular file
		atomic.AddInt64(&e.NumTotalFiles, 1)
		if strings.Contains(path, ".git") {
			e.log.Errorf("PATH %v", path)
			panic("WHY")
		}
		fp, err := os.Open(filepath.Clean(path))
		if err != nil {
			if !strings.Contains(err.Error(), "EOF") {
				log.Printf("Error %v reading %v", err, path)
			}
			atomic.AddInt64(&e.NumSkippedFiles, 1)
			return nil
		}
		//nolint
		defer fp.Close()

		// read only head of the file
		len := info.Size()
		if len < 2*e.HeaderLen {
			// short file
			len = 2 * e.HeaderLen
		}
		buf := make([]byte, len)
		_, err = fp.Read(buf)
		if err != nil {
			log.Printf("Error %v reading %v", err, path)
			atomic.AddInt64(&e.NumSkippedFiles, 1)
			return nil
		}
		content := string(buf)
		content = strings.ToLower(content)

		for _, rex := range e.cfg.Input.SkipTextRex {
			if rex.MatchString(content) {
				//log.Printf("--- skip file due to content: %v", path)
				atomic.AddInt64(&e.NumSkippedFiles, 1)
				return nil
			}
		}
		// make sure to close input file before visitor so it can change it if needed
		_ = fp.Close()

		// passed all tests - call visitor
		atomic.AddInt64(&e.NumFoundFiles, 1)
		visitor(e, path, info)
		return nil
	})

	if err != nil {
		fmt.Println("ERROR", err.Error())
	}

	return nil
}

func (e *Executor) SanityCheck() {

	atomic.StoreInt64(&e.NumTotalDirs, 0)
	atomic.StoreInt64(&e.NumTotalFiles, 0)
	atomic.StoreInt64(&e.NumFoundFiles, 0)
	atomic.StoreInt64(&e.NumSkippedDirs, 0)
	atomic.StoreInt64(&e.NumSkippedFiles, 0)

	c := e.cfg.Input
	codePath, _ := filepath.Abs(c.RootDir)
	info, err := os.Stat(codePath)
	if err != nil {
		log.Printf("Error %v - %v", codePath, err)
		os.Exit(1)
	}
	if !info.IsDir() {
		log.Printf("Error %v - not a dir", codePath)
		os.Exit(1)
	}
	p := path.Join(c.HeaderDir, c.HeaderName)
	headerFile, _ := filepath.Abs(p)
	info, err = os.Stat(headerFile)
	if err != nil {
		log.Printf("Error %v - %v", headerFile, err)
		os.Exit(1)
	}
	if info.IsDir() {
		log.Printf("Error %v - is a dir", headerFile)
		os.Exit(1)
	}
}
