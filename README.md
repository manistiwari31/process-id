Process Manager

Process Manager is a command-line tool for managing processes on Linux based on their names and optionally killing them.
Features

    List all running processes.
    Filter processes by name.
    Optionally kill processes matching a specified name.

Installation

    Make sure you have Go installed on your system.

    Clone the repository:

    $ git clone https://github.com/manistiwari31/process-id.git
    $ cd process-id


Build the program:

    $ go build -o proc


This will generate a binary named proc.

Optionally, move the binary to your bin directory to run it as a Linux command:


    $ sudo mv proc /usr/local/bin

Usage

bash

# Show all running processes
$ proc show

# Show processes matching a specific name
$ proc process_name

# Kill processes matching a specific name
$ proc process_name kill

Replace process_name with the name of the process you want to filter. Use kill as an action to terminate matching processes.
Example

Suppose you want to find and kill all processes named "myapp":

    $ proc myapp kill
Killing 1234
Process Killed

Contributing

Feel free to contribute by following these guidelines...
