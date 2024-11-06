package components

type Combat interface {
	Health() int
	AttackPower() int
	Damage(amount int)
}

type BasicCombat struct {
	health      int
	attackPower int
}

func (b *BasicCombat) AttackPower() int {
	return b.attackPower
}

func (b *BasicCombat) Health() int {
	return b.health
}

func (b *BasicCombat) Damage(amount int) {
	b.health -= amount
}

var _ Combat = (*BasicCombat)(nil)
