package ext

import (
	"app/tool"
	"app/util"
	"context"
	"net/http"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MiddleWare 文件服务中间件
func MiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//仅对.html文件处理
		htmlRe, _ := regexp.Compile(`^.+\.html[\?]*.*$`)
		if !htmlRe.MatchString(r.URL.String()) {
			h.ServeHTTP(w, r)
			return
		}

		//检查Cookie中的Uid是否合法
		loginCheck(w, r)
		//文件服务
		h.ServeHTTP(w, r)
	})
}

//MiddleWareAPI API中间件:检查Uid和Token的合理性
func MiddleWareAPI(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//检查Cookie中的Uid是否合法
		loginCheck(w, r)
		//API服务
		next(w, r)
	})
}

//loginCheck 检查Cookie中的Uid是否合法
func loginCheck(w http.ResponseWriter, r *http.Request) {
	//获取Token
	token, _ := r.Cookie("Token")
	if token == nil {
		util.DelCookie(w, "Uid")
		return
	}
	tv := token.Value
	if tv == "" {
		util.DelCookie(w, "Uid")
		return
	}

	//如果token匹配就向Cookie添加"Uid"
	ctoken := tool.MongoDBCLient.Database("myweb").Collection("token")
	var t bson.M
	ctoken.FindOne(context.TODO(), bson.M{"Token": tv}).Decode(&t)
	uid := t["Id"]
	if uid != nil {
		uids := uid.(primitive.ObjectID).Hex()
		util.SetCookie(w, "Uid", uids)
	} else {
		util.DelCookie(w, "Uid")
	}
}
