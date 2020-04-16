package user

import (
	"net/http"

	"buzzme/pkg/errors"
	"buzzme/pkg/respond"
	"buzzme/schema"
	"buzzme/utils"
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

// creates new user
func createUserHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	var input schema.UserReq

	if err := utils.Decode(r, &input); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}

	city, err := store.User().Create(&input)
	if err != nil {
		return err
	}

	respond.Created(w, city)
	return nil
}

// Get incident details by ID
// func getIncidentHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
// 	ctx := r.Context()
// 	incident, _ := ctx.Value("incident").(*schema.Incident)

// 	respond.OK(w, incident)
// 	return nil
// }
