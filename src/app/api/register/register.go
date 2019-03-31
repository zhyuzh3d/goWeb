package register

import (
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

//ReqDS 注册接口的请求数据格式
type ReqDS struct {
	Email string
	Pw    string
}

//HandleFunc 注册接口处理函数
func HandleFunc(w http.ResponseWriter, r *http.Request) {
	ds := ReqDS{}
	json.NewDecoder(r.Body).Decode(&ds)

	mailRe, _ := regexp.Compile(`^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	pwRe, _ := regexp.Compile(`^[0-9a-zA-Z_@]{6,18}$`)

	if !pwRe.MatchString(ds.Pw) {
		util.WWrite(w, 1, "密码格式错误。", nil)
		return
	}
	if !mailRe.MatchString(ds.Email) {
		util.WWrite(w, 1, "邮箱格式错误。", nil)
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
	res, err := dbc.InsertOne(ctx, bson.M{"Email": ds.Email, "Pw": ds.Pw})
	if err != nil {
		util.WWrite(w, 1, "写入数据库失败。", nil)
		return
	}
	d := res.InsertedID.(primitive.ObjectID).Hex()
	util.WWrite(w, 0, "注册成功。", d)
	return
}
