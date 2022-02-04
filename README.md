# Survey Test

Example for https://github.com/AlecAivazis/survey/issues/406

To reproduce:

```
make test
```

This will run tests in a loop 100 times and some of them will fail with the following error

```
--- FAIL: TestAskQuestion (0.02s)
    main_test.go:69: Unexpected output.
        Expected: 
        ? Do you like pizza? Yes                                                        
        Answer true ; 
        Found: 
        ? Do you like pizza? (y/N) y                                                    
        ? Do you like pizza? Yes                                                        
        Answer true
```

Tested with go version go1.17.6 linux/amd64 in Ubuntu 21.10