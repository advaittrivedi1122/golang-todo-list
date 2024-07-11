# Todos List in Golang
A Todos List - CRUD API with pagination.
(Golang + ScyllaDb)

<hr>

## Steps to start the server.

- Make sure you are in linux machine.
- Make sure you have docker installed in your machine.

### 1. Clone the repo.
### 2. Run these Commands.

#### -> Start Go server + ScyllaDb
`sudo bash start.sh`

#### -> Stop ScyllaDb
`sudo bash stop.sh`

<hr>

## API Endpoints

```
1) [POST] - /add-todo               (Add user todo)
2) [GET]  - /get-user-todo-by-id    (Get single todo for user)
3) [GET]  - /get-user-todos         (Get all todos for user)
4) [POST] - /update-user-todo-by-id (Update single todo for user)
5) [POST] - /delete-user-todo-by-id (Delete todo created by user)
6) [POST] - /delete-user-todos      (Delete all todos for user)
```

## API Request Params

```
1) /add-todo
        user_id*     (int)
        title*       (string)
        description* (string)
        status*      (string - pending/completed)

2) /get-user-todo-by-id
        user_id*     (int)
        id*          (int)

3) /get-user-todos
        user_id*     (int)
        sort         (string - asc/desc)
        filter       (string - pending/completed)
        limit        (int)
        offset       (int)

4) /update-user-todo-by-id
        user_id*      (int)
        id*           (int)
        title         (string)
        description   (string)
        status        (string)

5) /delete-user-todo-by-id
        user_id*      (int)
        id*           (int)

6) /delete-user-todos
        user_id*      (int)
```