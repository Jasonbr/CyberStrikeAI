package app

import (
	"secai/internal/config"
	"secai/internal/mcp"
	"secai/internal/vision"

	"go.uber.org/zap"
)

func registerVisionTools(mcpServer *mcp.Server, cfg *config.Config, logger *zap.Logger) {
	vision.RegisterAnalyzeImageTool(mcpServer, cfg, logger)
}
