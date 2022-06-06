package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMD5(t *testing.T) {
	assert.EqualValues(t, "992715f4af1861e8a15979a40713e4c9", MD5("/A?D(G-K"))
	assert.EqualValues(t, "56201a0bbb4b31bf24a18c6615dbae98", MD5("u7x!A%D*"))
	assert.EqualValues(t, "f55d8bf9ce4dac0be06a480f256792a6", MD5("Zq4t7w!z"))
	assert.EqualValues(t, "7bae927cf98fa541e4d1bdd52f6c2d1c", MD5("hVmYq3t6"))
	assert.EqualValues(t, "82a78f7729b16bb81f04135e9726aafa", MD5("PdSgVkYp"))
	assert.EqualValues(t, "a378c082ff07c258d79b292b3e9f6536", MD5("-JaNdRgU"))
	assert.EqualValues(t, "d1f51dc4f0c2cc638928642f84681c60", MD5("C&F)J@Nc"))
	assert.EqualValues(t, "6126d5ea5c6655d9d2a0397ad26f3390", MD5("9z$C&E)H"))
	assert.EqualValues(t, "35e4c806f427a43a9bac08cf3224a6b7", MD5("s6v9y$B&"))
	assert.EqualValues(t, "5d6a755967a2780ffedb5215a6df1b9f", MD5("Xp2s5v8y"))
}
