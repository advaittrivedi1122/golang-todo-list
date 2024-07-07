package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/advaittrivedi1122/todolist/env"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

var Session *gocql.Session

// Initialises the ScyllaDb Session with specified configurations.
func Initialise()  {
	env := env.GetEnv()

	// Initialise DB with required configuration
	cluster := gocql.NewCluster(env.DB_CLUSTER)
	cluster.Port = env.DB_PORT
	cluster.Keyspace = env.DB_KEYSPACE
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = time.Second * 10	// 10s timeout

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Unable to create session : %v", err)
	}

	// Make active session available across packages
	Session = session
}

func init() {
	// Load env file & setup env struct for easier access
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(err)
	}
	
	env := env.GetEnv()
	env.PORT, _ = strconv.Atoi(os.Getenv("PORT"))
	env.DB_CLUSTER = os.Getenv("DB_CLUSTER")
	env.DB_PORT, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	env.DB_KEYSPACE = os.Getenv("DB_KEYSPACE")
}