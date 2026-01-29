package feishu_task

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larktask "github.com/larksuite/oapi-sdk-go/v3/service/task/v2"
)

func GetAttachmentDetailsFromTask(taskID string, client *lark.Client) ([]string, []string, error) {
	req := larktask.NewListAttachmentReqBuilder().
		PageSize(50).
		ResourceType(`task`).
		ResourceId(taskID).
		Build()
	resp, err := client.Task.V2.Attachment.List(context.Background(), req)
	if err != nil {
		return nil, nil, err
	}
	if !resp.Success() {
		return nil, nil, fmt.Errorf("请求失败: %s", resp.CodeError.Msg)
	}
	var nameList []string
	var urlList []string
	for _, attachment := range resp.Data.Items {
		if attachment.Name != nil {
			nameList = append(nameList, *attachment.Name)
		}
		if attachment.Url != nil {
			urlList = append(urlList, *attachment.Url)
		}
	}
	return nameList, urlList, nil
}

// DownloadResult 单个文件下载结果
type DownloadResult struct {
	OriginalName string // 原始文件名
	FilePath     string // 实际保存路径
	Size         int64
	Err          error
}

func DownloadAttachments(nameList, urlList []string, destDir string) ([]DownloadResult, error) {
	if len(nameList) != len(urlList) {
		return nil, fmt.Errorf("文件名列表(%d)与URL列表(%d)长度不匹配", len(nameList), len(urlList))
	}
	if len(urlList) == 0 {
		return nil, nil
	}
	if destDir == "" {
		destDir = "./downloads"
	}
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return nil, fmt.Errorf("创建目录失败: %w", err)
	}

	results := make([]DownloadResult, len(urlList))
	var mu sync.Mutex
	var wg sync.WaitGroup

	semaphore := make(chan struct{}, 5)

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	for i, fileURL := range urlList {
		wg.Add(1)
		go func(idx int, url string) {
			defer wg.Done()

			semaphore <- struct{}{}        // 获取许可
			defer func() { <-semaphore }() // 释放许可

			fileName := nameList[idx]
			result := DownloadResult{
				OriginalName: fileName,
			}

			// 获取唯一文件路径
			filePath := getUniqueFilePath(filepath.Join(destDir, fileName))

			// 下载单个文件
			size, err := downloadSingle(client, url, filePath)
			result.FilePath = filePath
			result.Size = size
			result.Err = err

			mu.Lock()
			results[idx] = result
			mu.Unlock()

			if err != nil {
				fmt.Printf("❌ [%s] 失败: %v\n", fileName, err)
			} else {
				fmt.Printf("✅ [%s] 完成 (%.2f MB)\n", fileName, float64(size)/(1024*1024))
			}
		}(i, fileURL)
	}

	wg.Wait()
	return results, nil
}

// downloadSingle 单个文件下载（内部函数，不暴露）
func downloadSingle(client *http.Client, fileURL, filePath string) (int64, error) {
	req, err := http.NewRequest("GET", fileURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return 0, fmt.Errorf("创建文件失败: %w", err)
	}
	defer out.Close()

	written, err := io.Copy(out, resp.Body)
	if err != nil {
		return 0, fmt.Errorf("写入失败: %w", err)
	}
	if err := out.Sync(); err != nil {
		return written, err
	}

	return written, nil
}

func getUniqueFilePath(filePath string) string {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return filePath
	}

	ext := filepath.Ext(filePath)
	base := filePath[:len(filePath)-len(ext)]

	for i := 1; i < 1000; i++ {
		newPath := fmt.Sprintf("%s(%d)%s", base, i, ext)
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			return newPath
		}
	}
	return fmt.Sprintf("%s_%d%s", base, time.Now().Unix(), ext)
}

// func DownloadAttachment(nameList, urlList []string, destDir string) (string, error) {
// 	if destDir == "" {
// 		destDir = "."
// 	}
// 	if err := os.MkdirAll(destDir, 0755); err != nil {
// 		return "", fmt.Errorf("创建目录失败: %w", err)
// 	}
// 	for i, fileURL := range urlList {
// 		fileName := nameList[i]
// 		filePath := filepath.Join(destDir, fileName)
// 		if _, err := os.Stat(filePath); err == nil {
// 			ext := filepath.Ext(fileName)
// 			base := fileName[:len(fileName)-len(ext)]
// 			fileName = fmt.Sprintf("%s_%d%s", base, time.Now().Unix(), ext)
// 			filePath = filepath.Join(destDir, fileName)
// 		}
// 		client := &http.Client{
// 			Timeout: 120 * time.Second,
// 		}
// 		resp, err := client.Get(fileURL)
// 		if err != nil {
// 			return "", fmt.Errorf("请求失败: %w", err)
// 		}
// 		defer resp.Body.Close()

// 		if resp.StatusCode != http.StatusOK {
// 			return "", fmt.Errorf("服务器返回错误状态码: %d", resp.StatusCode)
// 		}
// 		// 创建本地 file
// 		out, err := os.Create(filePath)
// 		if err != nil {
// 			return "", fmt.Errorf("创建文件失败: %w", err)
// 		}
// 		defer out.Close()
// 		// 写入文件内容
// 		written, err := io.Copy(out, resp.Body)
// 		if err != nil {
// 			return "", fmt.Errorf("写入文件失败: %w", err)
// 		}
// 		fmt.Printf("下载完成: %s (大小: %.2f MB)\n", filePath, float64(written)/(1024*1024))
// 	}
// 	return "", nil
// }
