package main

import (
	"fmt"
	"testing"

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
