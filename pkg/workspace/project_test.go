package workspace

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestProjectRuntimeInfoRoundtrip(t *testing.T) {
	ri := NewProjectRuntimeInfo("nodejs", nil)
	byts, err := yaml.Marshal(ri)
	assert.NoError(t, err)

	var riRountrip ProjectRuntimeInfo
	err = yaml.Unmarshal(byts, &riRountrip)
	assert.NoError(t, err)
	assert.Equal(t, "nodejs", riRountrip.Name())
	assert.Equal(t, nil, riRountrip.Options())
}
