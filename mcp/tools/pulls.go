package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	
	"github.com/gitcode-org-com/gitcode-mcp/api"
)

// AddPullRequestTools 添加Pull Request相关工具到MCP服务器
func AddPullRequestTools(s *server.MCPServer, apiClient *api.GitCodeAPI) {
	// 列出Pull Requests
	listPRsTool := mcp.NewTool("list_pull_requests",
		mcp.WithDescription("列出仓库的Pull Requests"),
		mcp.WithString("owner",
			mcp.Required(),
			mcp.Description("仓库所有者"),
		),
		mcp.WithString("repo",
			mcp.Required(),
			mcp.Description("仓库名称"),
		),
	)
	s.AddTool(listPRsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		owner, _ := request.Params.Arguments["owner"].(string)
		repo, _ := request.Params.Arguments["repo"].(string)
		
		prs, err := apiClient.Pulls.ListPullRequests(owner, repo)
		if err != nil {
			return nil, fmt.Errorf("获取Pull Requests列表失败: %w", err)
		}
		return FormatJSONResult(prs)
	})
	
	// 获取Pull Request
	getPRTool := mcp.NewTool("get_pull_request",
		mcp.WithDescription("获取特定Pull Request的详细信息"),
		mcp.WithString("owner",
			mcp.Required(),
			mcp.Description("仓库所有者"),
		),
		mcp.WithString("repo",
			mcp.Required(),
			mcp.Description("仓库名称"),
		),
		mcp.WithNumber("pull_number",
			mcp.Required(),
			mcp.Description("Pull Request编号"),
		),
	)
	s.AddTool(getPRTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		owner, _ := request.Params.Arguments["owner"].(string)
		repo, _ := request.Params.Arguments["repo"].(string)
		prNumber, _ := request.Params.Arguments["pull_number"].(float64)
		
		pr, err := apiClient.Pulls.GetPullRequest(owner, repo, int(prNumber))
		if err != nil {
			return nil, fmt.Errorf("获取Pull Request详情失败: %w", err)
		}
		return FormatJSONResult(pr)
	})
	
	// 创建Pull Request
	createPRTool := mcp.NewTool("create_pull_request",
		mcp.WithDescription("创建新Pull Request"),
		mcp.WithString("owner",
			mcp.Required(),
			mcp.Description("仓库所有者"),
		),
		mcp.WithString("repo",
			mcp.Required(),
			mcp.Description("仓库名称"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Pull Request标题"),
		),
		mcp.WithString("head",
			mcp.Required(),
			mcp.Description("源分支"),
		),
		mcp.WithString("base",
			mcp.Required(),
			mcp.Description("目标分支"),
		),
		mcp.WithString("body",
			mcp.Description("Pull Request内容"),
		),
	)
	s.AddTool(createPRTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		owner, _ := request.Params.Arguments["owner"].(string)
		repo, _ := request.Params.Arguments["repo"].(string)
		title, _ := request.Params.Arguments["title"].(string)
		head, _ := request.Params.Arguments["head"].(string)
		base, _ := request.Params.Arguments["base"].(string)
		body, _ := request.Params.Arguments["body"].(string)
		
		pr, err := apiClient.Pulls.CreatePullRequest(owner, repo, title, head, base, body)
		if err != nil {
			return nil, fmt.Errorf("创建Pull Request失败: %w", err)
		}
		return FormatJSONResult(pr)
	})
} 