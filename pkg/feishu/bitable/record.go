package feishu_bitable

import (
	"context"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

func GetUserNameFromUnionID(unionID string, client *lark.Client) (string, error) {
	req := larkcontact.NewGetUserReqBuilder().
		UserId(unionID).
		UserIdType(`union_id`).
		Build()

	resp, err := client.Contact.V3.User.Get(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("请求失败: %w", err)
	}
	if !resp.Success() {
		return "", fmt.Errorf("请求失败: %s", resp.CodeError.Msg)
	}
	return *resp.Data.User.Name, nil
}
