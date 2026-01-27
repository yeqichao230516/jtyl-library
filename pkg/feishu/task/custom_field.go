package feishu_task

import (
	"context"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larktask "github.com/larksuite/oapi-sdk-go/v3/service/task/v2"
)

func GetCustomFieldIDsMultipleByGuid(guid string, fields []string, client *lark.Client) ([]string, error) {
	req := larktask.NewGetCustomFieldReqBuilder().
		CustomFieldGuid(guid).
		Build()

	resp, err := client.Task.V2.CustomField.Get(context.Background(), req)
	if err != nil {
		return nil, err
	}

	if !resp.Success() {
		return nil, fmt.Errorf("请求失败: %s", resp.CodeError.Msg)
	}
	var result []string
	for _, options := range resp.Data.CustomField.MultiSelectSetting.Options {
		for _, field := range fields {
			if options.Name != nil && *options.Name == field {
				result = append(result, *options.Guid)
			}
		}
	}
	return result, nil
}

func GetCustomFieldIDSingleByGuid(guid string, field string, client *lark.Client) (string, error) {
	req := larktask.NewGetCustomFieldReqBuilder().
		CustomFieldGuid(guid).
		Build()

	resp, err := client.Task.V2.CustomField.Get(context.Background(), req)
	if err != nil {
		return "", err
	}

	if !resp.Success() {
		return "", fmt.Errorf("请求失败: %s", resp.CodeError.Msg)
	}
	for _, options := range resp.Data.CustomField.SingleSelectSetting.Options {
		if options.Name != nil && *options.Name == field {
			return *options.Guid, nil
		}
	}
	return "", nil
}
