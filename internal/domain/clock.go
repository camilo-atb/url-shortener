package domain

import "time"

type Clock interface { // esto es un Output Port
	Now() time.Time
}
