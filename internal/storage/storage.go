package storage

import (
	"errors"
	"sync"

	"github.com/zviadadze/userver/internal/models"
)

var (
	mu     sync.Mutex
	nextID int
	idPool []int                = make([]int, 2)
	users  map[int]*models.User = make(map[int]*models.User, 10)
)

func init() {
	nextID = 0
	users = make(map[int]*models.User, 10)
	idPool = []int{nextID}
}

func getNewID() int {
	if len(idPool) == 0 {
		nextID++
		return nextID
	}
	newID := idPool[0]
	idPool = idPool[1:]
	return newID
}

func AppendUser(name string, age int) *models.User {
	mu.Lock()
	defer mu.Unlock()
	id := getNewID()
	user := &models.User{ID: id, Name: name, Age: age}
	users[id] = user
	return user
}

func GetUsers() map[int]*models.User {
	return users
}

func GetUser(id int) (*models.User, error) {
	mu.Lock()
	defer mu.Unlock()
	user, isExist := users[id]
	if !isExist {
		return nil, errors.New("user does not exist")
	}
	return user, nil
}

func RemoveUser(id int) (*models.User, error) {
	mu.Lock()
	defer mu.Unlock()
	removedUser, isExist := users[id]
	if !isExist {
		return nil, errors.New("user does not exist")
	}
	delete(users, id)
	idPool = append(idPool, id)
	return removedUser, nil
}
