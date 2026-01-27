package main

import (
	"testing"

	feishu_task "github.com/yeqichao230516/jtyl-library/pkg/feishu/task"
	"github.com/yeqichao230516/jtyl-library/pkg/system"
)

func TestMain(t *testing.T) {
	id, err := feishu_task.GetCustomFieldIDSingleByGuid("134ba936-864e-4878-a82e-67024e38be34", "项目1", system.FeiShu("cli_a81807b812b7901c", "wGTNLAxJiZBCoBvht4b7UbeBmSkWprYw"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

//89ebd072-c62b-45b5-8d6c-735f18ae982a
