package rabbitmq

func NewOptions() *Options {
	return nil
}

type Options struct {
}

func (o *Options) Validate() []error {
	return nil
}

// type Option func(l *Logger)
