package language_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/language"
)

const lang = api.BASE_URL + "lang"

func CreateLanguage(L language.Language) (language.Language, error) {
	entity := language.Language{}
	resp, _ := api.Rest().SetBody(L).Post(lang + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateLanguage(L language.Language) (language.Language, error) {
	entity := language.Language{}
	resp, _ := api.Rest().SetBody(L).Post(lang + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadLanguage(id string) (language.Language, error) {
	entity := language.Language{}
	resp, _ := api.Rest().Get(lang + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteLanguage(id string) (language.Language, error) {
	entity := language.Language{}
	resp, _ := api.Rest().Get(lang + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadLanguages() (language.Language, error) {
	entity := language.Language{}
	resp, _ := api.Rest().Get(lang + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
