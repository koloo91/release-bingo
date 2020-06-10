package model

import (
	"github.com/google/uuid"
	"time"
)

func NewEntry(text string) *Entry {
	return &Entry{
		Id:      uuid.New().String(),
		Text:    text,
		Checked: false,
		Created: time.Now(),
		Updated: time.Now(),
	}
}

type Entry struct {
	Id      string
	Text    string
	Checked bool
	Created time.Time
	Updated time.Time
}

type EntryVo struct {
	Id      string    `json:"id"`
	Text    string    `json:"text" binding:"required"`
	Checked bool      `json:"checked"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func EntryVoToEntity(vo *EntryVo) *Entry {
	return NewEntry(vo.Text)
}

func EntryEntityToVo(entity *Entry) *EntryVo {
	return &EntryVo{
		Id:      entity.Id,
		Text:    entity.Text,
		Checked: entity.Checked,
		Created: entity.Created,
		Updated: entity.Updated,
	}
}

func EntryEntitiesToVos(entries []*Entry) []*EntryVo {
	result := make([]*EntryVo, 0, len(entries))

	for _, entry := range entries {
		result = append(result, EntryEntityToVo(entry))
	}

	return result
}
