package system

import lark "github.com/larksuite/oapi-sdk-go/v3"

func FeiShu(appID, appSecret string) *lark.Client {
	return lark.NewClient(appID, appSecret)
}
