package gof

func NewDefaultRouterFactory() *RouterFactory {
	factory := NewRouterFactory()
	factory.Use(Logger())
	factory.Use(Recovery())
	return factory
}
