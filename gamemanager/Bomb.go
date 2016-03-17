package gamemanager

type Bomb struct {
	owner *Player
	field *Field
}

func NewBomb(p *Player, f *Field) *Bomb {
	return &Bomb{
		owner: p,
		field: f,
	}
}

func (bomb *Bomb) explode(gameMap *GameMap) {
	fields := gameMap.GetNOSWFieldsOfField(bomb.field)

	// Ausgangsfeld hinzufügen
	fields = append(fields, bomb.field)

	for i := range fields {
		// Wände werden durch Explosionsstrahl zerstört
		if fields[i].wall != nil {
			if fields[i].wall.isDestructible {
				fields[i].wall = nil
			}
		}

		// Specials werden durch Explosionsstrahl zerstört
		fields[i].wall = nil

		// Spieler werden durch Explosionsstrahl gelähmt
		for _, v := range fields[i].players {
			v.isParalyzed = true
		}
	}
}
