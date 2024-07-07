package env

type env struct {
	DB_CLUSTER 	string
	DB_PORT		int
	DB_KEYSPACE	string
	PORT		int
}


var myEnv env

func GetEnv() (*env)  {
	return &myEnv
}