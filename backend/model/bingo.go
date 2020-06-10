package model

import "github.com/gorilla/websocket"

type GameVo struct {
	Users []string     `json:"users"`
	User  *BingoUserVo `json:"user"`
}

type BingoUser struct {
	Name       string
	Connection *websocket.Conn
	Card       *BingoCard
}

type BingoUserVo struct {
	Name string       `json:"name"`
	Card *BingoCardVo `json:"card"`
}

type BingoCard struct {
	RowOne   []*BingoField
	RowTwo   []*BingoField
	RowThree []*BingoField
	RowFour  []*BingoField
}

type BingoCardVo struct {
	RowOne   []*BingoFieldVo `json:"rowOne"`
	RowTwo   []*BingoFieldVo `json:"rowTwo"`
	RowThree []*BingoFieldVo `json:"rowThree"`
	RowFour  []*BingoFieldVo `json:"rowFour"`
}

type BingoField struct {
	Id      string
	Text    string
	Checked bool
}

type BingoFieldVo struct {
	Id      string `json:"id"`
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
}

func MapBingoUserToVo(entity *BingoUser) *BingoUserVo {
	return &BingoUserVo{
		Name: entity.Name,
		Card: mapBingoCardToVo(entity.Card),
	}
}

func mapBingoCardToVo(entity *BingoCard) *BingoCardVo {
	return &BingoCardVo{
		RowOne:   mapBingoFieldsToVos(entity.RowOne),
		RowTwo:   mapBingoFieldsToVos(entity.RowTwo),
		RowThree: mapBingoFieldsToVos(entity.RowThree),
		RowFour:  mapBingoFieldsToVos(entity.RowFour),
	}
}

func mapBingoFieldToVo(entity *BingoField) *BingoFieldVo {
	return &BingoFieldVo{
		Id:      entity.Id,
		Text:    entity.Text,
		Checked: entity.Checked,
	}
}

func mapBingoFieldsToVos(entities []*BingoField) []*BingoFieldVo {
	result := make([]*BingoFieldVo, 0, len(entities))

	for _, entity := range entities {
		result = append(result, mapBingoFieldToVo(entity))
	}

	return result
}
