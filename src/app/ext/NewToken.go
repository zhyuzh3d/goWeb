package ext

import (
	"app/tool"
	"app/util"
	"context"
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

//NewToken 创建新token，返回一个基本的用户信息Token和Uid，用于用户手工登录注册和重置时候使用
func NewToken(w http.ResponseWriter, uids string) string {
	//创建token
	token, _ := uuid.NewV4()
	tokens := token.String()

	//删除旧的,创建新的
	coll := tool.MongoDBCLient.Database("myweb").Collection("token")
	du := bson.M{"Token": tokens, "Id": uids, "Ts": time.Now().Unix()}
	du2 := bson.M{"Id": uids}
	coll.DeleteMany(context.TODO(), du2)
	coll.InsertOne(context.TODO(), du)

	//返回对象
	data := map[string]string{
		"Token": tokens,
		"Uid":   uids}

	datas, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	//返回id，写入Token和Uid
	util.SetCookie(w, "Uid", uids)
	util.SetCookie(w, "Token", tokens)
	return string(datas)
}
