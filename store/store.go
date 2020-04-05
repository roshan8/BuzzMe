package store

import (
	"buzzme/config"
	"buzzme/schema"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// gorm postgres connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbConn *gorm.DB

// Init ...
func Init() {
	db, err := gorm.Open(config.DBDriver, config.DBDataSource)
	if err != nil {
		log.Fatal(err)
	}
	dbConn = db
	db.AutoMigrate(
		&schema.Users{},
		&schema.Incident{},
		//TODO: add escalation policy and other schema
	)
}

// Conn struct holds the store connection
type Conn struct {
	DB           *gorm.DB
	UserConn     User
	IncidentConn Incident
	// TODO: EscalationConn, WebhookConn etc
}

// NewStore inits new store connection
func NewStore() *Conn {
	Init()
	conn := &Conn{
		DB: dbConn,
	}
	conn.UserConn = NewUserStore(conn)
	conn.IncidentConn = NewIncidentStore(conn)
	// TODO: EscalationConn, WebhookConn etc

	return conn
}

// User implements the store interface and it returns the Users interface
func (s *Conn) User() User {
	return s.UserConn
}

// User implements the store interface and it returns the Users interface
func (s *Conn) Incident() Incident {
	return s.IncidentConn
}

func getCommonIndexes(tableName string) map[string]string {
	idx := fmt.Sprintf("idx_%s", tableName)
	return map[string]string{
		fmt.Sprintf("%s_created_at", idx): "created_at",
		fmt.Sprintf("%s_updated_at", idx): "updated_at",
	}
}

// recordExists should check if record is avail or not for particular table
// based on the given condition.
func recordExists(tableName, where string) (exists bool) {
	baseQ := fmt.Sprintf("select 1 from %s where %v", tableName, where)
	dbConn.Raw(fmt.Sprintf("select exists (%v)", baseQ)).Row().Scan(&exists)
	return
}
