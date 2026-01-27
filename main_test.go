package main

import (
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
