package version

import "fmt"

// Version represents the current version of Delve.
type Version struct {
	Major    string
	Minor    string
	Patch    string
	Metadata string
	Build    string
}

var (
	// DelveVersion is the current version of Delve.
	DelveVersion = Version{
		Major: "1", Minor: "5", Patch: "1", Metadata: "",
		Build: "$Id: bca418ea7ae2a4dcda985e623625da727d4525b5 $",
	}
)

func (v Version) String() string {
	ver := fmt.Sprintf("Version: %s.%s.%s", v.Major, v.Minor, v.Patch)
	if v.Metadata != "" {
		ver += "-" + v.Metadata
	}
	return fmt.Sprintf("%s\nBuild: %s", ver, v.Build)
}
