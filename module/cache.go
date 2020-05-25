package module

import "github.com/superloach/defunct/types"

type Cache struct {
	Module *Module
	Cache  types.Thing
}

func (c *Cache) Load(name string) types.Funct {
	panic("cache load stub")
	return types.Zilch
}

func (c *Cache) Call(under types.Funct, args []types.Funct) types.Funct {
	panic("cache call stub")
	return types.Zilch
}

func (c *Cache) GetProp(under types.Funct, name string) types.Funct {
	if len(name) == 0 {
		return types.Zilch
	}

	if name[0] == '.' {
		panic("relative import stub")
	}

	panic("import stub")

	return types.Zilch
}

func (c *Cache) SetProp(under types.Funct, name string, val types.Funct) types.Funct {
	panic("cache setprop stub")
	return types.Zilch
}

func (c *Cache) String() string {
	return "{Cache}"
}
