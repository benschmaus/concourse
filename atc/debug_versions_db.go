package atc

type DebugVersionsDB struct {
	ResourceVersions []DebugResourceVersion
	BuildOutputs     []DebugBuildOutput
	BuildInputs      []DebugBuildInput
	JobIDs           map[string]int
	ResourceIDs      map[string]int
}

type DebugResourceVersion struct {
	VersionID  int
	ResourceID int
	CheckOrder int
	Disabled   bool
}

type DebugBuildOutput struct {
	DebugResourceVersion
	BuildID int
	JobID   int
}

type DebugBuildInput struct {
	DebugResourceVersion
	BuildID         int
	JobID           int
	InputName       string
	FirstOccurrence bool
}
