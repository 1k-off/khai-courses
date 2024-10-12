package entities

type Schedule []struct {
	Day     string   `json:"day"`
	Lessons []Lesson `json:"lessons"`
}

type Lesson struct {
	Name    string `json:"name"`
	Teacher string `json:"teacher"`
	Room    string `json:"room"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

const (
	DayMonday    = "monday"
	DayTuesday   = "tuesday"
	DayWednesday = "wednesday"
	DayThursday  = "thursday"
	DayFriday    = "friday"
)

func ValidateDay(day string) bool {
	switch day {
	case DayMonday, DayTuesday, DayWednesday, DayThursday, DayFriday:
		return true
	default:
		return false
	}
}
