package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Env struct {
	db *sql.DB
}

func main() {
	for {
		db, err := sql.Open("mysql","go:password@tcp(127.0.0.1:3306)/bdd_passwords")
		if err != nil {
			panic(err.Error())
		}
		env := &Env{db:db}
		menu()
		userInput := getUserInput()
		var parameters []string
		switch userInput {
		case "0":
			os.Exit(0)
		case "1":
			questions := handleJsonArray("genPw.json")
			for _,a := range questions {
				fmt.Println(a["question"])
				parameters = addParameter(parameters)
			}
			fmt.Println("Generating Password...")
			pw := generatePassword(parameters)
			fmt.Println(pw)

		case "2":
			seeDataBase(env.db)
		case "3":
			fmt.Println("Quelle est le site pour lequel vous-voulez ajouter un mot de passe ? ")
			site := getUserInput()
			fmt.Println("Quelle est le mot de passe ? ")
			mdp := getUserInput()
			addCred(env.db,site,mdp)
		case "4":
			fmt.Println("Pour quelle site voulez-vous supprimer le mdp ?")
			site := getUserInput()
			deleteCred(env.db,site)
		default:
			fmt.Println("Une erreur s'est produite: Entrée invalide")


		}
	}
}

