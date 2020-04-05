package schema

// Struct UserDetails holds the all user information
type Incident struct {
	BaseSchema
	CheckName string `json:"check_name"`
	Slug      string `json:"-"`
	State     string `json:"state"`
	Deleted   bool   `json:"deleted" sql:"default:false"`
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
