package workers

type TFActionType string

const (
	TfActionPlanType   = TFActionType("plan")
	TfActionInitType   = TFActionType("init")
	TfActionDeployType = TFActionType("deploy")
)

type TfAction interface {
	Type() string
}
