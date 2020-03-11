package user

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/buzzme/schema"

	"gopkg.in/yaml.v2"
)

// InitUsers fetches and unmarshal the user data from yaml config files
func InitUsers() {

	users := schema.Users{}

	UsersYamlFile, err := ioutil.ReadFile("config/users.yaml") //Todo: need to fetch an abosolute path later
	fmt.Printf("%s", UsersYamlFile)

	if err != nil {
		log.Printf("Enable to fetch users.yaml  #%v ", err)
	}
	err = yaml.Unmarshal([]byte(UsersYamlFile), &users)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Printf("--- t:\n%v\n\n", users)

}
