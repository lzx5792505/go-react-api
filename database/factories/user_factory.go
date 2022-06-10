package factories

import (
	"liu/app/models/user"
	"time"

	"github.com/bxcodec/faker/v3"
)

func MakeUsers(times int) []user.User {
	var objs []user.User

	faker.SetGenerateUniqueValues(true)
	for i := 0; i < times; i++ {
		model := user.User{
			Name:        faker.Username(),
			User:        faker.Username(),
			LastLoginAt: time.Now(),
			Password:    "$2a$14$FUoopLlZcIEOvu90twOwfexA/mvFE.nICU3mauyYCv5AJB052inxa",
		}
		objs = append(objs, model)
	}
	return objs
}
