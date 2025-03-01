package cfg

/*
Represents a generic configuration object that has many QoL features like default values,
marshaling/unmarshaling to/from byte streams, and more.
*/
type IConfig interface {
	/*
		Sets the default values for this configuration object. Example function body
		(assuming implementor is called Foo):

			func (c *Foo) Defaults() (cfg.IConfig, error) {
				return cfg.DefaultsHelper(c, nil)
			}
	*/
	Defaults() (IConfig, error)
	
	//SaveDefault() error
	//Save() error
	SaveAs(path string) error
	
	//Load() (IConfig, error)
	//LoadAs(path string) (IConfig, error)
	
	//Init() (IConfig, error)
	//InitAs(path string) (IConfig, error)

	//defaultPathName() string //Specifies the default location to drop the default configuration.
}

//func SaveDefault() {}		// Saves the default config object to the default path.
//func Save() {}			// Saves a config object to the default path.
//func SaveAs() {}			// Saves a config object to a given path.
//func Load() {}			// Saves a config object to the default path.
//func LoadAs() {}			// Loads a config object from a given path.
//func Init() {}			// Loads a config object from the default path, saving the default version if it doesn't exist.
//func InitAs() {}			// Loads a config object from a given path, saving the default version if it doesn't exist.

//TODO: try to use a custom struct tag `kname` that indicates the name of the key
