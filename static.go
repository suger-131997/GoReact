package static

import "embed"

//go:embed assets/*
var Assets embed.FS

//go:embed .output/page
var PageFS embed.FS
