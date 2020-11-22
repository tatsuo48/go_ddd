package command

type CircleCreateCommand struct {
	userID     string
	circleName string
}

func NewCircleCreateCommand(id string, name string) CircleCreateCommand {
	return CircleCreateCommand{
		userID:     id,
		circleName: name,
	}
}

type CircleJoinCommand struct {
	userID   string
	circleID string
}

func NewCircleJoinCommand(userID string, circleID string) CircleJoinCommand {
	return CircleJoinCommand{
		userID:   userID,
		circleID: circleID,
	}
}
