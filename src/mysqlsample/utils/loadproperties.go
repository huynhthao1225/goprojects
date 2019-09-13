package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mysqlsample/config"
	"os"
)

// LoadProperties will read application.json file
func LoadProperties() *config.ConnectionProperties {
	// Open our jsonFile
	jsonFile, err := os.Open("E:\\projects\\goprojects\\src\\mysqlsample\\application.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened application.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	
	conProperties := new(config.ConnectionProperties)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'conProperties' which we defined above
	json.Unmarshal(byteValue, &conProperties)
	fmt.Println("I am done with LoadProperties")

	return conProperties

}
