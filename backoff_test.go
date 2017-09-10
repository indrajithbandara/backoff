package backoff

import (
	"testing"
	"time"

	"github.com/tj/assert"
)

func Test1(t *testing.T) {

	b := &Backoff{
		Min:    100 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
	}

	assert.Equal(t, b.Duration(), 100*time.Millisecond)
	assert.Equal(t, b.Duration(), 200*time.Millisecond)
	assert.Equal(t, b.Duration(), 400*time.Millisecond)
	b.Reset()
	assert.Equal(t, b.Duration(), 100*time.Millisecond)
}

func TestForAttempt(t *testing.T) {

	b := &Backoff{
		Min:    100 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
	}

	assert.Equal(t, b.ForAttempt(0), 100*time.Millisecond)
	assert.Equal(t, b.ForAttempt(1), 200*time.Millisecond)
	assert.Equal(t, b.ForAttempt(2), 400*time.Millisecond)
	b.Reset()
	assert.Equal(t, b.ForAttempt(0), 100*time.Millisecond)
}

func Test2(t *testing.T) {

	b := &Backoff{
		Min:    100 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 1.5,
	}

	assert.Equal(t, b.Duration(), 100*time.Millisecond)
	assert.Equal(t, b.Duration(), 150*time.Millisecond)
	assert.Equal(t, b.Duration(), 225*time.Millisecond)
	b.Reset()
	assert.Equal(t, b.Duration(), 100*time.Millisecond)
}

func Test3(t *testing.T) {

	b := &Backoff{
		Min:    100 * time.Nanosecond,
		Max:    10 * time.Second,
		Factor: 1.75,
	}

	assert.Equal(t, b.Duration(), 100*time.Nanosecond)
	assert.Equal(t, b.Duration(), 175*time.Nanosecond)
	assert.Equal(t, b.Duration(), 306*time.Nanosecond)
	b.Reset()
	assert.Equal(t, b.Duration(), 100*time.Nanosecond)
}

func Test4(t *testing.T) {
	b := &Backoff{
		Min:    500 * time.Second,
		Max:    100 * time.Second,
		Factor: 1,
	}

	assert.Equal(t, b.Duration(), b.Max)
}

func TestGetAttempt(t *testing.T) {
	b := &Backoff{
		Min:    100 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
	}
	assert.Equal(t, b.Attempt(), 0)
	assert.Equal(t, b.Duration(), 100*time.Millisecond)
	assert.Equal(t, b.Attempt(), 1)
	assert.Equal(t, b.Duration(), 200*time.Millisecond)
	assert.Equal(t, b.Attempt(), 2)
	assert.Equal(t, b.Duration(), 400*time.Millisecond)
	assert.Equal(t, b.Attempt(), 3)
	b.Reset()
	assert.Equal(t, b.Attempt(), 0)
	assert.Equal(t, b.Duration(), 100*time.Millisecond)
	assert.Equal(t, b.Attempt(), 1)
}

func TestJitter(t *testing.T) {
	b := &Backoff{
		Min:    100 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
		Jitter: true,
	}

	assert.Equal(t, b.Duration(), 100*time.Millisecond)
	between(t, b.Duration(), 100*time.Millisecond, 200*time.Millisecond)
	between(t, b.Duration(), 100*time.Millisecond, 400*time.Millisecond)
	b.Reset()
	assert.Equal(t, b.Duration(), 100*time.Millisecond)
}

func between(t *testing.T, actual, low, high time.Duration) {
	if actual < low {
		t.Fatalf("Got %s, Expecting >= %s", actual, low)
	}
	if actual > high {
		t.Fatalf("Got %s, Expecting <= %s", actual, high)
	}
}
