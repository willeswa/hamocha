package cmd

var fileTypes = map[string][]string {
	"Image": {".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg"},
	"Audio": {".mp3", ".wav", ".flac", ".aac", ".ogg"},
	"Video": {".mp4", ".mkv", ".avi", ".mov", ".flv", ".wmv"},
	"Document": {".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt"},
	"Archive": {".zip", ".tar", ".gz", ".7z", ".rar"},
}

type File struct {
	Name string
	Size int
	IsDir bool
	Ext string
}