package service

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/koloo91/release-bingo/model"
	"github.com/koloo91/release-bingo/repository"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	random = rand.New(rand.NewSource(time.Now().Unix()))

	updatePlayerDataMutex sync.Mutex
)

func UpdatePlayerData(ctx context.Context) error {

	entries, err := repository.GetAllEntries(ctx)
	if err != nil {
		return err
	}

	updatePlayerDataMutex.Lock()

	entriesMap := entriesListToMap(entries)
	for _, player := range repository.GetAllUsers() {
		for _, field := range player.Card.RowOne {
			if entry, ok := entriesMap[field.Id]; ok {
				field.Checked = entry.Checked
			}
		}

		for _, field := range player.Card.RowTwo {
			if entry, ok := entriesMap[field.Id]; ok {
				field.Checked = entry.Checked
			}
		}

		for _, field := range player.Card.RowThree {
			if entry, ok := entriesMap[field.Id]; ok {
				field.Checked = entry.Checked
			}
		}

		for _, field := range player.Card.RowFour {
			if entry, ok := entriesMap[field.Id]; ok {
				field.Checked = entry.Checked
			}
		}
	}

	broadcastGameData()
	updatePlayerDataMutex.Unlock()
	return nil
}

func entriesListToMap(entries []*model.Entry) map[string]*model.Entry {
	result := make(map[string]*model.Entry, 0)
	for _, entry := range entries {
		result[entry.Id] = entry
	}
	return result
}

func ConnectToGame(userName string, connection *websocket.Conn) {
	userNameToLower := strings.ToLower(userName)
	if user := repository.GetUserByName(userNameToLower); user != nil {
		log.Printf("user with name '%s' already exists, disconnecting", userName)
		if err := connection.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(400, "user with name already exists"), time.Now().Add(time.Second)); err != nil {
			log.Println(err)
		}
		return
	}

	log.Printf("User '%s' connected", userName)

	card, err := generateRandomCard()
	if err != nil {
		connection.Close()
		return
	}

	bingoUser := &model.BingoUser{
		Name:       userName,
		Connection: connection,
		Card:       card,
	}

	repository.AddUser(bingoUser)

	connection.SetCloseHandler(addConnectionClose(userName))

	go func() {
		for {
			if _, _, err := connection.ReadMessage(); err != nil {
				return
			}
		}
	}()

	broadcastGameData()
}

func addConnectionClose(userName string) func(int, string) error {
	return func(i int, s string) error {
		log.Printf("User '%s' disconected", userName)
		repository.RemoveUser(userName)
		broadcastGameData()
		return nil
	}
}

func generateRandomCard() (*model.BingoCard, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entries, err := repository.GetAllEntries(ctx)
	if err != nil {
		return nil, err
	}

	card := &model.BingoCard{}

	entriesLength := len(entries)
	if entriesLength == 0 {
		return card, nil
	}

	var randomIndex int
	var entry *model.Entry

	for i := 0; i < 4; i++ {
		randomIndex = random.Intn(entriesLength)
		entry = entries[randomIndex]
		card.RowOne = append(card.RowOne, &model.BingoField{
			Id:      entry.Id,
			Text:    entry.Text,
			Checked: entry.Checked,
		})

		randomIndex = random.Intn(entriesLength)
		entry = entries[randomIndex]
		card.RowTwo = append(card.RowTwo, &model.BingoField{
			Id:      entry.Id,
			Text:    entry.Text,
			Checked: entry.Checked,
		})

		randomIndex = random.Intn(entriesLength)
		entry = entries[randomIndex]
		card.RowThree = append(card.RowThree, &model.BingoField{
			Id:      entry.Id,
			Text:    entry.Text,
			Checked: entry.Checked,
		})

		randomIndex = random.Intn(entriesLength)
		entry = entries[randomIndex]
		card.RowFour = append(card.RowFour, &model.BingoField{
			Id:      entry.Id,
			Text:    entry.Text,
			Checked: entry.Checked,
		})
	}

	return card, nil
}

func broadcastGameData() {
	userNames := repository.GetAllUserNames()

	for _, user := range repository.GetAllUsers() {
		gameData := model.GameVo{
			Users: userNames,
			User:  model.MapBingoUserToVo(user),
		}

		if err := user.Connection.WriteJSON(gameData); err != nil {
			log.Printf("Error sending current game state to user '%s': %v", user.Name, err)
		}
	}
}
