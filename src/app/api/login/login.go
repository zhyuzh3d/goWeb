package login

import (
	"app/tool"
	"app/util"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

//ReqDS 注册接口的请求数据格式
type ReqDS struct {
	Email string
	Pw    string
}

//HandleFunc 注册接口处理函数
func HandleFunc(w http.ResponseWriter, r *http.Request) {
	ds := ReqDS{}
	json.NewDecoder(r.Body).Decode(&ds)

	// //访问数据集
	dbc := tool.MongoDBCLient.Database("myweb").Collection("user")

	//验证用户邮箱是否与用户名匹配
	var u bson.M
	dbc.FindOne(context.TODO(), bson.M{"Email": ds.Email}).Decode(&u)
	if u["Pw"] == ds.Pw {
		util.WWrite(w, 0, "登录成功", u["_id"])
	} else {
		util.WWrite(w, 1, "邮箱与用户名不匹配", nil)
	}
	return
}
