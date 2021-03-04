package persondependencyinjection

// PersonDependencyInjection injects handlers on its buses
type PersonDependencyInjection interface {
	RegisterCommandHandlers()
	RegisterQueryHandlers()
}
