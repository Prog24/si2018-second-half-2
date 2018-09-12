package main

import (
	"math/rand"
	"time"

	"github.com/eure/si2018-server-side/entities"
	"github.com/eure/si2018-server-side/repositories"
	"github.com/go-openapi/strfmt"
)

func dummyImage() {
	maleImg := []string{
		"https://si-2018-000.eure.jp/assets/eye_yorime_yoseme_man.png",
		"https://si-2018-000.eure.jp/assets/hair_kariage.png",
		"https://si-2018-000.eure.jp/assets/kyosyu_man.png",
		"https://si-2018-000.eure.jp/assets/nage_kiss_man.png",
		"https://si-2018-000.eure.jp/assets/perm_hair_man.png",
		"https://si-2018-000.eure.jp/assets/pose_ayashii_man.png",
		"https://si-2018-000.eure.jp/assets/pose_kandou_man.png",
		"https://si-2018-000.eure.jp/assets/pose_nigawarai_man.png",
		"https://si-2018-000.eure.jp/assets/pose_shock_man.png",
		"https://si-2018-000.eure.jp/assets/tehepero2_youngman.png",
	}

	femaleImg := []string{
		"https://si-2018-000.eure.jp/assets/eye_yorime_yoseme_man.png",
		"https://si-2018-000.eure.jp/assets/hair_kariage.png",
		"https://si-2018-000.eure.jp/assets/kyosyu_man.png",
		"https://si-2018-000.eure.jp/assets/nage_kiss_man.png",
		"https://si-2018-000.eure.jp/assets/perm_hair_man.png",
		"https://si-2018-000.eure.jp/assets/pose_ayashii_man.png",
		"https://si-2018-000.eure.jp/assets/pose_kandou_man.png",
		"https://si-2018-000.eure.jp/assets/pose_nigawarai_man.png",
		"https://si-2018-000.eure.jp/assets/pose_shock_man.png",
		"https://si-2018-000.eure.jp/assets/tehepero2_youngman.png",
	}

	r := repositories.NewUserImageRepository()

	// for male
	for i := maleIDStart; i <= maleIDEnd; i++ {
		now := strfmt.DateTime(time.Now())
		image := entities.UserImage{
			UserID:    int64(i),
			Path:      maleImg[rand.Intn(len(maleImg))],
			CreatedAt: now,
			UpdatedAt: now,
		}
		r.Create(image)
	}

	// for female
	for i := femaleIDStart; i <= femaleIDEnd; i++ {
		now := strfmt.DateTime(time.Now())
		image := entities.UserImage{
			UserID:    int64(i),
			Path:      femaleImg[rand.Intn(len(femaleImg))],
			CreatedAt: now,
			UpdatedAt: now,
		}
		r.Create(image)
	}
}
