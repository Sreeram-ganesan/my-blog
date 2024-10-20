package mapper

import (
	"fmt"

	"github.com/Sreeram-ganesan/my-blog/internal/adapters/persist/internal/repo"
	"github.com/Sreeram-ganesan/my-blog/internal/core/model"

	"github.com/samber/lo"
)

func ContactEntityToModel(e *repo.ContactWithPhonesEntity) *model.Contact {
	phones := make([]*model.ContactPhone, len(e.Phones))
	for i, ph := range e.Phones {
		phones[i] = &model.ContactPhone{
			PhoneType:   phoneTypeEntityToModel(ph.PhoneType),
			PhoneNumber: ph.PhoneNumber,
		}
	}
	return &model.Contact{
		ID:        RepoIdToModelId(e.ID),
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Phones:    phones,
	}
}

func BlogEntityToModel(e *repo.BlogEntity) *model.Blog {
	return &model.Blog{
		ID:      RepoIdToModelId(e.ID),
		Title:   e.Title,
		Content: e.Content,
		Author:  e.Author,
	}
}

func ContactToSaveModelToEntity(m *model.ContactToSave) *repo.ContactWithPhonesEntity {
	return &repo.ContactWithPhonesEntity{
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Phones: lo.Map(m.Phones, func(item *model.ContactPhoneToSave, _ int) *repo.PhoneEntity {
			return &repo.PhoneEntity{
				PhoneType:   phoneTypeModelToEntity(item.PhoneType),
				PhoneNumber: item.PhoneNumber,
			}
		}),
	}
}

func BlogToSaveModelToEntity(m *model.BlogToSave) *repo.BlogEntity {
	return &repo.BlogEntity{
		Title:   m.Title,
		Content: m.Content,
		Author:  m.Author,
	}
}

func phoneTypeEntityToModel(phoneType string) model.ContactPhoneType {
	switch phoneType {
	case "mobile":
		return model.ContactPhoneTypeMobile
	case "home":
		return model.ContactPhoneTypeHome
	case "work":
		return model.ContactPhoneTypeWork
	default:
		panic(fmt.Sprintf("unexpected database mobile phone type: %s", phoneType))
	}
}

func phoneTypeModelToEntity(phoneType model.ContactPhoneType) string {
	switch phoneType {
	case model.ContactPhoneTypeMobile:
		return "mobile"
	case model.ContactPhoneTypeHome:
		return "home"
	case model.ContactPhoneTypeWork:
		return "work"
	default:
		panic(fmt.Sprintf("unexpected model phone type: %s", phoneType))
	}
}
