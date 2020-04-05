package store

import (
	"buzzme/pkg/errors"
	"buzzme/schema"
	"fmt"
)

// UserStore implements the cities interface
type IncidentStore struct {
	*Conn
}

// NewUserStore ...
func NewIncidentStore(st *Conn) *IncidentStore {
	cs := &IncidentStore{st}
	go cs.createTableIfNotExists()
	return cs
}

func (cs *IncidentStore) createTableIfNotExists() {
	if !cs.DB.HasTable(&schema.Incident{}) {
		if err := cs.DB.CreateTable(&schema.Incident{}).Error; err != nil {
			fmt.Println(err)
		}
	}

	go cs.createIndexesIfNotExists()
}

func (cs *IncidentStore) createIndexesIfNotExists() {
	scope := cs.DB.NewScope(&schema.Incident{})
	commonIndexes := getCommonIndexes(scope.TableName())
	for k, v := range commonIndexes {
		if !scope.Dialect().HasIndex(scope.TableName(), k) {
			err := cs.DB.Model(&schema.Incident{}).AddIndex(k, v).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	uniqueIndexes := map[string][]string{
		"idx_Incidents_name": []string{"name"},
	}
	for k, v := range uniqueIndexes {
		if !scope.Dialect().HasIndex(scope.TableName(), k) {
			if err := cs.DB.Model(&schema.Incident{}).AddUniqueIndex(k, v...).Error; err != nil {
				fmt.Println(err)
			}
		}
	}
}

// All returns all the Incidents
func (cs *IncidentStore) All() ([]*schema.Incident, *errors.AppError) {
	var Incidents []*schema.Incident
	if err := cs.DB.Find(&Incidents).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return Incidents, nil
}

// Create a new Incident
func (cs *IncidentStore) Create(req *schema.IncidentReq) (*schema.Incident, *errors.AppError) {
	if recordExists("Incident", fmt.Sprintf("name='%s'", req.IncidentName)) {
		return nil, errors.BadRequest("city name alreay registered")
	}

	incident := &schema.Incident{
		IncidentName: req.IncidentName,
		State:        req.State,
	}
	if err := cs.DB.Save(incident).Error; err != nil {
		return nil, errors.InternalServerStd().AddDebug(err)
	}

	return incident, nil
}

// // Create a new User
// func (cs *UserStore) Create(req *schema.UserReq) (*schema.User, *errors.AppError) {
// 	if recordExists("cities", fmt.Sprintf("name='%s' and deleted_at=null", req.Name)) {
// 		return nil, errors.BadRequest("User name alreay registered")
// 	}

// 	User := &schema.User{
// 		Name:      req.Name,
// 		Latitude:  req.Latitude,
// 		Longitude: req.Longitude,
// 	}
// 	if err := cs.DB.Save(User).Error; err != nil {
// 		return nil, errors.InternalServerStd().AddDebug(err)
// 	}

// 	return User, nil
// }

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
