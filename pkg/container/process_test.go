package container

import (
	"github.com/initialed85/km1s/pkg/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	var p *Process
	var err error
	var logs string

	t.Run("Integration", func(t *testing.T) {
		defer func() {
			if p == nil {
				return
			}

			p.Close()
		}()

		p = NewProcess(
			"while true; do echo \"hi $NAME @ $(date +%s%3N)\"; sleep 0.1; done",
			128,
			map[string]string{
				"NAME": "Joe Bloggs",
			},
		)
		logs = p.Logs()
		assert.Equal(t, 0, len(logs))

		err = p.Open()
		test.FailOnError(t, err)

		time.Sleep(time.Millisecond * 100)
		logs = p.Logs()
		assert.Greater(t, len(logs), 0)
		assert.Less(t, len(logs), 128)

		time.Sleep(time.Millisecond * 500)
		logs1 := p.Logs()
		assert.Equal(t, 128, len(logs1))

		time.Sleep(time.Millisecond * 250)
		logs2 := p.Logs()
		assert.Equal(t, 128, len(logs2))

		assert.NotEqual(t, logs1, logs2)

		p.Close()
	})
}
