package MyWxSdk

import (
	"encoding/base64"
	"fmt"
	"os"
	"regexp"
)

// TConfig 配置
var (
	OriginId       string // 原始ID
	AppId          string // 应用ID
	AppSecret      string // 应用密钥
	Token          string // 令牌
	EncodingAESKey []byte // 消息加解密密钥

)

// Initialize 配置并初始化
func Initialize(originId, appId, appSecret, token, encodingAESKey string) {

	//sugarLogger = logger.Sugar()
	if matched, err := regexp.MatchString("^gh_[0-9a-f]{12}$", originId); err != nil || !matched {

	}
	if matched, err := regexp.MatchString("^wx[0-9a-f]{16}$", appId); err != nil || !matched {

	}
	if matched, err := regexp.MatchString("^[0-9a-f]{32}$", appSecret); err != nil || !matched {

	}
	if matched, err := regexp.MatchString("^[0-9a-zA-Z]{3,32}$", token); err != nil || !matched {

	}
	if matched, err := regexp.MatchString("^[0-9a-zA-Z]{43}$", encodingAESKey); err != nil || !matched {

	}

	OriginId = originId   // 原始ID
	AppId = appId         // 应用ID
	AppSecret = appSecret // 应用密钥
	Token = token         // 令牌

	var err error
	EncodingAESKey, err = base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		fmt.Printf("appSecret config error: %s", err)
		os.Exit(1)
	}

	// refresh access token
	RefreshAccessToken(AppId, AppSecret)
}
