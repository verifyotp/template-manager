package s3

func WithEnv(env string) Option {
	return func(s *S3) {
		s.env = env
	}

}

func WithRegion(region string) Option {
	return func(s *S3) {
		s.region = region
	}
}

func WithFolder(folder string) Option {
	return func(s *S3) {
		s.folder = folder
	}
}
