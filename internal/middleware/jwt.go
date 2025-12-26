package middleware

import (
	"context"
	"fmt"
	"server/internal/infra"
	"server/internal/pkg/apperror"
	"server/internal/pkg/response"
	"server/internal/pkg/token"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("token")
		if accessToken == "" {
			err := apperror.New(apperror.Unauthorized, "没有有效的认证信息")
			response.Fail(c, err)
			c.Abort()
			return
		}

		// check token valid
		userInfo, err := token.ValidateToken(accessToken, infra.AppConfig.Jwt.AccessSignedKey)
		if err != nil {
			err := apperror.New(apperror.InvalidToken, "无效的token")
			response.Fail(c, err)
			c.Abort()
			return
		}

		// check token in whitelist
		if infra.RDS != nil {
			ctx := context.Background()
			accessTokenKey := fmt.Sprintf("whitelist:access_token:%s", accessToken)
			exists, err := infra.RDS.Exists(ctx, accessTokenKey).Result()
			if err != nil || exists == 0 {
				err := apperror.New(apperror.InvalidToken, "无效的token")
				response.Fail(c, err)
				c.Abort()
				return
			}
		}

		// check token exist
		c.Set("username", userInfo.Username)
		c.Set("userId", userInfo.ID)

		c.Next()
	}
}
