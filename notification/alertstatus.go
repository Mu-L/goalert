package notification

type AlertStatus struct {
	Dest       Dest
	CallbackID string
	LogEntry   string
	AlertID    int
}

var _ Message = &AlertStatus{}

func (s AlertStatus) Type() MessageType { return MessageTypeAlertStatus }
func (s AlertStatus) ID() string        { return s.CallbackID }
func (s AlertStatus) Destination() Dest { return s.Dest }
