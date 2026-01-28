package feishu_task

import (
	"context"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larktask "github.com/larksuite/oapi-sdk-go/v3/service/task/v2"
)

func GetTask(guid string, client *lark.Client) (*larktask.Task, error) {
	req := larktask.NewGetTaskReqBuilder().
		TaskGuid(guid).
		Build()
	resp, err := client.Task.V2.Task.Get(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if !resp.Success() {
		return nil, fmt.Errorf("请求失败: %s", resp.CodeError.Msg)
	}
	return resp.Data.Task, nil
}
