# GitCode MCP 平台配置说明

本文档介绍如何在不同AI平台上配置和使用GitCode MCP服务。

## 配置文件概述

项目提供了四个主要AI平台的配置文件：

- `claude_config.json` - Claude平台配置
- `cline_config.json` - Cline平台配置
- `cursor_config.json` - Cursor平台配置
- `windsurf_config.json` - Windsurf平台配置

这些配置文件使用相同的基本结构，都采用STDIO模式与GitCode MCP服务器通信。

## 使用说明

各平台的配置文件在docs目录下，使用前需要替换<您的GitCode访问令牌>为您自己的令牌。

安装脚本会自动将配置文件复制到~/.gitcode_mcp/docs/目录下。

