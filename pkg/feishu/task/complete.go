package feishu_task

import (
	"context"
	"fmt"
	"strconv"
	"time"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larktask "github.com/larksuite/oapi-sdk-go/v3/service/task/v2"
)

func CompleteTask(guid string, client *lark.Client) error {
	req := larktask.NewPatchTaskReqBuilder().
		TaskGuid(guid).
		Body(larktask.NewPatchTaskReqBodyBuilder().
			Task(larktask.NewInputTaskBuilder().
				CompletedAt(strconv.FormatInt(time.Now().UnixMilli(), 10)).
				Build()).
			UpdateFields([]string{`completed_at`}).
			Build()).
		Build()
	resp, err := client.Task.V2.Task.Patch(context.Background(), req)
	if err != nil {
		return err
	}
	if !resp.Success() {
		return fmt.Errorf("请求失败: %s", resp.CodeError.Msg)
	}
	return nil
}
