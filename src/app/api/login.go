package api

import (
	"app/ext"
	"app/tool"
	"app/util"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type loginReqDS struct {
	Email string
	Pw    string
}

//Login 注册接口处理函数
func Login(w http.ResponseWriter, r *http.Request) {
	ds := loginReqDS{}
	json.NewDecoder(r.Body).Decode(&ds)

	// //访问数据集
	dbc := tool.MongoDBCLient.Database("myweb").Collection("user")

	//验证用户邮箱是否与用户名匹配
	var u bson.M
	dbc.FindOne(context.TODO(), bson.M{"Email": ds.Email}).Decode(&u)
	if u["Pw"] == ds.Pw {
		uids := u["_id"].(primitive.ObjectID).Hex()

		datas := ext.NewToken(w, uids)
		util.WWrite(w, 0, "登录成功", datas)
	} else {
		util.WWrite(w, 1, "邮箱与用户名不匹配", nil)
	}
	return
}
