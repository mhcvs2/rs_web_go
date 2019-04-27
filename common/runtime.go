package common

func HandleCrash(additionalHandlers ...func(interface{})) {
	if r := recover(); r != nil {
		for _, fn := range additionalHandlers {
			fn(r)
		}
	}
}
