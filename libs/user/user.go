package user

import (
	"github.com/eure/si2018-second-half-2/entities"
	"github.com/eure/si2018-second-half-2/repositories"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func BuildUserEntityByModel(meID int64, p si.PutProfileBody) entities.User {
	return entities.User{
		ID: meID,

		Nickname:       p.Nickname,
		ImageURI:       p.ImageURI,
		Tweet:          p.Tweet,
		Introduction:   p.Introduction,
		ResidenceState: p.ResidenceState,
		HomeState:      p.HomeState,
		Education:      p.Education,
		Job:            p.Job,
		AnnualIncome:   p.AnnualIncome,
		Height:         p.Height,
		BodyBuild:      p.BodyBuild,
		MaritalStatus:  p.MaritalStatus,
		Child:          p.Child,
		WhenMarry:      p.WhenMarry,
		WantChild:      p.WantChild,
		Smoking:        p.Smoking,
		Drinking:       p.Drinking,
		Holiday:        p.Holiday,
		HowToMeet:      p.HowToMeet,
		CostOfDate:     p.CostOfDate,
		NthChild:       p.NthChild,
		Housework:      p.Housework,
	}
}

func SetUsersImage(s *repositories.Session, users []entities.User) ([]entities.User, error) {
	imageMap := make(map[int64]entities.User, len(users))
	var userIDs []int64
	for _, m := range users {
		userIDs = append(userIDs, m.ID)
		imageMap[m.ID] = m
	}

	rp := repositories.NewUserImageRepository(s)
	userImage, err := rp.GetByUserIDs(userIDs)
	if err != nil {
		return nil, err
	}

	userList := make([]entities.User, len(users))
	for _, m := range userImage {
		user := imageMap[m.UserID]
		user.ImageURI = m.Path
		imageMap[m.UserID] = user
	}

	for i, m := range users {
		userList[i] = imageMap[m.ID]
	}
	return userList, err
}

func SetUserImage(s *repositories.Session, user entities.User) (entities.User, error) {
	rp := repositories.NewUserImageRepository(s)
	userImage, err := rp.GetByUserID(user.ID)
	if err != nil {
		return entities.User{}, err
	}

	user.ImageURI = userImage.Path
	return user, err
}
