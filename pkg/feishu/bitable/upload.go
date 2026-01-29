package feishu_bitable

import (
	"context"
	"fmt"
	"os"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkdrive "github.com/larksuite/oapi-sdk-go/v3/service/drive/v1"
)

func UploadFileToBitable(appToken, fileName string, fileSize int, file *os.File, client *lark.Client) (string, error) {
	req := larkdrive.NewUploadAllMediaReqBuilder().
		Body(larkdrive.NewUploadAllMediaReqBodyBuilder().
			FileName(fileName).
			ParentType(`bitable_file`).
			ParentNode(appToken).
			Size(fileSize).
			File(file).
			Build()).
		Build()
	resp, err := client.Drive.V1.Media.UploadAll(context.Background(), req)
	if err != nil {
		return "", err
	}
	if !resp.Success() {
		return "", fmt.Errorf("请求失败: %s", resp.CodeError.Msg)
	}
	return *resp.Data.FileToken, nil
}
