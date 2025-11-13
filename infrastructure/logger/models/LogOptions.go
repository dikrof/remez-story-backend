package logModels

type Option func(opts *Options)

type Options struct {
	withStackTrace bool
	component      string
	fields         []*LogField
}

func (o *Options) WithStackTrace() bool {
	return o.withStackTrace
}

func (o *Options) GetComponent() string {
	return o.component
}

func (o *Options) GetFields() []*LogField {
	return o.fields
}

func WithComponent(component string) Option {
	return func(opts *Options) {
		opts.component = component
	}
}

func WithStackTrace() Option {
	return func(opts *Options) {
		opts.withStackTrace = true
	}
}

func WithStringField(key string, value string) Option {
	return func(opts *Options) {
		opts.fields = append(opts.fields, &LogField{Key: key, String: value})
	}
}
