package lib

type Download struct {
	URL string
	Dest string
	Progress chan uint64
	ProgInt uint64
	TotalSize uint64
}

type Mod struct {
	ID uint `json:"id"`
	ModID uint `json:"mod_id"`
	ModKey string `json:"mod_key"`
	MinecraftVersion []string `json:"minecraft_version"`
	JavaVersion []string `json:"java_version"`
	Filename string `json:"file_name"`
	FileSize string `json:"file_size"`
	FileMD5 string `json:"file_md5"`
	ReleaseType string `json:"release_type"`
	DownloadURL string `json:"download_url"`
	Uploaded string `json:"uploaded"`
	DownloadCount uint32 `json:"download_count"`
	ModDependencies []string `json:"mod_dependencies"`
}

type CurseResp struct {
	Status string `json:"status"`
	ExecutionTime uint `json:"execution_time"`
	Result []Mod `json:"result"`
}

type JSONPack struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Version string `json:"version"`
	MCVersion string `json:"mcversion"`
	Mods []JSONMod `json:"mods"`
}

type JSONMod struct {
	Title string `json:"title"`
	CurseName string `json:"cursename"`
	Release uint `json:"release"`
}