package infra

import (
	"fmt"
	"github.com/FirstSS-Sub/Dictionary-search/domain/model"
	"github.com/FirstSS-Sub/Dictionary-search/domain/repository"
)

type DictRepository struct {
	SqlHandler
}

func NewDictRepository(sqlHandler SqlHandler) repository.DictRepository {
	dictRepository := DictRepository{sqlHandler}
	return &dictRepository
}

func (dictRep *DictRepository) FindWord(word string) (dicts []*model.Dict, err error) {
	rows, err := dictRep.SqlHandler.Conn.Query(
		"SELECT * FROM dicts WHERE JpWord LIKE ? or EngWord LIKE ? or Body LIKE ?", "%"+word+"%", "%"+word+"%", "%"+word+"%")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		dict := model.Dict{}

		rows.Scan(&dict.ID, &dict.JpName, &dict.EngName, &dict.Body, &dict.Tags)

		dicts = append(dicts, &dict)
	}
	return
}

func (dictRep *DictRepository) FindTag(tag string) (dicts []*model.Dict, err error) {
	rows, err := dictRep.SqlHandler.Conn.Query("SELECT * FROM dicts WHERE Tags=?", tag)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		dict := model.Dict{}

		rows.Scan(&dict.ID, &dict.JpName, &dict.EngName, &dict.Body, &dict.Tags)

		dicts = append(dicts, &dict)
	}
	return
}
