package ps

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"golang.org/x/time/rate"
	"strconv"
	"zhuhui.com/microkit/kit_provider/util"
)

/**
	Endpoint 定义Request和Response格式，
	并可以使用装饰器包装函数，以此来实现各种中间件嵌套
 */

type UserRequest struct {
	Uid    int `json:"uid"`
	Method string
	Token  string
}

type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEndponit(us IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		uname := ctx.Value("LoginUser")
		fmt.Println("current user:", uname)
		rs := "nothing"
		//var logger kitlog.Logger
		//{
		//	logger = kitlog.NewLogfmtLogger(os.Stdout)
		//	logger = kitlog.WithPrefix(logger, "mykit", "1.0")
		//	logger = kitlog.With(logger, "time", kitlog.DefaultTimestampUTC)
		//	logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
		//}
		if r.Method == "GET" {
			rs = us.GetName(r.Uid) + strconv.Itoa(util.ServicePort)
			//logger.Log("method", r.Method, "event", "get user", "userid", r.Uid)
		} else if r.Method == "DELETE" {
			err := us.DelUser(r.Uid)
			if err != nil {
				rs = err.Error()
			} else {
				rs = fmt.Sprintf("del user, id=%d", r.Uid)
			}
		} else {

		}

		return UserResponse{Result: rs}, nil
	}
}

/**
限流中间件
 */
func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, util.NewMyError(429, "too many request")
			}
			return next(ctx, request)
		}
	}
}

/**
日志中间件
 */
func UserServiceLogMiddleware(logger kitlog.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(UserRequest)
			logger.Log("method", r.Method, "event", "get user", "userid", r.Uid)
			return next(ctx, request)
		}
	}
}

/**
Token中间件
 */
func CheckTokenMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(UserRequest)
			uc := UserClaim{}
			getToken, err := jwt.ParseWithClaims(r.Token, &uc, func(token *jwt.Token) (i interface{}, e error) {
				return []byte(secKey), nil
			})
			if getToken != nil && getToken.Valid {
				name := getToken.Claims.(*UserClaim).Uname
				fmt.Println()
				newCtx := context.WithValue(ctx, "LoginUser", name)
				return next(newCtx, request)
			} else {
				return nil, util.NewMyError(403, "error token")
			}
		}
	}
}
