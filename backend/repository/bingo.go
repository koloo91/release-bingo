package repository

import (
	"github.com/koloo91/release-bingo/model"
	"strings"
	"sync"
)

var (
	userDatabase = make(map[string]*model.BingoUser)
	mutex        sync.Mutex
)

func GetAllUsers() []*model.BingoUser {
	result := make([]*model.BingoUser, 0, len(userDatabase))

	for _, user := range userDatabase {
		result = append(result, user)
	}

	return result
}

func GetAllUserNames() []string {
	result := make([]string, 0, len(userDatabase))

	for _, user := range userDatabase {
		result = append(result, user.Name)
	}

	return result
}

func GetUserByName(name string) *model.BingoUser {
	if user, ok := userDatabase[name]; ok {
		return user
	}
	return nil
}

func AddUser(user *model.BingoUser) {
	mutex.Lock()
	nameToLower := strings.ToLower(user.Name)
	userDatabase[nameToLower] = user
	mutex.Unlock()
}

func RemoveUser(userName string) {
	mutex.Lock()
	nameToLower := strings.ToLower(userName)
	delete(userDatabase, nameToLower)
	mutex.Unlock()
}
