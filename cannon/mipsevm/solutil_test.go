package mipsevm

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSourcemap(t *testing.T) {
	contract, err := LoadContract("MIPS")
	require.NoError(t, err)
	srcMap, err := contract.SourceMap([]string{"../../packages/contracts-bedrock/contracts/cannon/MIPS.sol"})
	require.NoError(t, err)
	for i := 0; i < len(contract.DeployedBytecode.Object); i++ {
		info := srcMap.FormattedInfo(uint64(i))
		if !strings.HasPrefix(info, "generated:") && !strings.HasPrefix(info, "../../packages/contracts-bedrock/contracts/cannon/MIPS.sol") {
			t.Fatalf("unexpected info: %q", info)
		}
	}
}
