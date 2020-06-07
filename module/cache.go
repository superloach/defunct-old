package module

import "github.com/superloach/defunct/types"

type Cache struct {
	Module *Module
	Cache  types.Thing
}

func (c *Cache) Load(name string) types.Funct {
	println("TODO: cache load")
	return types.Zilch
}

func (c *Cache) Call(under types.Funct, args []types.Funct) types.Funct {
	println("TODO: cache call")
	return types.Zilch
}

func (c *Cache) GetProp(under types.Funct, name string) types.Funct {
	if len(name) == 0 {
		return types.Zilch
	}

	if name[0] == '.' {
		println("TODO: relative import")
	}

	println("TODO: import")

	return types.Zilch
}

func (c *Cache) SetProp(under types.Funct, name string, val types.Funct) types.Funct {
	println("TODO: cache setprop")
	return types.Zilch
}

func (c *Cache) String() string {
	return "{Cache}"
}
