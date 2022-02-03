package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var (
	version      string
	gitTag       string
	gitCommit    = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState = "not a git tree"       // state of git tree, either "clean" or "dirty"
	buildDate    = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// Info contains versioning information.
type Info struct {
	Version      string `json:"version"`
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// string returns info as a human-friendly version string.
func (info Info) string() string {
	return info.GitTag
}

// get return version info of current build
func get() Info {
	return Info{
		Version:      version,
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// GetString return version string
func GetString() (v string, err error) {
	vStruct := get()
	marshalled, err := json.MarshalIndent(&vStruct, "", "  ")
	if err != nil {
		return "", err
	}
	return string(marshalled), nil
}
