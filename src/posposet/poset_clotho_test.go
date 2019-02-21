package posposet

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosetSimpleClotho(t *testing.T) {
	testSpecialNamedCC(t, `
a01     b01     c01
║       ║       ║
a11 ─ ─ ╬ ─ ─ ─ ╣       d01
║       ║       ║       ║
║       ╠ ─ ─ ─ c11 ─ ─ ╣
║       ║       ║       ║       e01
╠ ─ ─ ─ B12++ ─ ╣       ║       ║
║       ║       ║       ║       ║
║       ║       ╠ ─ ─ ─ D12++ ─ ╣
║       ║       ║       ║       ║
A22++ ─ ╫ ─ ─ ─ ╬ ─ ─ ─ ╣       ║
║       ║       ║       ║       ║
╠ ─ ─ ─ ╫ ─ ─ ─ ╫ ─ ─ ─ ╬ ─ ─ ─ E12++
║       ║       ║       ║       ║
╠ ─ ─ ─ ╫ ─ ─ ─ C22++ ─ ╣       ║
║       ║       ║       ║       ║
╠ ─ ─ ─ B23+  ─ ╣       ║       ║
║       ║       ║       ║       ║
║       ║       ╠ ─ ─ ─ D23+  ─ ╣
║       ║       ║       ║       ║
║       ╠ ─ ─ ─ ╫ ─ ─ ─ ╬ ─ ─ ─ E23+
║       ║       ║       ║       ║
A33+  ─ ╬ ─ ─ ─ ╣       ║       ║
║       ║       ║       ║       ║
║       ╠ ─ ─ ─ C33+    ║       ║
║       ║       ║       ║       ║
╠ ─ ─ ─ b33 ─ ─ ╣       ║       ║
║       ║       ║       ║       ║
a43 ─ ─ ╬ ─ ─ ─ ╣       ║       ║
║       ║       ║       ║       ║
║       ╠ ─ ─ ─ C44 ─ ─ ╣       ║
║       ║       ║       ║       ║
╠ ─ ─ ─ B44 ─ ─ ╣       ║       ║
║       ║       ║       ║       ║
║       ║       ╠ ─ ─ ─ D34 ─ ─ ╣
║       ║       ║       ║       ║
A54 ─ ─ ╫ ─ ─ ─ ╬ ─ ─ ─ ╣       ║
║       ║       ║       ║       ║
╠ ─ ─ ─ ╫ ─ ─ ─ c54 ─ ─ ╣       ║
║       ║       ║       ║       ║
║       ║       ╠ ─ ─ ─ ╬ ─ ─ ─ E34
║       ║       ║       ║       ║
`)
}

/*
 * Utils:
 */

// testSpecialNamedCC is a general test of root selection.
// Node name means:
// - 1st letter uppercase - node should be root;
// - 2nd number - index by node;
// - 3rd number - frame where node should be in;
// - last "+" - single if ClothoCandidate, double if Clotho;
func testSpecialNamedCC(t *testing.T, asciiScheme string) {
	assert := assert.New(t)
	// init
	nodes, _, names := ParseEvents(asciiScheme)
	p := FakePoset(nodes)
	// process events
	for _, event := range names {
		p.PushEventSync(*event)
	}
	// check each
	for name, event := range names {
		// check roots
		mustBeRoot := (name == strings.ToUpper(name))
		frame, isRoot := p.FrameOfEvent(event.Hash())
		if !assert.Equal(mustBeRoot, isRoot, name+" is root") {
			break
		}
		// check frames
		mustBeFrame, err := strconv.ParseUint(name[2:3], 10, 64)
		if !assert.NoError(err, "name the nodes properly: <UpperCaseForRoot><Index><FrameN>") {
			return
		}
		if !assert.Equal(mustBeFrame, frame.Index, "frame of "+name) {
			break
		}
		// check Clotho
		mustBeCC := len(name) > 3 && name[3:4] == "+"
		isCC := frame.ClothoCandidates[event.Creator].Contains(event.Hash())
		if !assert.Equal(mustBeCC, isCC, name+" is CC") {
			break
		}
		continue // TODO: temporary
		mustBeClotho := len(name) > 4 && name[3:5] == "++"
		isClotho := frame.ClothoList[event.Creator].Contains(event.Hash())
		if !assert.Equal(mustBeClotho, isClotho, name+" is Clotho") {
			break
		}

	}
}
