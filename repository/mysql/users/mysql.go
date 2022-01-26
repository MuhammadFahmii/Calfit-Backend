package users

import (
	"CalFit/business/paginations"
	"CalFit/business/users"
	"CalFit/exceptions"
	"CalFit/repository/mysql/addresses"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UsersRepo struct {
	DBConn *gorm.DB
}

<<<<<<< HEAD
type ProfilesRepo struct {
=======
type ProfileUsersRepo struct {
>>>>>>> master
	DBConn *gorm.DB
}

func NewUsersRepo(db *gorm.DB) users.Repository {
	return &UsersRepo{
		DBConn: db,
	}
}

<<<<<<< HEAD
func NewProfileRepo(db *gorm.DB) users.ProfileRepository {
	return &ProfilesRepo{
=======
func NewProfileUsersRepo(db *gorm.DB) users.ProfileRepository {
	return &ProfileUsersRepo{
>>>>>>> master
		DBConn: db,
	}
}

func (repo *UsersRepo) LoginOAuth(ctx context.Context, domain users.Domain) (users.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Where("email=?", data.Email).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data.MembershipTypeID = 1
			data.CreatedAt = time.Now()
			address := addresses.Address{
				Address:     "default",
				District:    "default",
				City:        "default",
				Postal_code: "11111",
			}
			repo.DBConn.Create(&address)
			data.AddressID = address.Id
			repo.DBConn.Create(&data)
			return data.ToDomain(), nil
		}
		return users.Domain{}, err
	}
	return data.ToDomain(), nil
}

func (repo *UsersRepo) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Where("email=?", data.Email).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data.MembershipTypeID = 1
			data.CreatedAt = time.Now()
			address := addresses.Address{
				Address:     "default",
				District:    "default",
				City:        "default",
				Postal_code: "11111",
			}
			repo.DBConn.Create(&address)
			data.AddressID = address.Id
			repo.DBConn.Create(&data)
			return data.ToDomain(), nil
		}
		return users.Domain{}, err
	}
	return users.Domain{}, exceptions.ErrUserAlreadyExists
}

func (repo *UsersRepo) GetByUsername(ctx context.Context, email string) (users.Domain, error) {
	data := User{}
	if err := repo.DBConn.Where("email=?", email).First(&data).Error; err != nil {
		return users.Domain{}, err
	}
	domain := data.ToDomain()
	type Relation struct {
		MembershipName string
	}
	relation := Relation{}
	repo.DBConn.Table("users").Select("membership_types.name AS membership_name").Joins("LEFT JOIN membership_types ON users.membership_type_id = membership_types.id").Scan(&relation)
	domain.MembershipName = relation.MembershipName
	return domain, nil
}

func (repo *UsersRepo) Update(ctx context.Context, domain users.Domain) (users.Domain, error) {
	data := FromDomain(domain)
	if err := repo.DBConn.Where("email=?", data.Email).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.Domain{}, exceptions.ErrNotFound
		}
		return users.Domain{}, err
	}
	data.FullName = domain.FullName
	if domain.Password != "" {
		data.Password = domain.Password
	}
	data.UpdatedAt = time.Now()
	if err := repo.DBConn.Save(&data).Error; err != nil {
		return users.Domain{}, err
	}
	return data.ToDomain(), nil
}

<<<<<<< HEAD
func (b *ProfilesRepo) GetAll(ctx context.Context, pagination paginations.Domain) ([]users.Domain, error) {
=======
func (b *ProfileUsersRepo) GetAll(ctx context.Context, pagination paginations.Domain) ([]users.Domain, error) {
>>>>>>> master
	var usersModel []User

	offset := (pagination.Page - 1) * pagination.Limit
	if err := b.DBConn.Preload("Address").Preload("BookingDetails").Limit(pagination.Limit).Offset(offset).Find(&usersModel).Error; err != nil {
		return nil, err
	}
	var result []users.Domain = ToListDomain(usersModel)
	return result, nil
}

<<<<<<< HEAD
func (b *ProfilesRepo) CountAll(ctx context.Context) (int, error) {
=======
func (b *ProfileUsersRepo) CountAll(ctx context.Context) (int, error) {
>>>>>>> master
	var count int64
	if err := b.DBConn.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

<<<<<<< HEAD
func (b *ProfilesRepo) GetById(ctx context.Context, id string) (users.Domain, error) {
=======
func (b *ProfileUsersRepo) GetById(ctx context.Context, id string) (users.Domain, error) {
>>>>>>> master
	var user User
	if err := b.DBConn.Preload("Address").Preload("Classes").Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return users.Domain{}, exceptions.ErrNotFound
		}
		return users.Domain{}, err
	}
	return user.ToDomain(), nil
}

<<<<<<< HEAD
func (b *ProfilesRepo) Update(ctx context.Context, id string, user users.Domain) (users.Domain, error) {
=======
func (b *ProfileUsersRepo) Update(ctx context.Context, id string, user users.Domain) (users.Domain, error) {
>>>>>>> master
	var userModel User
	if err := b.DBConn.Where("id = ?", id).Preload("Address").First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return users.Domain{}, exceptions.ErrUserNotFound
		}
		return users.Domain{}, err
	}

	userModel.Email = user.Email
	userModel.Password = user.Password
	userModel.AddressID = user.AddressID
	userModel.Photo = user.Photo
	userModel.MembershipTypeID = user.MembershipTypeID
	userModel.FullName = user.FullName
	userModel.UpdatedAt = time.Now()

	updateErr := b.DBConn.Save(&userModel).Error
	if updateErr != nil {
		return users.Domain{}, updateErr
	}
	return userModel.ToDomain(), nil
}
func (repo *UsersRepo) GetByID(ctx context.Context, id int) (users.Domain, error) {
	data := User{}
	if err := repo.DBConn.Where("id=?", id).First(&data).Error; err != nil {
		return users.Domain{}, err
	}
	domain := data.ToDomain()
	type Relation struct {
		MembershipName string
	}
	relation := Relation{}
	repo.DBConn.Table("users").Select("membership_types.name AS membership_name").Joins("LEFT JOIN membership_types ON users.membership_type_id = membership_types.id").Scan(&relation)
	domain.MembershipName = relation.MembershipName
	return domain, nil
}
