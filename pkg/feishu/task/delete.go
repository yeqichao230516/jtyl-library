package feishu_task

import (
	"context"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larktask "github.com/larksuite/oapi-sdk-go/v3/service/task/v2"
)

func DeleteTask(guid string, client *lark.Client) error {
	req := larktask.NewDeleteTaskReqBuilder().
		TaskGuid(guid).
		Build()

	resp, err := client.Task.V2.Task.Delete(context.Background(), req)

	if err != nil {
		return err
	}
	if !resp.Success() {
		return fmt.Errorf("请求失败: %s", resp.CodeError.Msg)
	}
	return nil
}
