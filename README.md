# Now : A simple time marking tool
Now enables you to mark different event of the day with a timestamp and tags.
Ultimately this allow you to figure the answer for the question "What did I do today?"

This tool is in its nacent stage.With continuous developement and active contributors I plan to implement folllowing features,

1. Summary command
2. Command line gui with more 
3. A todo list feature(most via plan command)
4. A complete nerdy!! commmand line task planning and tracking solution

## Installation
Follwing commands after clone is repository will install the now tool.

    go mod tidy
    go install now

## Usage
Now tool uses intuitive commands that enables ease of use.

You can always use help for usage and function of different commands.
    
    now help [command]

To add an entry that indicates start of a task
    
    now doing taskTitle --tags tag1,tag2

To add an entry that indicates completion of a task

    now done taskTitle --tags tag1,tag2

To add an break entry 

    now break
Break command also supports tags.

To add an continue entry that indicate completion of a break

    now continue
Continue command also suports tags.

To display entries

    now show #Displays the last record
    now show -c n #Displays last n record