package envsource

type Source interface {
	LookupEnv(k string) (string, bool)
}

type Composed struct {
	sources []Source
}

func (c *Composed) LookupEnv(k string) (string, bool) {
	for _, source := range c.sources {
		if v, ok := source.LookupEnv(k); ok {
			return v, true
		}
	}
	return "", false
}

func ComposeSources(source Source, sources ...Source) *Composed {
	sources = prepend(sources, source)
	return &Composed{sources: sources}
}

func prepend[T interface{}](arr []T, v T) []T {
	if len(arr) == 0 {
		arr = append(arr, v)
	} else {
		arr = append([]T{v}, arr...)
	}
	return arr
}
