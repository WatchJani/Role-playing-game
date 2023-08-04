package enemy

import s "github.com/WatchJani/Role-playing-game/model/skeleton"

type Enemy struct {
	s.Skeleton
	Movement float64
	Class    string
}
