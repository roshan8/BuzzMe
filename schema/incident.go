package schema

// Struct UserDetails holds the all user information
type Incident struct {
	BaseSchema
	IncidentName string `json:"check_name"`
	State        string `json:"state"`
}

type IncidentShort struct {
	IncidentName string `json:"check_name"`
	State        string `json:"state"`
}

// Schema describing the new incident creation payload request
type IncidentReq struct {
	IncidentName string `json:"check_name"`
	State        string `json:"state"`
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
