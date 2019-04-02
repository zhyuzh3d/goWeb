package api

import (
	"app/tool"
	"app/util"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type sendRegVerifyMailReqDS struct {
	Email string
}

//SendRegVerifyMail 注册接口处理函数
func SendRegVerifyMail(w http.ResponseWriter, r *http.Request) {
	ds := sendRegVerifyMailReqDS{}
	json.NewDecoder(r.Body).Decode(&ds)

	mailRe, _ := regexp.Compile(`^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	if !mailRe.MatchString(ds.Email) {
		util.WWrite(w, 1, "邮箱格式错误。", nil)
		return
	}

	//检查是否存在，如果已经存在且时间小于1分钟就就不再发送
	dbc := tool.MongoDBCLient.Database("myweb").Collection("regVerify")
	var u bson.M
	dbc.FindOne(context.TODO(), bson.M{"Email": ds.Email}).Decode(&u)
	now := time.Now().Unix()
	if u["Ts"] != nil && now-u["Ts"].(int64) < 60 {
		util.WWrite(w, 1, "请不要重复发送邮件。", nil)
		return
	}

	//生成随机6位数，并发送
	code := rand.Intn(899999) + 100000
	s := strconv.Itoa(code)
	err := tool.SendMail(ds.Email, "您在www.myweb.com的注册码是:"+s, "来自Myweb的注册验证码")
	if err != nil {
		util.WWrite(w, 1, "发送邮件失败。", nil)
		fmt.Println(err)
		return
	}

	//删除原有数据，创建新数据
	dbc.DeleteOne(context.TODO(), bson.M{"Email": ds.Email})
	dt := bson.M{"Code": s, "Email": ds.Email, "Ts": now}
	_, err = dbc.InsertOne(context.TODO(), dt)
	if err != nil {
		util.WWrite(w, 1, "写入数据库出错。", nil)
		fmt.Println(err)
	} else {
		util.WWrite(w, 0, "发送成功，请检查邮箱。", nil)
	}
	return
}
