package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func handleJsonArray(jsonfile string) []map[string]interface{}{
	var myjson []map[string]interface{}
	file,err := ioutil.ReadFile(jsonfile)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(file,&myjson)
	return myjson
}

func addParameter(parameters []string) []string {
	p := getUserInput()
	parameters = append(parameters,p)
	return parameters
}


func getUserInput() string{
	reader := bufio.NewReader(os.Stdin)
	text,_ := reader.ReadString('\n')
	text = strings.Replace(text,"\n","",-1)
	return text
}
