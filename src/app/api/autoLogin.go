package api

import (
	"app/tool"
	"app/util"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type autoLoginReqDS struct {
	Email string
}

//AutoLogin 注册接口处理函数
func AutoLogin(w http.ResponseWriter, r *http.Request) {
	ds := autoLoginReqDS{}
	json.NewDecoder(r.Body).Decode(&ds)

	//直接信任Cookie中的Uid
	uid, _ := r.Cookie("Uid")

	//没登录返回空
	if uid == nil || uid.Value == "" {
		util.WWrite(w, 1, "自动登录失败。", nil)
		return
	}

	//登录成功返回对象
	var u bson.M
	coll := tool.MongoDBCLient.Database("myweb").Collection("user")
	idobj, err := primitive.ObjectIDFromHex(uid.Value)
	if err != nil {
		util.WWrite(w, 1, "自动登录Cookie.Uid异常。", nil)
		return
	}
	coll.FindOne(context.TODO(), bson.M{"_id": idobj}).Decode(&u)

	data := map[string]string{
		"Email": u["Email"].(string),
		"Uid":   uid.Value}
	datas, err := json.Marshal(data)
	if err != nil {
		util.WWrite(w, 1, "自动登录数据库内容异常。", nil)
		return
	}

	util.WWrite(w, 0, "自动登录成功。", string(datas))
	return
}
