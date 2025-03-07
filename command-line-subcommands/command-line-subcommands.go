package main



/*Some command-line tools, like the go tool or git have many subcommands, each with its own set of flags. 
For example, go build and go get are two different subcommands of the go tool. 
The flag package lets us easily define simple subcommands that have their own flags.*/
import (
    "flag"
    "fmt"
    "os"
)

func main() {
	/*We declare a subcommand using the NewFlagSet function, and proceed to define new flags specific for this subcommand.*/
    fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
    fooEnable := fooCmd.Bool("enable", false, "enable")
    fooName := fooCmd.String("name", "", "name")

	/*For a different subcommand we can define different supported flags.*/
    barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
    barLevel := barCmd.Int("level", 0, "level")

	/*The subcommand is expected as the first argument to the program.*/
    if len(os.Args) < 2 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

	/*Check which subcommand is invoked.*/
    switch os.Args[1] {

		/*For every subcommand, we parse its own flags and have access to trailing positional arguments.*/
    case "foo":
        fooCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'foo'")
        fmt.Println("  enable:", *fooEnable)
        fmt.Println("  name:", *fooName)
        fmt.Println("  tail:", fooCmd.Args())
    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println("  level:", *barLevel)
        fmt.Println("  tail:", barCmd.Args())
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }
}

/*$ go build command-line-subcommands.go 

First invoke the foo subcommand.
	

$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

Now try bar.
	

$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

But bar won’t accept foo’s flags.
	

$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
        level*/