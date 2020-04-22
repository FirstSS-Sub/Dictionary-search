package mock_repository

import (
	model "github.com/FirstSS-Sub/Dictionary-search/domain/model"
	"github.com/FirstSS-Sub/Dictionary-search/usecase"
	"github.com/golang/mock/gomock"
	reflect "reflect"
	"testing"
)
/*
func TestView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected []*model.Dict
	var err error

	mockSample := NewMockDictRepository(ctrl)
	mockSample.EXPECT().FindAll().Return(expected, err)

	// mockを利用してdictUsecase.View()をテストする
	dictUsecase := usecase.NewDictUsecase(mockSample)
	result, err := dictUsecase.View()

	if err != nil {
		t.Error("Actual FindAll() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual FindAll() is not same as expected")
	}

}

 */

func TestSearchWord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected []*model.Dict
	var err error
	word := "test"

	mockSample := NewMockDictRepository(ctrl)
	mockSample.EXPECT().FindWord(word).Return(expected, err)

	// mockを利用してdictUsecase.Search(word string)をテストする
	dictUsecase := usecase.NewDictUsecase(mockSample)
	result, err := dictUsecase.SearchWord(word)

	if err != nil {
		t.Error("Actual FindWord(word string) is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual FindWord(word string) is not same as expected")
	}

}

func TestSearchTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected []*model.Dict
	var err error
	tag := "test"

	mockSample := NewMockDictRepository(ctrl)
	mockSample.EXPECT().FindTag(tag).Return(expected, err)

	// mockを利用してdictUsecase.Search(word string)をテストする
	dictUsecase := usecase.NewDictUsecase(mockSample)
	result, err := dictUsecase.SearchTag(tag)

	if err != nil {
		t.Error("Actual FindTag(tag string) is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual FindTag(tag string) is not same as expected")
	}

}

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *model.Dict
	var err error

	mockSample := NewMockDictRepository(ctrl)
	mockSample.EXPECT().Create(expected).Return(expected, err)

	// mockを利用してdictUsecase.Add(dict *model.Dict)をテストする
	dictUsecase := usecase.NewDictUsecase(mockSample)
	err = dictUsecase.Add(expected)

	if err != nil {
		t.Error("Actual Find(word string) is not same as expected")
	}

}

func TestEdit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *model.Dict
	var err error

	mockSample := NewMockDictRepository(ctrl)
	mockSample.EXPECT().Update(expected).Return(expected, err)

	// mockを利用してdictUsecase.Edit(dict *model.Dict)をテストする
	dictUsecase := usecase.NewDictUsecase(mockSample)
	err = dictUsecase.Edit(expected)

	if err != nil {
		t.Error("Actual Find(word string) is not same as expected")
	}

}
