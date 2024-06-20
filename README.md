Process Manager

A command-line tool for managing processes on Linux based on their names and optionally killing them.
Features

    List all running processes.
    Filter processes by name.
    Optionally kill processes matching a specified name.

Installation

    Make sure you have Go installed on your system.
    Clone the repository:

    bash

$ git clone https://github.com/manistiwari31/process-id.git
$ cd your-repo

Build the program:

bash

$ go build

Run the program:

bash

    $ ./process-manager [process_name] [action]

    Replace [process_name] with the name of the process you want to filter or "show" to display all processes. Use [action] as "kill" to terminate matching processes.

Usage

bash

# Show all running processes
$ ./process-manager show

# Show processes matching a specific name
$ ./process-manager process_name

# Kill processes matching a specific name
$ ./process-manager process_name kill

Example

Suppose you want to find and kill all processes named "myapp":

bash

$ ./process-manager myapp kill
Killing 1234
Process Killed

Contributing

Feel free to contribute by following these guidelines...
