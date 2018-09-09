package entities

import (
	"github.com/eure/si2018-second-half-2/models"
	"github.com/go-openapi/strfmt"
)

type User struct {
	ID             int64           `xorm:"id"`
	Gender         string          `xorm:"gender"`
	Birthday       strfmt.Date     `xorm:"birthday"`
	Nickname       string          `xorm:"nickname"`
	Tweet          string          `xorm:"tweet"`
	Introduction   string          `xorm:"introduction"`
	ResidenceState string          `xorm:"residence_state"`
	HomeState      string          `xorm:"home_state"`
	Education      string          `xorm:"education"`
	Job            string          `xorm:"job"`
	AnnualIncome   string          `xorm:"annual_income"`
	Height         string          `xorm:"height"`
	BodyBuild      string          `xorm:"body_build"`
	MaritalStatus  string          `xorm:"marital_status"`
	Child          string          `xorm:"child"`
	WhenMarry      string          `xorm:"when_marry"`
	WantChild      string          `xorm:"want_child"`
	Smoking        string          `xorm:"smoking"`
	Drinking       string          `xorm:"drinking"`
	Holiday        string          `xorm:"holiday"`
	HowToMeet      string          `xorm:"how_to_meet"`
	CostOfDate     string          `xorm:"cost_of_date"`
	NthChild       string          `xorm:"nth_child"`
	Housework      string          `xorm:"housework"`
	CreatedAt      strfmt.DateTime `xorm:"created_at"`
	UpdatedAt      strfmt.DateTime `xorm:"updated_at"`

	ImageURI string `xorm:"-"`
}

func (u User) Build() models.User {
	return models.User{
		ID:             u.ID,
		Gender:         u.Gender,
		Birthday:       u.Birthday,
		Nickname:       u.Nickname,
		ImageURI:       u.ImageURI,
		Tweet:          u.Tweet,
		Introduction:   u.Introduction,
		ResidenceState: u.ResidenceState,
		HomeState:      u.HomeState,
		Education:      u.Education,
		Job:            u.Job,
		AnnualIncome:   u.AnnualIncome,
		Height:         u.Height,
		BodyBuild:      u.BodyBuild,
		MaritalStatus:  u.MaritalStatus,
		Child:          u.Child,
		WhenMarry:      u.WhenMarry,
		WantChild:      u.WantChild,
		Smoking:        u.Smoking,
		Drinking:       u.Drinking,
		Holiday:        u.Holiday,
		HowToMeet:      u.HowToMeet,
		CostOfDate:     u.CostOfDate,
		NthChild:       u.NthChild,
		Housework:      u.Housework,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}

func (u User) GetOppositeGender() string {
	if u.Gender == "F" {
		return "M"
	}
	return "F"
}

type Users []User

func (users *Users) Build() []*models.User {
	var sUsers []*models.User

	for _, u := range *users {
		swaggerUser := u.Build()
		sUsers = append(sUsers, &swaggerUser)
	}
	return sUsers
}
