package types

type RetryStrategy string

const (
	SimpleRetry      RetryStrategy = "simple"      // SimpleRetry is a retry strategy with fixed delay
	ExponentialRetry RetryStrategy = "exponential" // ExponentialRetry is a retry strategy with exponential backoff
	JitterRetry      RetryStrategy = "jitter"      // JitterRetry is a retry strategy with random delay
)

func (r RetryStrategy) String() string {
	return string(r)
}

func (r RetryStrategy) IsValid() bool {
	switch r {
	case SimpleRetry, ExponentialRetry, JitterRetry:
		return true
	}
	return false
}
