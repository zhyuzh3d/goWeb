package api

import (
	"app/ext"
	"app/tool"
	"app/util"
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type registerReqDS struct {
	Email string
	Pw    string
	Code  string
}

//Register 注册接口处理函数
func Register(w http.ResponseWriter, r *http.Request) {
	ds := registerReqDS{}
	json.NewDecoder(r.Body).Decode(&ds)

	//格式检查
	pwRe, _ := regexp.Compile(`^[0-9a-z]{32}$`)
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
	dbcv := tool.MongoDBCLient.Database("myweb").Collection("regVerify")
	var v bson.M
	dbcv.FindOne(context.TODO(), bson.M{"Email": ds.Email}).Decode(&v)
	if v["Code"] == nil || v["Code"] != ds.Code {
		util.WWrite(w, 1, "验证码错误。", nil)
		return
	}

	//访问数据集
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	dbc := tool.MongoDBCLient.Database("myweb").Collection("user")
	defer cancel()

	//验证用户邮箱是否存在
	count, err := dbc.CountDocuments(context.TODO(), bson.M{"Email": ds.Email})
	if err != nil {
		util.WWrite(w, 1, "读取数据库失败。", nil)
		return
	}
	if count > 0 {
		util.WWrite(w, 1, "邮箱已存在。", nil)
		return
	}

	//写入数据库
	u := bson.M{"Email": ds.Email, "Pw": ds.Pw, "Ts": time.Now().Unix()}
	res, err := dbc.InsertOne(ctx, u)
	if err != nil {
		util.WWrite(w, 1, "写入数据库失败。", nil)
		return
	}
	uids := res.InsertedID.(primitive.ObjectID).Hex()

	datas := ext.NewToken(w, uids)
	util.WWrite(w, 0, "注册成功。", datas)
	return
}
