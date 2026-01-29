package feishu_address_book

import (
	"context"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

func GetUserNameByUnionID(unionID string, client *lark.Client) (string, error) {
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

func GetUserNameByOpenID(openID string, client *lark.Client) (string, error) {
	req := larkcontact.NewGetUserReqBuilder().
		UserId(openID).
		UserIdType(`open_id`).
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
