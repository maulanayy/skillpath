package usecase

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"skillpath/model"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) Login(req model.Login) (model.UserInfo, error) {

	user, err := uc.userDB.GetByEmail(req.Email)

	if err != nil {
		fmt.Println("ERROR GET BY EMAIl")
		return model.UserInfo{}, err
	}

	if err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password)); err != nil {
		return model.UserInfo{}, errors.New("Password does not match..")
	}
	now := time.Now()
	h := sha256.New()
	c := strconv.FormatInt(user.ID, 10) + user.Password + strconv.FormatInt(now.Unix(), 10)
	h.Write([]byte(c))
	token := fmt.Sprintf("%x", h.Sum(nil))

	user.Info.TokenGeneratedTime, user.Token = time.Now().Unix(), token

	data := model.UserInfo{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
		Name:  user.Info.Name,
	}

	return data, nil
}

func (uc *usecase) Create(req model.User) error {
	hashPassword := []byte(string(req.Password))
	hashPassword, err := bcrypt.GenerateFromPassword(hashPassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashPassword)

	return uc.userDB.Create(&req)
}
