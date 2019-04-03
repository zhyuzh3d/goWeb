package api

import (
	"app/tool"
	"app/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
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

		//创建token并写入数据库
		token, _ := uuid.NewV4()
		tokens := token.String()
		ctoken := tool.MongoDBCLient.Database("myweb").Collection("token")
		du := bson.M{"Token": tokens, "Id": u["_id"], "Ts": time.Now().Unix()}
		ctoken.InsertOne(context.TODO(), du)

		//返回id，写入Token和Uid
		util.SetCookie(w, "Uid", uids)
		util.SetCookie(w, "Token", tokens)
		fmt.Println(tokens, uids)

		util.WWrite(w, 0, "登录成功", u["_id"])
	} else {
		util.WWrite(w, 1, "邮箱与用户名不匹配", nil)
	}

	return
}
