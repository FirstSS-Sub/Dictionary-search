package injector

import (
	"github.com/FirstSS-Sub/Dictionary-search/domain/repository"
	"github.com/FirstSS-Sub/Dictionary-search/handler"
	"github.com/FirstSS-Sub/Dictionary-search/infra"
	"github.com/FirstSS-Sub/Dictionary-search/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return *sqlhandler
}

/*
DictRepository(interface)に実装であるSqlHandler(struct)を渡し生成する。
*/
func InjectDictRepository() repository.DictRepository {
	sqlHandler := InjectDB()
	return infra.NewDictRepository(sqlHandler)
}

func InjectDictUsecase() usecase.DictUsecase {
	DictRepo := InjectDictRepository()
	return usecase.NewDictUsecase(DictRepo)
}

func InjectDictHandler() handler.DictHandler {
	return handler.NewDictHandler(InjectDictUsecase())
}
