package usecase

import (
	"github.com/FirstSS-Sub/Dictionary-search/domain/model"
	"github.com/FirstSS-Sub/Dictionary-search/domain/repository"
)

type DictUsecase interface {
	SearchWord(string) (dict []*model.Dict, err error)
	SearchTag(string) (dict []*model.Dict, err error)
	// View() (dict []*model.Dict, err error)
	Add(*model.Dict) (err error)
	Edit(*model.Dict) (err error)
}

type dictUsecase struct {
	dictRep repository.DictRepository
}

func NewDictUsecase(dictRep repository.DictRepository) DictUsecase {
	dictUsecase := dictUsecase{dictRep: dictRep}
	return &dictUsecase
}

// SearchWord 入力された内容でDictを検索する
func (usecase *dictUsecase) SearchWord(word string) (dict []*model.Dict, err error) {
	dict, err = usecase.dictRep.FindWord(word)
	return
}

// SearchTag 入力されたタグでDictを検索する
func (usecase *dictUsecase) SearchTag(tag string) (dict []*model.Dict, err error) {
	dict, err = usecase.dictRep.FindTag(tag)
	return
}

/*
// View はDictの一覧を表示する
func (usecase *dictUsecase) View() (dict []*model.Dict, err error) {
	dict, err = usecase.dictRep.FindAll()
	return
}
*/

// Add はDictを新規追加する
func (usecase *dictUsecase) Add(dict *model.Dict) (err error) {
	_, err = usecase.dictRep.Create(dict)
	return
}

// Edit はDictを編集する
func (usecase *dictUsecase) Edit(dict *model.Dict) (err error) {
	_, err = usecase.dictRep.Update(dict)
	return
}
