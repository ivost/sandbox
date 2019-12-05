package rand

import (
	"io/ioutil"
	"log"

	"github.com/spf13/afero"
)

func runme {
	log.Printf("Hello")
	var fs = afero.NewOsFs()
	f, err := fs.Create("/tmp/foo")
	if err != nil {
		panic(err)
	}
	f.WriteString("Hello\n")
	f.Close()

	if err != nil {
		panic(err)
	}
	f, err = fs.Open("/tmp/foo")
	if err != nil {
		panic(err)
	}
	d, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	log.Printf("%s", d)
}
