package api

import (
	"app/tool"
	"app/util"
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type resetPwReqDS struct {
	Email string
	Pw    string
	Code  string
}

//ResetPw 重置密码的接口
func ResetPw(w http.ResponseWriter, r *http.Request) {
	ds := resetPwReqDS{}
	json.NewDecoder(r.Body).Decode(&ds)

	//格式检查
	pwRe, _ := regexp.Compile(`^[0-9a-zA-Z_@]{6,18}$`)
	if !pwRe.MatchString(ds.Pw) {
		util.WWrite(w, 1, "密码格式错误。", nil)
		return
	}

	codeRe, _ := regexp.Compile(`^[0-9]{6}$`)
	if !codeRe.MatchString(ds.Code) {
		util.WWrite(w, 1, "验证码格式错误", nil)
		return
	}

	mailRe, _ := regexp.Compile(`^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	if !mailRe.MatchString(ds.Email) {
		util.WWrite(w, 1, "邮箱格式错误。", nil)
		return
	}

	//检查验证码是否正确
	dbcv := tool.MongoDBCLient.Database("myweb").Collection("rstPwVerify")
	var v bson.M
	dbcv.FindOne(context.TODO(), bson.M{"Email": ds.Email}).Decode(&v)
	if v["Code"] == nil || v["Code"] != ds.Code {
		util.WWrite(w, 1, "验证码错误。", nil)
		return
	}

	//更新数据库中的密码
	dbc := tool.MongoDBCLient.Database("myweb").Collection("user")
	filter := bson.M{"Email": ds.Email}
	up := bson.M{"$set": bson.M{"Pw": ds.Pw, "TsRst": time.Now().Unix()}}
	_, err := dbc.UpdateOne(context.TODO(), filter, up)
	if err != nil {
		util.WWrite(w, 1, "写入数据库失败。", nil)
		return
	}

	//返回修改的账号
	var nu bson.M
	dbc.FindOne(context.TODO(), bson.M{"Email": ds.Email}).Decode(&nu)
	util.WWrite(w, 0, "修改成功。", nu["_id"])
	return
}
