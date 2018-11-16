### Contribute:

Simply make a pull request, I have yet to turn one down.

**NOTE:** Currently, I am just relying on TODOS in the comments of the code, as a temporary (as in, will change) replacement for 'issues'



**IMPORTANT:** When adding a new task, you must follow this pattern!

1. Create a new file in the *tasks* directory and write all of your code there.

3. If you feel any code in your class may be used in other tasks, feel free to put it in `utils.go`.

4. Ensure your code is documented well (running *golint* is helpful).

5. New tasks should have an associated test file (e.g. `mytask_test.go`) in the same folder.



If the new feature is complete:

1. Add the case to the switch statement in `main.go`, so your new task can be called.

2. Finished!

