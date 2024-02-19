package payment_processors

type Processors struct {
	registry map[string]ProcessorProvider
}

func (p *Processors) GetRegistry(key string) ProcessorProvider {
	return p.registry[key]
}

type ProcessorProvider interface {
}
