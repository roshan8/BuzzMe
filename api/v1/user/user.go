package user

import (
	"net/http"

	"buzzme/pkg/errors"
	"buzzme/pkg/respond"
)

// InitUsers fetches and unmarshal the user data from yaml config files
func getAllUsersHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {

	// users := schema.Users{}

	// UsersYamlFile, err := ioutil.ReadFile("config/users.yaml") //Todo: need to fetch an abosolute path later
	// fmt.Printf("%s", UsersYamlFile)

	// if err != nil {
	// 	log.Printf("Enable to fetch users.yaml  #%v ", err)
	// }
	// err = yaml.Unmarshal([]byte(UsersYamlFile), &users)
	// if err != nil {
	// 	log.Fatalf("Unmarshal: %v", err)
	// }

	// fmt.Printf("--- t:\n%v\n\n", users)

	// return nil

	users, err := store.User().All()
	if err != nil {
		return err
	}

	respond.OK(w, users)
	return nil
}
