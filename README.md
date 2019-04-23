# Data parser
Use to parse data. As well as, it reads file and parse data. Commonly use to read config file, messages etc in application.

## JSON

### How to use?
- Can load data in bytes using .Load(data)
- Can load data from file using .LoadFile(path) 
````
var config parser.JSON

_ = config.LoadFile("examples/config.json")
````

#### Available methods 
- Get
- GetString
- GetInt
- GetFloat
- GetBool
- GetTime
- GetDuration

Can get data from inner object/array too. Use dot(.) to join keys.

#### Example 
````
// JSON
{
  "test": 1,
  "outer": [{
    "inner": "Hello World",
    "bool": false
  }]
}

// Get
config.GetString("outer.0.inner")
````