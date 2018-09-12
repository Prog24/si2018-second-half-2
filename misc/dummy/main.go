package main

import (
	"math/rand"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/eure/si2018-second-half-2/entities"
	"github.com/eure/si2018-second-half-2/repositories"
)

const (
	firstUserID   = 1
	maleIDStart   = 1
	maleIDEnd     = 1000
	femaleIDStart = 1001
	femaleIDEnd   = 2000
	lastUserID    = 2000

	maleMessageUserID   = 222
	femaleMessageUserID = 1222
)

func main() {
	dummyUser()              // Male 1-1000, Female 1001-2000
	dummyToken()             // Token for each Users
	dummyImage()             // Images for each Users
	dummyManyMessageCouple() // マッチ & メッセージしてるカップル M222 と F2222
	dummyManyGotLikeUser()   // F1111 に M1〜M100の男性からの被いいね
	dummyManyMatchUser()     // F1112 UID 1〜200の男性が UID 1112の女性とマッチ
}

func dummyManyGotLikeUser() {
	r := repositories.NewUserLikeRepository()

	for i := 1; i <= 100; i++ {
		rand.Seed(time.Now().UnixNano())
		createdDaysAgo := rand.Intn(600)
		minute1 := rand.Intn(1440)
		randTime := strfmt.DateTime(time.Now().AddDate(0, 0, -createdDaysAgo).Add(-time.Duration(minute1) * time.Minute))

		ent := entities.UserLike{
			UserID:    int64(i),
			PartnerID: 1111,
			CreatedAt: randTime,
			UpdatedAt: randTime,
		}
		r.Create(ent)
	}
}

func dummyManyMatchUser() {
	firstLikeDate := strfmt.DateTime(time.Now().AddDate(0, 0, -3))
	responseLikeDate := strfmt.DateTime(time.Now().AddDate(0, 0, -2))

	lr := repositories.NewUserLikeRepository()
	mr := repositories.NewUserMatchRepository()

	// Male 1-100 & Female 1112
	// =====================================================

	// first like
	for i := 1; i <= 100; i++ {
		ent := entities.UserLike{
			UserID:    int64(i),
			PartnerID: 1112,
			CreatedAt: firstLikeDate,
			UpdatedAt: firstLikeDate,
		}
		lr.Create(ent)
	}

	// resp like
	for i := 1; i <= 100; i++ {
		ent := entities.UserLike{
			UserID:    1112,
			PartnerID: int64(i),
			CreatedAt: responseLikeDate,
			UpdatedAt: responseLikeDate,
		}
		lr.Create(ent)
	}

	// match
	for i := 1; i <= 100; i++ {
		ent := entities.UserMatch{
			UserID:    int64(i),
			PartnerID: 1112,
			CreatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Duration(i) * time.Minute)), // Paginationのため分刻みでイテレート
			UpdatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Duration(i) * time.Minute)), // Paginationのため分刻みでイテレート
		}
		mr.Create(ent)
	}

	// Male 101-200 & Female 1112
	// =====================================================

	// first like
	for i := 101; i <= 200; i++ {
		ent := entities.UserLike{
			UserID:    1112,
			PartnerID: int64(i),
			CreatedAt: firstLikeDate,
			UpdatedAt: firstLikeDate,
		}
		lr.Create(ent)
	}

	// resp like
	for i := 101; i <= 200; i++ {
		ent := entities.UserLike{
			UserID:    int64(i),
			PartnerID: 1112,
			CreatedAt: responseLikeDate,
			UpdatedAt: responseLikeDate,
		}
		lr.Create(ent)
	}

	// match
	for i := 101; i <= 200; i++ {
		ent := entities.UserMatch{
			UserID:    1112,
			PartnerID: int64(i),
			CreatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Duration(i) * time.Minute)), // Paginationのため分刻みでイテレート
			UpdatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Duration(i) * time.Minute)), // Paginationのため分刻みでイテレート
		}
		mr.Create(ent)
	}
}

func dummyManyMessageCouple() {
	today := strfmt.DateTime(time.Now())
	yesterday := strfmt.DateTime(time.Now().AddDate(0, 0, -1))

	// マッチの前提として相互いいねが必要
	lr := repositories.NewUserLikeRepository()
	lr.Create(
		entities.UserLike{
			UserID:    maleMessageUserID,
			PartnerID: femaleMessageUserID,
			CreatedAt: yesterday,
			UpdatedAt: yesterday,
		})
	lr.Create(
		entities.UserLike{
			UserID:    femaleMessageUserID,
			PartnerID: maleMessageUserID,
			CreatedAt: today,
			UpdatedAt: today,
		})

	// メッセージの前提としてマッチが必要
	mr := repositories.NewUserMatchRepository()
	mr.Create(
		entities.UserMatch{
			UserID:    maleMessageUserID,
			PartnerID: femaleMessageUserID,
			CreatedAt: today,
			UpdatedAt: today,
		})

	rand.Seed(time.Now().UnixNano())
	createdDaysAgo := rand.Intn(600)
	minute1 := rand.Intn(1440)
	randTime := strfmt.DateTime(time.Now().AddDate(0, 0, -createdDaysAgo).Add(-time.Duration(minute1) * time.Minute))

	// 双方向に130件ずつメッセージ
	msgr := repositories.NewUserMessageRepository()
	for i := 0; i <= 130; i++ {
		ent := entities.UserMessage{
			UserID:    maleMessageUserID,
			PartnerID: femaleMessageUserID,
			Message:   "hello",
			CreatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -createdDaysAgo).Add(-time.Duration(i*60) * time.Minute)),
			UpdatedAt: randTime,
		}
		msgr.Create(ent)
	}
	for i := 0; i <= 130; i++ {
		ent := entities.UserMessage{
			UserID:    femaleMessageUserID,
			PartnerID: maleMessageUserID,
			Message:   "hi!",
			CreatedAt: strfmt.DateTime(time.Now().AddDate(0, 0, -createdDaysAgo).Add(-time.Duration(i*60+1) * time.Minute)),
			UpdatedAt: randTime,
		}
		msgr.Create(ent)
	}
}

// covert "1994-12-24" style string to strfmt.Date
func stringToStrFmtDate(str string) strfmt.Date {
	var date strfmt.Date
	date.UnmarshalText([]byte(str))
	return date
}
