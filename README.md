# GitCode MCP Go Server

这是GitCode MCP服务器的Go语言实现版本，提供了GitCode API的标准MCP接口封装。

## 功能特点

- 完整支持GitCode API的主要功能
- 基于标准MCP协议实现，使用mark3labs/mcp-go SDK
- 支持STDIO和SSE两种传输方式
- 轻量级，响应速度快
- 并发处理能力强，适合高负载场景
- 模块化的代码结构，便于扩展和维护

## 安装要求

- Go 1.16+
- 网络连接以访问GitCode API

## 环境变量配置

项目使用`.env`文件来管理环境变量。您可以复制`.env.example`文件并重命名为`.env`，然后设置以下环境变量：

```
# GitCode API配置
GITCODE_TOKEN=<您的GitCode访问令牌>
GITCODE_API_URL=https://api.gitcode.com/api/v5

# MCP传输配置
# 可选值: stdio
MCP_TRANSPORT=stdio

# API配置
API_TIMEOUT=30
```

## 安装说明

### 方法一：使用安装脚本（推荐）

```bash
# 克隆仓库
git clone https://github.com/gitcode-org-com/gitcode-mcp.git
cd gitcode-mcp

# 运行安装脚本
./install.sh
```

安装脚本会：
1. 编译项目生成可执行文件
2. 创建配置目录 `~/.gitcode_mcp`
3. 复制配置文件到配置目录
4. 提示输入您的GitCode访问令牌
5. 将可执行文件安装到系统路径（需要管理员权限）或用户目录

安装完成后，您可以在任何位置运行 `gitcode_mcp_go` 命令。

### 方法二：使用 Go Install

```bash
# 安装最新版本
go install github.com/gitcode-org-com/gitcode-mcp@latest

# 或者克隆仓库后安装
git clone https://github.com/gitcode-org-com/gitcode-mcp.git
cd gitcode-mcp
go install
```

使用 Go Install 安装后，程序会被安装到 `$GOPATH/bin` 目录下。请确保该目录已添加到您的 PATH 环境变量中。

### 方法三：手动编译安装

```bash
# 克隆仓库
git clone https://github.com/gitcode-org-com/gitcode-mcp.git
cd gitcode-mcp

# 编译项目
go build -o gitcode_mcp_go

# 复制到系统路径（需要管理员权限）
sudo cp gitcode_mcp_go /usr/local/bin/
sudo chmod +x /usr/local/bin/gitcode_mcp_go

# 或者复制到用户目录
mkdir -p ~/bin
cp gitcode_mcp_go ~/bin/
chmod +x ~/bin/gitcode_mcp_go
# 确保 ~/bin 在您的 PATH 中
```

## 配置

首次运行前，请确保设置了您的GitCode访问令牌：

1. 创建配置目录：`mkdir -p ~/.gitcode_mcp`
2. 复制示例配置：`cp .env.example ~/.gitcode_mcp/.env`
3. 编辑配置文件：`nano ~/.gitcode_mcp/.env`
4. 设置您的访问令牌：`GITCODE_TOKEN=您的访问令牌`

## 快速开始

1. 克隆仓库

```bash
git clone https://github.com/gitcode-org-com/gitcode-mcp.git
cd gitcode-mcp
```

2. 安装依赖

```bash
go mod download
```

3. 运行MCP服务器

```bash
# 标准输入输出模式 (STDIO)
go run main.go

# 服务器发送事件模式 (SSE)
MCP_TRANSPORT=sse go run main.go
```

## 在Cursor中使用

GitCode MCP Go Server可以作为Cursor编辑器的MCP服务使用，使您能够在编辑器中直接操作GitCode仓库。

### 配置方法

1. **启动GitCode MCP服务器**

   在STDIO模式下启动MCP服务器：

   ```bash
   go run main.go
   ```

   或者在SSE模式下启动：

   ```bash
   MCP_TRANSPORT=sse go run main.go
   ```

