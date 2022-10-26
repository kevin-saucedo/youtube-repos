package connection

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var (
	Conn *sql.DB
	once sync.Once
)

const layer = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

func InitConn(host, port, user, password, dbname string) (err error) {
	once.Do(func() {
		dsn := fmt.Sprintf(layer, host, port, user, password, dbname)
		Conn, err = sql.Open("postgres", dsn)
		if err != nil {
			return
		}
		err = Conn.Ping()
	})
	return
}
