package stgmachine

type InitialState struct {
	WorkflowID     int    `json:"workflow_id"`
	ResourceTypeID int    `json:"resource_type_id"`
	ResourceID     string `json:"resource_id"`
	StageID        int    `json:"stage_id"`
}
