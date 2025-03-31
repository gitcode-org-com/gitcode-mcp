package tools

import (
	"encoding/json"
	"fmt"
	
	"github.com/mark3labs/mcp-go/mcp"
)

// FormatJSONResult 将数据格式化为JSON结果
func FormatJSONResult(data interface{}) (*mcp.CallToolResult, error) {
	// 将数据转换为JSON字符串
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("JSON编码失败: %w", err)
	}
	
	// 使用 NewToolResultText 创建结果
	return mcp.NewToolResultText(string(jsonBytes)), nil
} 