package file

import (
	"log"
	"os"
	"testing"
)

func TestFile(t *testing.T) {

	var path string = "/Users/alice/gocode/go-learning/file/test1"
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("托尔斯泰发的说法是奋斗\n")
	f.WriteString("agadsfd\n")
	f.WriteString("1233")
	f.WriteString("454546")

}
