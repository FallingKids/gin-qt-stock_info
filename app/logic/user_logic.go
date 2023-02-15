package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/dao"
	"github.com/gin-qt-business/app/type/data"
	pinyin "github.com/mozillazg/go-pinyin"
)

type User struct{}

func (u *User) LoginByPassword(ctx *gin.Context) (data.LoginByPasswordRes, error) {
	res := data.LoginByPasswordRes{
		Token: "",
	}
	return res, nil
}

func (u *User) Register(input data.RegisterReq) (data.RegisterRes, error) {
	// uid 的命名规则为中文名称的，例如 胡奕舟 -> huyz
	pinyinArgs := pinyin.NewArgs()
	pinyinArgs.Style = pinyin.Normal
	pinyinStr := pinyin.Pinyin(input.Username, pinyinArgs)
	var uid string
	for _, s := range pinyinStr {
		if len(uid) != 0 {
			for _, c := range s {
				uid += c
			}
		}
		uid += string(s[0])
	}
	_, err := dao.AddUser(dao.User{
		Uid:      uid,
		Username: input.Username,
		Password: input.Password,
		Phone:    input.Phone,
	})
	if err != nil {
		return data.RegisterRes{IsOk: false}, err
	}
	return data.RegisterRes{IsOk: true}, nil
}
