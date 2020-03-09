package main

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

func main() {

	users := Users{}

	UsersYamlFile, err := ioutil.ReadFile("users.yaml")
	fmt.Printf("%s", UsersYamlFile)

	if err != nil {
		log.Printf("Enable to fetch users.yaml  #%v ", err)
	}
	err = yaml.Unmarshal([]byte(UsersYamlFile), &users)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Printf("--- t:\n%v\n\n", users)

	// err := yaml.Unmarshal([]byte(data), &user)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- t:\n%v\n\n", t)

	// d, err := yaml.Marshal(&t)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- t dump:\n%s\n\n", string(d))

	// m := make(map[interface{}]interface{})

	// err = yaml.Unmarshal([]byte(data), &m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m:\n%v\n\n", m)

	// d, err = yaml.Marshal(&m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
