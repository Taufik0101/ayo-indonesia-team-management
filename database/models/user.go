package models

import (
	"fmt"
	"gin-ayo/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

type User struct {
	BaseModel
	Name     string         `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Email    string         `json:"email" gorm:"column:email;type:varchar(255);not null"`
	Password *string        `json:"password" gorm:"column:password;type:text;null"`
	Role     utils.UserType `json:"role" gorm:"column:role;type:role_types;not null;default: 'user'"`
}

func (*User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	if u.Password != nil && *u.Password != "" {
		pass := strings.TrimSpace(*u.Password)
		u.Password = &pass

		err = u.ValidatePassword()
		if err != nil {
			return err
		}

		if err := u.HashPassword(); err != nil {
			return err
		}
	}

	return nil
}

// BeforeUpdate Prepare user for update user
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	u.Name = strings.TrimSpace(u.Name)

	return nil
}

// HashPassword Hash user password with bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	strHashed := string(hashedPassword)
	u.Password = &strHashed

	return nil
}

// SanitizePassword Sanitize user password
func (u *User) SanitizePassword() {
	u.Password = nil
}

// ComparePasswords Compare user password and payload
func (u *User) ComparePasswords(password string) error {
	if u.Password != nil && *u.Password != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(password)); err != nil {
			return err
		}
	}

	return nil
}

// ValidatePassword validatePassword checks if the given password meets the specified criteria.
func (u *User) ValidatePassword() error {
	// Check password length
	if len(*u.Password) < 8 {
		return fmt.Errorf("password length must be at least 8 characters")
	}

	// Check for at least 1 lowercase letter
	re := regexp.MustCompile(`[a-z]`)
	if !re.MatchString(*u.Password) {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	// Check for at least 1 uppercase letter
	re = regexp.MustCompile(`[A-Z]`)
	if !re.MatchString(*u.Password) {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	// Check for at least 1 digit
	re = regexp.MustCompile(`[0-9]`)
	if !re.MatchString(*u.Password) {
		return fmt.Errorf("password must contain at least one digit")
	}

	// Check for at least 1 special character
	re = regexp.MustCompile(`[!@#$%^&?~]`)
	if !re.MatchString(*u.Password) {
		return fmt.Errorf("password must contain at least one special character: !@#$%%^&?~")
	}

	return nil
}
