package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/advaittrivedi1122/todolist/env"
	"github.com/advaittrivedi1122/todolist/internal/types"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

var Session *gocql.Session
var UsersCount int = 0

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
	} else {
		fmt.Println("\n[ScyllaDb session created successfully]")
	}

	// Counter table for user todos
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v.todos_count (id INT PRIMARY KEY, count INT)`, cluster.Keyspace)
	fmt.Printf("\n[Query] : %v\n", query)
	if err := session.Query(query).Exec(); err != nil {
		log.Fatalf("Unable to create users_count: %v", err)
	} else {
		fmt.Println("\n[Table Created successfully] : todos_count")
	}

	query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v.user_todos (id INT, user_id INT, title TEXT, description TEXT, status TEXT, created TIMESTAMP, updated TIMESTAMP, PRIMARY KEY(user_id, id)) WITH CLUSTERING ORDER BY (id ASC)", env.DB_KEYSPACE)
	fmt.Printf("\n[Query] : %v\n", query)
	if err := session.Query(query).Exec(); err != nil {
		log.Fatalf("Unable to create table todos: %v", err)
	} else {
		fmt.Println("\n[Table Created successfully] : todos")
	}

	// Make active session available across packages
	Session = session
	// Set UsersCount var with value if present in db
	GetUsersCount()
}

func ExecuteQuery(query string) error {
	fmt.Printf("\n[Query] : %v\n", query)

	if err := Session.Query(query).Exec(); err != nil {
		log.Printf("Unable to execute query: %v", err)
		return err
	}

	return nil
}

// Inserts user todos into seperate sorted tables for ASC and DESC filter
func InsertUserTodo(req types.AddTodoRequest) error {
	env := env.GetEnv()
	userTodoId := GetUserTodosCount(req.UserId)+1

	query := fmt.Sprintf(`INSERT INTO %v.user_todos (id, user_id, title, description, status, created, updated) VALUES (%v, %v, '%s', '%s', '%s', toTimestamp(now()), toTimestamp(now()))`, env.DB_KEYSPACE, userTodoId, req.UserId, req.Title, req.Description, req.Status)

	fmt.Printf("\n[Query] : %v\n", query)
	if err := Session.Query(query).Exec(); err != nil {
		return err
	}

	return nil
}

func UpdateUserTodoById(req types.UpdateUserTodoByIdRequest) error {
	env := env.GetEnv()
	var changes string

	if (req.Title != "") {
		changes += fmt.Sprintf("title = '%s', ", req.Title)
	}
	if (req.Description != "") {
		changes += fmt.Sprintf("description = '%s', ", req.Description)
	}
	if (req.Status != "") {
		changes += fmt.Sprintf("status = '%s'", req.Status)
	}

	query := fmt.Sprintf("UPDATE %v.user_todos SET %v, updated = toTimestamp(now()) WHERE user_id = %v AND id = %v", env.DB_KEYSPACE, changes, req.UserId, req.TodoId)
	if err := Session.Query(query).Exec(); err != nil {
		return err
	}

	return nil
}

func DeleteUserTodoById(req types.DeleteUserTodoByIdRequest) error {
	env := env.GetEnv()

	query := fmt.Sprintf("DELETE FROM %v.user_todos WHERE user_id = %v AND id = %v", env.DB_KEYSPACE, req.UserId, req.TodoId)
	if err := Session.Query(query).Exec(); err != nil {
		return err
	}

	return nil
}

func DeleteUserTodos(userId int) error {
	env := env.GetEnv()

	query := fmt.Sprintf("DELETE FROM %v.user_todos WHERE user_id = %v", env.DB_KEYSPACE, userId)
	if err := Session.Query(query).Exec(); err != nil {
		return err
	}

	return nil
}

func GetUsersCount() int {
	env := env.GetEnv()
	var count int = 0;

	query := fmt.Sprintf("SELECT COUNT(*) FROM %v.todos_count", env.DB_KEYSPACE)
	if err := Session.Query(query).Consistency(gocql.One).Scan(&count); err != nil {
		return 0
	}

	UsersCount = count
	return count
}

func GetUserTodosCount(userId int) int {
	env := env.GetEnv()
	var count int = 0;

	query := fmt.Sprintf("SELECT count FROM %v.todos_count WHERE id = %v", env.DB_KEYSPACE, userId)
	if err := Session.Query(query).Consistency(gocql.One).Scan(&count); err != nil {
		return 0
	}

	return count
}

func IncrementUserTodosCount(userId int) error {
	env := env.GetEnv()
	userTodosCount := GetUserTodosCount(userId)

	query := fmt.Sprintf("UPDATE %v.todos_count SET count = %v WHERE id = %v", env.DB_KEYSPACE, userTodosCount+1, userId)
	if err := Session.Query(query).Exec(); err != nil {
		return err
	}

	return nil
}

func DecrementUserTodosCount(userId int) error {
	env := env.GetEnv()
	userTodosCount := GetUserTodosCount(userId)

	if (userTodosCount <= 0) {
		return nil
	}
	query := fmt.Sprintf("UPDATE %v.todos_count SET count = %v WHERE id = %v", env.DB_KEYSPACE, userTodosCount-1, userId)
	if err := Session.Query(query).Exec(); err != nil {
		return err
	}

	return nil
}

func ResetUserTodosCount(userId int) error {
	env := env.GetEnv()
	userTodosCount := GetUserTodosCount(userId)

	if (userTodosCount <= 0) {
		return nil
	}
	query := fmt.Sprintf("UPDATE %v.todos_count SET count = 0 WHERE id = %v", env.DB_KEYSPACE, userId)
	if err := Session.Query(query).Exec(); err != nil {
		return err
	}

	return nil
}

func GetUserTodoById(userId int, id int) (todo types.Todo, err error) {
	env := env.GetEnv()

	query := fmt.Sprintf("SELECT id, user_id, title, description, status, created, updated FROM %v.user_todos WHERE user_id = %v AND id = %v ALLOW FILTERING", env.DB_KEYSPACE, userId, id)

	if err := Session.Query(query).Consistency(gocql.One).Scan(&todo.TodoId, &todo.UserId, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated); err != nil {
		return todo, err
	}
	
	return todo, nil
}

func GetUserTodos(req types.GetUserTodosRequest) (todos []types.Todo, err error) {
	env := env.GetEnv()
	var query string

	if (req.Filter == "" && req.Sort == "") {
		// Without filter and sort
		query = fmt.Sprintf("SELECT id, user_id, title, description, status, created, updated FROM %v.user_todos WHERE user_id = %v ALLOW FILTERING", env.DB_KEYSPACE, req.UserId)
	} else if (req.Filter != "" && req.Sort == "") {
		// With filter, without sort
		query = fmt.Sprintf("SELECT id, user_id, title, description, status, created, updated FROM %v.user_todos WHERE user_id = %v AND status = '%v' ALLOW FILTERING", env.DB_KEYSPACE, req.UserId, req.Filter)
	} else if (req.Filter == "" && req.Sort != "") {
		// Without filter, with sort (only sort if desc)
		query = fmt.Sprintf("SELECT id, user_id, title, description, status, created, updated FROM %v.user_todos WHERE user_id = %v ORDER BY id DESC ALLOW FILTERING", env.DB_KEYSPACE, req.UserId)
	} else if (req.Filter != "" && req.Sort != "") {
		// With filter and sort
		query = fmt.Sprintf("SELECT id, user_id, title, description, status, created, updated FROM %v.user_todos WHERE user_id = %v AND status = '%v' ORDER BY id DESC ALLOW FILTERING", env.DB_KEYSPACE, req.UserId, req.Filter)
	}

	fmt.Printf("\n\n[Query] : %v\n\n", query)

	iter := Session.Query(query).Iter()
	
	var todo types.Todo

	if (req.Limit > 0) {
		// Pagination
		offset := 0
		for iter.Scan(&todo.TodoId, &todo.UserId, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated) {
			// Pagination Limit
			if (len(todos) == req.Limit) {
				return todos, nil
			}
			// Pagination Offset
			if (offset >= req.Offset) {
				todos = append(todos, todo)
			}
			offset++
		}
	} else {
		// Without Pagination
		for iter.Scan(&todo.TodoId, &todo.UserId, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated) {
			todos = append(todos, todo)
		}
	}

	if err := iter.Close(); err != nil {
		return todos, err
	}
	
	return todos, nil
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