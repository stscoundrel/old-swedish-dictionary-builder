package reader

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func getDictionaryPath(dictionary string) string {
	_, base, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(base)

	return basepath + "/resources/" + dictionary + ".xml"
}

func ReadXmlDictionary(dictionary string) []byte {
	file, err := os.Open(getDictionaryPath(dictionary))

	if err != nil {
		fmt.Println("Could not read dictionary file: ", err)
		os.Exit(1)
	}

	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	return bytes
}
