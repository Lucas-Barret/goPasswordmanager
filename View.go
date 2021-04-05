package main

import (
	"fmt"
)


type Menu struct {
	Id string `json:"id"`
	Action string `json:"action"`
}


func menu() {
	menuActions := handleJsonArray("menu.json")
	for _,a := range menuActions {
		fmt.Println(a["id"].(string)+"."+a["action"].(string))//a.Id+"."+a.Action)
	}

}

