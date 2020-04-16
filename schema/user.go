package schema

// Struct UserDetails holds the all user information
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// Struct UserReq holds the required fields for new user creation
type UserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// TODO: Needs to figure out the part on how can we validate the fields in nested struct
// Ok implements the Ok interface, it validates user input
// func (c *Users) Ok() error {
// 	switch {
// 	case strings.TrimSpace(c.UserConfigs.Username) == "":
// 		return errors.IsRequiredErr("Username")
// 	}
// 	return nil
// }
