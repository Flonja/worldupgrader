package worldupgrader

import (
	"github.com/df-mc/worldupgrader/itemupgrader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemRemappedMetaDowngrade(t *testing.T) {
	item := itemupgrader.ItemMeta{
		Name: "minecraft:green_concrete",
		Meta: 0,
	}
	upgraded := itemupgrader.Downgrade(item, 111)
	assert.Equal(t, "minecraft:concrete", upgraded.Name)
	assert.Equal(t, int16(13), upgraded.Meta)
}

func TestItemRemappedMetaUnableToDowngrade(t *testing.T) {
	item := itemupgrader.ItemMeta{
		Name: "minecraft:green_concrete",
		Meta: 0,
	}
	upgraded := itemupgrader.Downgrade(item, 121)
	assert.NotEqual(t, "minecraft:concrete", upgraded.Name)
	assert.NotEqual(t, int16(13), upgraded.Meta)
}
