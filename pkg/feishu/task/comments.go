package feishu_task

import (
	"context"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larktask "github.com/larksuite/oapi-sdk-go/v3/service/task/v2"
)

func GetComments(guid string, client *lark.Client) (*larktask.ListCommentRespData, error) {
	req := larktask.NewListCommentReqBuilder().
		PageSize(50).
		ResourceType(`task`).
		ResourceId(guid).
		Direction(`asc`).
		Build()
	resp, err := client.Task.V2.Comment.List(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if !resp.Success() {
		return nil, fmt.Errorf("请求失败: %s", resp.CodeError.Msg)
	}
	return resp.Data, nil
}
