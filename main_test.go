package main

import (
	"fmt"
	"testing"

	feishu_bitable "github.com/yeqichao230516/jtyl-library/pkg/feishu/bitable"
	feishu_task "github.com/yeqichao230516/jtyl-library/pkg/feishu/task"
	"github.com/yeqichao230516/jtyl-library/pkg/system"
)

func TestMain(t *testing.T) {
	err := feishu_task.DeleteTask("063318ee-7f8a-4076-924c-efb4815fc791", system.FeiShu("cli_a81807b812b7901c", "wGTNLAxJiZBCoBvht4b7UbeBmSkWprYw"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Task deleted successfully")
}

func TestCompleteTask(t *testing.T) {
	err := feishu_task.CompleteTask("31717367-daad-46eb-91f7-1b2498ee44ed", system.FeiShu("cli_a81807b812b7901c", "wGTNLAxJiZBCoBvht4b7UbeBmSkWprYw"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Task completed successfully")
}

func TestGetTask(t *testing.T) {
	task, err := feishu_task.GetTask("31717367-daad-46eb-91f7-1b2498ee44ed", system.FeiShu("cli_a81807b812b7901c", "wGTNLAxJiZBCoBvht4b7UbeBmSkWprYw"))
	if err != nil {
		t.Fatal(err)
	}
	for _, value := range task.CustomFields {
		if *value.Name == "任务进度" {
			fmt.Printf("任务进度: %v\n", *value.NumberValue)
		}
	}
	// data, _ := json.Marshal(task.CustomFields)
	// fmt.Printf("Task details:\n%s\n", string(data))
}

func TestGetComments(t *testing.T) {
	comments, err := feishu_task.GetComments("ba011d87-6663-43f8-a639-ab402987593d", system.FeiShu("cli_a81807b812b7901c", "wGTNLAxJiZBCoBvht4b7UbeBmSkWprYw"))
	if err != nil {
		t.Fatal(err)
	}
	for _, comment := range comments.Items {
		fmt.Printf("Comment: %s\n", *comment.Creator.Id)
		fmt.Printf("Comment: %s\n", *comment.Content)
	}
}

func TestUpdateRecord(t *testing.T) {
	err := feishu_bitable.UpdateRecord("V8AxbmAOXapXQesLIIFcJbkunae", "tblFuFNqnSyQfDBl", "recv9n66O7A1h3", map[string]any{"任务评论": "人员1：测试评论详情\n人员2：你好"}, system.FeiShu("cli_a81807b812b7901c", "wGTNLAxJiZBCoBvht4b7UbeBmSkWprYw"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Record updated successfully")
}
