
# ArgParser

Argparser function for golang to treat flags.

The objective is to simplify the creation of an object containing the parsed arguments.

## Normal Usage

You will need to create a struct with the flags you want to parse in order to keep the data on a object. Then you will just need to set the flags and update the object with the data.

``` Golang
...
// Create a struct to store the data (this is not required as you can use the alternative method for parsing)
type args struct {
	verbose int
}

// Iniciate the ArgParser
arguments := ArgParser.Init(ArgParser{})

// setup the flags you want to parse
arguments.SetFlags("-v", // Key
    "--verbose", // Name
    1, // Default value
    true, // Required (bool)
    "Verbosity of the prints, Value ranges from 1 to 3") // Description to appear on the help menu

// Parse the console arguments
arguments.Parse()

// Set arguments to the 
aparsed_arguments := args{
    verbose: ap.GetFieldValue("verbose").(int), // get field by name, will need to specify the var type
}

// Obtain the value
fmt.Printf("%v", aparsed_arguments.verbose)

```

## Simplified Usage

In this method we lose the object with the assignment and keep only the struct object

``` Golang
...

// Iniciate the ArgParser
arguments := ArgParser.Init(ArgParser{})

// setup the flags you want to parse
arguments.SetFlags("-v", // Key
    "--verbose", // Name
    1, // Default value
    true, // Required (bool)
    "Verbosity of the prints, Value ranges from 1 to 3") // Description to appear on the help menu

// Parse the console arguments
arguments.Parse()

// Obtain the value method 1
x := arguments.GetFieldValue("verbose").(int)// get field by name, will need to specify the var type
fmt.Printf("%v", x)

// Obtain the value method 2
for _, arg := range arguments.flags {
    fmt.Printf("Key: %v, Value: %v\n", arg.key, arg.value)
}

```

## V1.0.0
- Inicial Commit