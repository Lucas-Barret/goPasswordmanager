package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func generatePassword(parameters []string) interface{} {

	url := "https://api.motdepasse.xyz/create/?include_lowercase"

	if parameters[0] == "y" {
		url = url + "&include_uppercase"
	}

	if parameters[1] == "y" {
		url = url + "&include_digits"
	}

	if parameters[2] == "y" {
		url = url + "&include_special"
	}


	url = url + "&password_length=" + parameters[3] + "&quantity=1"
	resp,err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, _ := io.ReadAll(resp.Body)
	var myjson map[string]interface{}
	json.Unmarshal(body,&myjson)
	return myjson["passwords"]


}


func seeDataBase(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM mymdp")
	var (
		site string
		pass string
	)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&site,&pass)
		fmt.Printf("%s %s\n",site,pass)
	}
}

func addCred(db *sql.DB,site string,mdp string) {
	query := "INSERT INTO mymdp VALUES('"+site+"','"+mdp+"')"
	_,err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func deleteCred(db *sql.DB,site string) {
	query := "DELETE FROM mymdp WHERE site='" +  site + "'"
	_,err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}

}