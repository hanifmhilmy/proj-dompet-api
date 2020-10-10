package database

import (
	"errors"
	"log"
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	// DBMain main database used for the apps
	DBMain = "dompet_main"
)

// DBConnInterface wrapper database connection struct
type DBConnInterface interface {
	Connect(names []string)
	GetDB(name string) (Client, error)
}

type (
	// Connection store the database connection
	Connection struct {
		conn map[string]replication
	}

	// DBConfig store the database config
	DBConfig struct {
		Conn *Connection
		Name map[string]option
	}

	option struct {
		master      string
		maxlifeconn int64
		maxidle     int
		maxconn     int
	}

	replication struct {
		Master *sqlx.DB
	}
)

// NewDB to create database config
func NewDB(conf config.Config) DBConnInterface {
	mapOpt := make(map[string]option)
	for k := range conf.Database {
		o := option{
			master:      conf.Database[k].Master,
			maxconn:     conf.Database[k].MasterMaxConn,
			maxidle:     conf.Database[k].MasterMaxIdle,
			maxlifeconn: conf.Database[k].MaxLifeConn,
		}
		mapOpt[k] = o
	}

	return &DBConfig{
		Name: mapOpt,
	}
}

// Connect create database connection
func (db *DBConfig) Connect(names []string) {
	mapConn := make(map[string]replication, len(names))
	for _, name := range names {
		dbName, ok := db.Name[name]
		if !ok {
			log.Fatalf("[Database] No config set for [%s]\n", name)
			break
		}
		if dbName.master != "" {
			sqldb, err := sqlx.Connect("postgres", dbName.master)
			if err != nil {
				log.Println("[Database] DB master: ", name)
				log.Fatalln(err)
			}

			// setting up options for db conn
			sqldb.SetMaxOpenConns(dbName.maxconn)
			sqldb.SetMaxIdleConns(dbName.maxidle)
			sqldb.SetConnMaxLifetime(time.Duration(dbName.maxlifeconn) * time.Millisecond)

			mapConn[name] = replication{
				Master: sqldb,
			}
		}
	}
	db.Conn = &Connection{conn: mapConn}
}

// GetDB get database connection
func (db *DBConfig) GetDB(name string) (client Client, err error) {
	conn, ok := db.Conn.conn[name]
	if !ok {
		err = errors.New("map nil")
		log.Printf("[Database] %v, reason: No DB connection with name: %s found!\n", name, err)
	}
	return NewClient(conn.Master), nil
}
