package bst

type ComplexityResults struct {
	Porownania          int
	OdczytyPodstawienia int
	Wysokosc            int
}

func (c *ComplexityResults) Add(other ComplexityResults) {
	c.Porownania += other.Porownania
	c.OdczytyPodstawienia += other.OdczytyPodstawienia
	c.Wysokosc += other.Wysokosc
}

func (c *ComplexityResults) Divide(n int) {
	c.Porownania /= n
	c.OdczytyPodstawienia /= n
	c.Wysokosc /= n
}
