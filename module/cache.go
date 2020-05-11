package module

import "github.com/superloach/defunct/types"

type Cache struct {
	Modules types.Thing
}

func (c *Cache) Call(under types.Funct, args []types.Funct) types.Funct {
	return types.Zilch
}

func (c *Cache) GetProp(under types.Funct, name string) types.Funct {
	return types.Zilch
}

func (c *Cache) SetProp(under types.Funct, name string, val types.Funct) types.Funct {
	return types.Zilch
}

func (c *Cache) String() string {
	return "{Cache}"
}
