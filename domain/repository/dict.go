package repository

import "github.com/FirstSS-Sub/Dictionary-search/domain/model"

// modelの永続化を行うリポジトリ、DictRepository を定義
// ここではClean Architectureの依存性の順番を守るためinterface（抽象）のみを定義し、実際の実装はinfra層で行う
type DictRepository interface {
	FindWord(word string) (dict []*model.Dict, err error)
	FindTag(tag string) (dict []*model.Dict, err error)
	Create(dict *model.Dict) (*model.Dict, error)
	Update(dict *model.Dict) (*model.Dict, error)
}
