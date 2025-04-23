package domain

type Dialogue struct {
	ID        uint
	UserID    string
	SessionID string
	ModelType string
	Request   string
	Response  string
}

type Dialogues struct {
	SystemRole, UserRole string
	DialogueList         []*Dialogue
}

func (d *Dialogues) ToDialogueModel() []map[string]string {
	dm := make([]map[string]string, 0, len(d.DialogueList)*2)
	for i := 0; i < len(d.DialogueList); i++ {
		dm = append(dm, map[string]string{
			"role":    d.UserRole,
			"content": d.DialogueList[i].Request,
		})

		if d.DialogueList[i].Response != "" {
			dm = append(dm, map[string]string{
				"role":    d.SystemRole,
				"content": d.DialogueList[i].Response,
			})
		}
	}

	return dm
}
