Todoq
-----

An extremely simple CLI app to help you track your TODOs as a queue of tasks.

# Usage
`todoq` is completely interactive, so just run it:
```
$ todoq
```
and use commands to manipulate your task queue:
- `list` - List your current tasks
- `push [TASK]` - Push a task on top of the queue _(name shouldn't contain whitespace!)_
- `pop` - Pop the bottom task off the queue
- `q`/`quit`/`exit` - Quit the program _(naturally)_

## Examples
```
> list
Current task queue:

> push math-hw
> list
Current task queue:
- math-hw

> push dishes
> list
Current task queue:
- math-hw
- dishes

> pop
Task "math-hw" completed!
> list
Current task queue:
- dishes

> q
```

After you exit, `todoq` saves the current task list in `$HOME/todoq.tsk` on Linux and alike, undefined behaviour on Windows (for now).
