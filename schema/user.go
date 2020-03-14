package schema

type UserDetails struct {
	Username string `yaml:"username"`
	Email    string `yaml:"email"`
	Phone    string `yaml:"phone"`
}

type Users struct {
	UserConfigs []UserDetails `yaml:"Users"`
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
