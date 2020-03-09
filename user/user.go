package user

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Note: struct fields must be public in order for unmarshal to (Variable name should start with Capital letter)
type UserDetails struct {
	Username string `yaml:"username"`
	Email    string `yaml:"email"`
	Phone    string `yaml:"phone"`
}

type Users struct {
	UserConfigs []UserDetails `yaml:"Users"`
}

func InitUsers() {

	users := Users{}

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
