package sqlconnection

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"server/setting"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
)

func Init(config setting.Config) {
	var err error
	withCA := true

	if config.SSLMode == "false" {
		withCA = false
	}

	DBConn, err = OpenConnection(config, withCA)
	if err != nil {
		panic(err)
	}
}

// OpenConnection opens a connection to the database.
func OpenConnection(config setting.Config, WithCert bool) (*sql.DB, error) {
	if WithCert {
		certBytes, err := base64.StdEncoding.DecodeString(config.CACERTBASE64)
		if err != nil {
			log.Fatalf("unable to read in the cert file: %s", err)
		}

		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(certBytes); !ok {
			log.Fatal("failed-to-parse-sql-ca")
		}

		tlsConfig := &tls.Config{
			RootCAs:            caCertPool,
			InsecureSkipVerify: true,
		}

		if err := mysql.RegisterTLSConfig("custom", tlsConfig); err != nil {
			panic(err)
		}
	}

	for {
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&tls=%s&timeout=10s",
			config.User,
			config.Password,
			config.DatabaseHost,
			config.DatabasePort,
			config.Name,
			config.SSLMode,
		))

		if err != nil {
			PrintInfo(config, err)
		} else {
			err = db.Ping()
			if err != nil {
				PrintInfo(config, err)
			} else {
				log.Printf("Connect To Database %v At Port %v \n", "Mysql", config.DatabaseHost)
				return db, nil
			}
		}
	}
}

func PrintInfo(config setting.Config, err error) {
	log.Println("Exit with error: ", err)
	log.Println("Reconnecting to database with info: ", config.User, " ", config.Password, " ", config.DatabaseHost, " ", config.DatabasePort, " ", config.Name, " ", config.SSLMode)
	time.Sleep(time.Second * 3)
}
