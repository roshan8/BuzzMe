package store

import (
	"buzzme/pkg/errors"
	"buzzme/schema"
	"fmt"
)

// UserStore implements the cities interface
type UserStore struct {
	*Conn
}

// NewUserStore ...
func NewUserStore(st *Conn) *UserStore {
	cs := &UserStore{st}
	go cs.createTableIfNotExists()
	return cs
}

func (cs *UserStore) createTableIfNotExists() {
	if !cs.DB.HasTable(&schema.User{}) {
		if err := cs.DB.CreateTable(&schema.User{}).Error; err != nil {
			fmt.Println(err)
		}
	}

	go cs.createIndexesIfNotExists()
}

func (cs *UserStore) createIndexesIfNotExists() {
	scope := cs.DB.NewScope(&schema.User{})
	commonIndexes := getCommonIndexes(scope.TableName())
	for k, v := range commonIndexes {
		if !scope.Dialect().HasIndex(scope.TableName(), k) {
			err := cs.DB.Model(&schema.User{}).AddIndex(k, v).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	uniqueIndexes := map[string][]string{
		"idx_users_name": []string{"name"},
	}
	for k, v := range uniqueIndexes {
		if !scope.Dialect().HasIndex(scope.TableName(), k) {
			if err := cs.DB.Model(&schema.User{}).AddUniqueIndex(k, v...).Error; err != nil {
				fmt.Println(err)
			}
		}
	}
}

// All returns all the users
func (cs *UserStore) All() ([]*schema.User, *errors.AppError) {
	var users []*schema.User
	if err := cs.DB.Find(&users).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return users, nil
}

// Create a new User
func (cs *UserStore) Create(req *schema.UserReq) (*schema.User, *errors.AppError) {
	if recordExists("users", fmt.Sprintf("email='%s'", req.Email)) {
		return nil, errors.BadRequest("User name alreay registered")
	}

	User := &schema.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
	}
	if err := cs.DB.Save(User).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return User, nil
}

// // GetByID returns the matched record for the given id
// func (cs *UserStore) GetByID(UserID uint) (*schema.User, *errors.AppError) {
// 	var User schema.User
// 	if err := cs.DB.First(&User, "id=? and deleted=?", UserID, false).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, errors.BadRequest("invalid User id").AddDebug(err)
// 		}
// 		return nil, errors.InternalServerStd().AddDebug(err)
// 	}

// 	return &User, nil
// }

// // Update the User name, lat, lon values
// func (cs *UserStore) Update(User *schema.User, update *schema.User) (*schema.User, *errors.AppError) {
// 	if err := cs.DB.Model(User).Updates(update).Error; err != nil {
// 		return nil, errors.InternalServerStd().AddDebug(err)
// 	}

// 	return User, nil
// }

// // Delete soft deletes the User for the given id
// func (cs *UserStore) Delete(UserID uint) *errors.AppError {
// 	if err := cs.DB.Delete(&schema.User{}, "id=?", UserID).Error; err != nil {
// 		return errors.InternalServerStd().AddDebug(err)
// 	}

// 	return nil
// }
