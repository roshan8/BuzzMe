package schema

type UserDetails struct {
	Username string `yaml:"username"`
	Email    string `yaml:"email"`
	Phone    string `yaml:"phone"`
}

type Users struct {
	UserConfigs []UserDetails `yaml:"Users"`
}
