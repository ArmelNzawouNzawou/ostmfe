package member_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/member"
)

const memb = api.BASE_URL + "memb"

func CreateMember(M member.Member) (member.Member, error) {
	entity := member.Member{}
	resp, _ := api.Rest().SetBody(M).Post(memb + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UdateMember(M member.Member) (member.Member, error) {
	entity := member.Member{}
	resp, _ := api.Rest().SetBody(M).Post(memb + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadMember(id string) (member.Member, error) {
	entity := member.Member{}
	resp, _ := api.Rest().Get(memb + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteMember(id string) (member.Member, error) {
	entity := member.Member{}
	resp, _ := api.Rest().Get(memb + "delete?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadMembers() (member.Member, error) {
	entity := member.Member{}
	resp, _ := api.Rest().Get(memb + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
