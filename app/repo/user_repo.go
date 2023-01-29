package repo

import (
	"gorm.io/gorm"
	"langgo/app/models"
	"langgo/app/pkg/sqls"
)

// UserRepo .
var UserRepo = newUserRepo()

func newUserRepo() *userRepo {
	return &userRepo{}
}

type userRepo struct{}

// GetByUUID .
func (r *userRepo) GetByUUID(db *gorm.DB, userID string) (*models.User, error) {
	ret := &models.User{}
	if err := db.Where("status = ?", 1).Where("uuid = ?", userID).First(ret).Error; err != nil {
		return ret, err
	}
	return ret, nil
}

// Find .
func (r *userRepo) Find(db *gorm.DB, cnd *sqls.Condition) ([]models.User, error) {
	var ret []models.User
	if err := cnd.Find(db, &ret); err != nil {
		return ret, err
	}
	return ret, nil
}

// FindOne .
func (r *userRepo) FindOne(db *gorm.DB, cnd *sqls.Condition) (*models.User, error) {
	ret := &models.User{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return ret, err
	}
	return ret, nil
}

// Count .
func (r *userRepo) Count(db *gorm.DB, cnd *sqls.Condition) int64 {
	return cnd.Count(db, &models.User{})
}

// FindPage 翻页查询
func (r *userRepo) FindPage(db *gorm.DB, cnd *sqls.Condition) ([]models.User, *sqls.PageInfo, error) {
	count := cnd.Count(db, &models.User{})
	paging := &sqls.PageInfo{
		Page:  cnd.PageInfo.Page,
		Limit: cnd.PageInfo.Limit,
		Total: count,
	}

	ret, err := r.Find(db, cnd)
	return ret, paging, err
}

// Create .
func (r *userRepo) Create(db *gorm.DB, m *models.User) error {
	err := db.Create(m).Error
	return err
}

// UpdateAll 更新整条数据
func (r *userRepo) UpdateAll(db *gorm.DB, m *models.User) error {
	err := db.Save(m).Error
	return err
}

// Updates .
func (r *userRepo) Updates(db *gorm.DB, userID string, columns map[string]interface{}) error {
	err := db.Model(&models.User{}).Where("uuid = ?", userID).Updates(columns).Error
	return err
}

// UpdateColumn .
func (r *userRepo) UpdateColumn(db *gorm.DB, userID string, name string, value interface{}) error {
	err := db.Model(&models.User{}).Where("uuid = ?", userID).UpdateColumn(name, value).Error
	return err
}

// Delete .
func (r *userRepo) Delete(db *gorm.DB, userID string) error {
	err := db.Model(&models.User{}).Where("uuid = ?", userID).Where("status = ?", 1).
		UpdateColumn("status", -1).Error
	return err
}
