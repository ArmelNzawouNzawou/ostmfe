package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const eventPrj = api.BASE_URL + "eventPrj"

func CreateEventProject(prj project.Project) (project.Project, error) {
	entity := project.Project{}
	resp, _ := api.Rest().SetBody(prj).Post(eventPrj + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventProject(prj project.Project) (project.Project, error) {
	entity := project.Project{}
	resp, _ := api.Rest().SetBody(prj).Post(eventPrj + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventProject(id string) (project.Project, error) {
	entity := project.Project{}
	resp, _ := api.Rest().Get(eventPrj + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteEventProject(id string) (project.Project, error) {
	entity := project.Project{}
	resp, _ := api.Rest().Get(eventPrj + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventProjects() (project.Project, error) {
	entity := project.Project{}
	resp, _ := api.Rest().Get(eventPrj + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
