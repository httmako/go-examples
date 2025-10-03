# go-subcommands
An example of how to create a Golang application with subcommands like docker or kubectl.

This project is named "tharja" internally, which is also the main module name. You can change it ofcourse.  
The 2 provided modules are just examples, one serving the current directory as a webserver and the other calculating the sha256 hash of the input or file provided.

## Structure

When a module is included, the "init" function is run (and not the main function).  
We have a handler module which has a map of all subcommands.  
First we include the handler into the main function and then include the modules to "_".  

This way each module is registering itself into the handler by calling "handler.Register" in the init function.  
The main function only has to call the "handler.ExecuteCommand()" function, which parses the os.Args and executes the corresponding subcommand.

You can register multiple subcommands in a single "init" function of a module.  
You can not use the main package as the handler because it would cause a circular dependency error.  