2. **在Cursor中配置MCP**

   有以下几种配置方式：

   **方式一：使用STDIO模式（推荐）**

   在Cursor的设置中，添加以下MCP配置：

   ```json
   {
     "mcpServers": {
       "gitcode": {
         "command": "gitcode_mcp_go",
         "args": [],
         "env": {
           "GITCODE_TOKEN": "<您的GitCode访问令牌>",
           "GITCODE_API_URL": "https://api.gitcode.com/api/v5"
         }
       }
     }
   }
   ```

   您也可以直接使用项目提供的配置文件：

   ```bash
   # 复制Cursor配置文件
   cp ~/.gitcode_mcp/docs/cursor_config.json ~/cursor-config.json
   
   # 编辑配置文件，添加您的GitCode访问令牌
   nano ~/cursor-config.json
   ```

   **方式二：使用SSE模式**

   ```json
   {
     "mcpServers": {
       "gitcode": {
         "url": "http://localhost:8000",
         "env": {
           "GITCODE_TOKEN": "<您的GitCode访问令牌>"
         }
       }
     }
   }
   ```

3. **支持的平台**

   项目docs目录下提供了各平台的配置文件：
   - claude_config.json - Claude平台配置
   - cline_config.json - Cline平台配置
   - cursor_config.json - Cursor平台配置 
   - windsurf_config.json - Windsurf平台配置

### 使用GitCode MCP工具

配置完成后，您可以在Cursor中通过聊天面板使用以下GitCode功能：

#### 仓库操作

- 列出自己的仓库 (list_repositories)
- 获取特定仓库的详情 (get_repository)
- 创建新仓库 (create_repository)

#### 分支管理

- 查看仓库的分支列表 (list_branches)
- 查看分支详情 (get_branch)
- 创建新分支 (create_branch)

#### Issue管理

- 查看仓库的Issue列表 (list_issues)
- 获取Issue详情 (get_issue)
- 创建新Issue (create_issue)

#### Pull Request操作

- 获取PR列表 (list_pull_requests)
- 查看PR详情 (get_pull_request)
- 创建新PR (create_pull_request)

#### 代码搜索

- 搜索代码 (search_code)
- 搜索仓库 (search_repositories)
- 搜索Issue (search_issues)
- 搜索用户 (search_users)

### 实战场景示例

**场景1：处理Issue并创建PR**

1. 查看仓库Issues：
   ```
   可以帮我列出当前仓库的所有Issue吗？
   ```

2. 获取Issue详情：
   ```
   查看Issue #123的详细信息
   ```

3. 创建分支处理Issue：
   ```
   创建一个新分支fix-issue-123来处理这个问题
   ```

4. 实现并提交代码

5. 创建PR：
   ```
   创建一个PR，将fix-issue-123分支合并到main分支，标题为"修复Issue #123"
   ```

**场景2：搜索代码和仓库**

```
搜索GitCode上有关"微服务架构"的仓库
```

```
在当前仓库中搜索所有使用了"Redis缓存"的代码
```

## MCP工具清单

GitCode MCP提供以下工具：

| 工具名称 | 描述 | 参数 |
|---------|------|-----|
| list_repositories | 列出当前用户的仓库 | 无 |
| get_repository | 获取特定仓库的详细信息 | owner, repo |
| create_repository | 创建新仓库 | name, description?, private? |
| list_branches | 列出仓库的分支 | owner, repo |
| get_branch | 获取特定分支的详细信息 | owner, repo, branch |
| create_branch | 创建新分支 | owner, repo, branch, ref |
| list_issues | 列出仓库的Issues | owner, repo |
| get_issue | 获取特定Issue的详细信息 | owner, repo, issue_number |
| create_issue | 创建新Issue | owner, repo, title, body? |
| list_pull_requests | 列出仓库的Pull Requests | owner, repo |
| get_pull_request | 获取特定Pull Request的详细信息 | owner, repo, pull_number |
| create_pull_request | 创建新Pull Request | owner, repo, title, head, base, body? |
| search_code | 搜索代码 | query |
| search_repositories | 搜索仓库 | query |
| search_issues | 搜索Issues | query |
| search_users | 搜索用户 | query |

## 许可证

该项目采用MIT许可证。详情请参阅LICENSE文件。

## 平台配置文件

在docs目录下，提供了四个主要AI平台的配置文件示例：

- **Claude平台**: docs/claude_config.json
- **Cline平台**: docs/cline_config.json  
- **Cursor平台**: docs/cursor_config.json
- **Windsurf平台**: docs/windsurf_config.json

这些配置文件已经包含了基本设置，您只需要替换`<您的GitCode访问令牌>`为您自己的访问令牌即可使用。

### 使用配置文件

```bash
# 复制对应平台的配置文件到适当位置
cp docs/cursor_config.json ~/cursor-config.json

# 编辑配置文件，添加您的GitCode访问令牌
nano ~/cursor-config.json
```

然后根据各平台的配置方法，将配置文件路径添加到相应的设置中。