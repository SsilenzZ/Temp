package social

import (
	"Web-chat/pkg/models"
	"gorm.io/gorm"
)

type SocialRepositoryI interface{}

type SocialRepository struct {
	DB *gorm.DB
}

func ProvideSocialRepository(DB *gorm.DB) SocialRepository {
	return SocialRepository{DB: DB}
}

func (s SocialRepository) Friend(id1, id2 int) bool {
	var social models.Social
	if s.CheckBlocked(id1, id2) {
		return false
	}
	if s.getStatus(id1, id2, 1, social) {
		return false
	} else {
		social = models.Social{User_ID1: id1, User_ID2: id2, Status: 1}
		s.DB.Create(&social)
		return true
	}
}

func (s SocialRepository) Block(id1, id2 int) bool {
	var social models.Social
	if s.CheckBlocked(id1, id2) {
		return false
	}
	if s.getStatus(id1, id2, 1, social) {
		s.DB.Model(&social).Where("user_id1 = ? AND user_id2 = ? AND status = 1", id1, id2).
			Updates(map[string]interface{}{"status": 0})
		return true
	} else {
		social = models.Social{User_ID1: id1, User_ID2: id2, Status: 0}
		s.DB.Create(&social)
		return true
	}
}

func (s SocialRepository) Unfriend(id1, id2 int) {
	var social models.Social
	s.DB.Where("user_id1 = ? AND user_id2 = ? AND status = 1", id1, id2).Delete(&social)
}

func (s SocialRepository) Unblock(id1, id2 int) {
	var social models.Social
	s.DB.Where("user_id1 = ? AND user_id2 = ? AND status = 0", id1, id2).Delete(&social)
}

func (s SocialRepository) GetFriends(id int) []models.Social {
	var social []models.Social
	s.DB.Where("user_id1 = ? AND status = 1", id).Find(&social)
	return social
}

func (s SocialRepository) GetBlocked(id int) []models.Social {
	var social []models.Social
	s.DB.Where("user_id1 = ? AND status = 0", id).Find(&social)
	return social
}

func (s SocialRepository) getStatus(id1, id2, status int, social models.Social) bool {
	err := s.DB.Where("user_id1 = ? AND user_id2 = ? and status = ?", id1, id2, status).Take(&social).Error
	return err == nil
}

func (s SocialRepository) CheckBlocked(id1, id2 int) bool {
	var social models.Social
	return s.getStatus(id1, id2, 0, social)
}
