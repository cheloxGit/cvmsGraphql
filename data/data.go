package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cheloxGit/cvmsGraphql/resolvers"
)

//RetrieveCVMSFromFile func
func RetrieveCVMSFromFile() func() []resolvers.CV {
	return func() []resolvers.CV {
		jsonf, err := os.Open("./data.json")

		if err != nil {
			fmt.Printf("failed to open json file, error: %v", err)
		}

		jsonDataFromFile, _ := ioutil.ReadAll(jsonf)
		defer jsonf.Close()

		var cvmsData []resolvers.CV

		err = json.Unmarshal(jsonDataFromFile, &cvmsData)

		if err != nil {
			fmt.Printf("failed to parse json, error: %v", err)
		}
		return cvmsData
	}
}
