package repositories

import "github.com/eure/si2018-second-half-2/entities"

type UserLikeRepository struct {
	RootRepository
}

func NewUserLikeRepository(s *Session) UserLikeRepository {
	return UserLikeRepository{NewRootRepository(s)}
}

func (r *UserLikeRepository) Create(ent entities.UserLike) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

// 自分が既にLikeしている/されている状態の全てのUserのIDを返す.
func (r *UserLikeRepository) FindLikeAll(userID int64) ([]int64, error) {
	var likes []entities.UserLike
	var ids []int64

	s := r.GetSession()
	err := s.Where("partner_id = ?", userID).Or("user_id = ?", userID).Find(&likes)
	if err != nil {
		return ids, err
	}

	for _, l := range likes {
		if l.UserID == userID {
			ids = append(ids, l.PartnerID)
			continue
		}
		ids = append(ids, l.UserID)
	}

	return ids, nil
}

// いいねを1件取得する.
// userIDはいいねを送った人, partnerIDはいいねを受け取った人.
func (r *UserLikeRepository) GetLikeBySenderIDReceiverID(userID, partnerID int64) (*entities.UserLike, error) {
	var ent entities.UserLike

	s := r.GetSession()
	has, err := s.Where("user_id = ?", userID).And("partner_id = ?", partnerID).Get(&ent)
	if err != nil {
		return nil, err
	}
	if has {
		return &ent, nil
	}
	return nil, nil
}

// マッチ済みのお相手を除き、もらったいいねを、limit/offsetで取得する.
func (r *UserLikeRepository) FindGotLikeWithLimitOffset(userID int64, limit, offset int, matchIDs []int64) ([]entities.UserLike, error) {
	var likes []entities.UserLike

	s := r.GetSession()
	s.Where("partner_id = ?", userID)
	if len(matchIDs) > 0 {
		s.NotIn("user_id", matchIDs)
	}
	s.Limit(limit, offset)
	s.Desc("created_at")
	err := s.Find(&likes)
	if err != nil {
		return likes, err
	}

	return likes, nil
}
