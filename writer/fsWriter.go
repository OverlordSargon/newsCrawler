package writer

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"../misc"
)

func WriteToFS(outPath string, filename string, entity interface{}) error {
	res, _ := json.Marshal(entity)
	os.MkdirAll(outPath, 0777)
	filepath := path.Join(outPath, filename)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		ioutil.WriteFile(filepath, res, 0777)
	} else {
		misc.LogError(err)
	}
	return nil
}
