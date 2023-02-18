package logic

import (
	"fmt"
	"time"

	"github.com/gin-qt-business/app/dao"
	"github.com/gin-qt-business/app/type/data"
	"github.com/gin-qt-business/app/utils"
	pinyin "github.com/mozillazg/go-pinyin"
	"github.com/spf13/viper"
)

type UserService struct{}

func (u *UserService) LoginByPassword(input data.LoginByPasswordReq) (data.LoginByPasswordRes, error) {
	res := data.LoginByPasswordRes{
		Token: "",
	}
	return res, nil
}

func (u *UserService) Register(input data.RegisterReq) (data.RegisterRes, error) {
	key := viper.Get("app.aes_key").(string)
	iv := viper.Get("app.aes_iv").(string)

	// uid 的命名规则为中文名称的，例如 胡奕舟 -> huyz
	pinyinArgs := pinyin.NewArgs()
	pinyinArgs.Style = pinyin.Normal
	pinyinStr := pinyin.Pinyin(input.Username, pinyinArgs)
	var uid string
	for _, s := range pinyinStr {
		if len(uid) == 0 {
			uid += string(s[0])
		} else {
			uid += string(s[0][0])
		}
	}

	password, err := utils.Decrypt(input.Password, []byte(key), []byte(iv))
	if err != nil {
		return data.RegisterRes{IsOk: false}, err
	}

	user := dao.UserDao{
		Uid: uid,
	}

	err = checkUid(&user)
	if err != nil {
		return data.RegisterRes{IsOk: false}, err
	}

	user.Username = input.Username
	user.Password = string(password)
	user.Phone = input.Phone

	err = user.AddUser()
	if err != nil {
		return data.RegisterRes{IsOk: false}, err
	}
	return data.RegisterRes{IsOk: true}, nil
}

func checkUid(user *dao.UserDao) error {
	preUid := user.Uid
	err := user.GetUser()
	if err != nil {
		return err
	}
	if user.ID != 0 {
		user.Uid = preUid + fmt.Sprint(time.Now().Month()) + fmt.Sprint(time.Now().Day())
		checkUid(user)
	}
	return nil
}
