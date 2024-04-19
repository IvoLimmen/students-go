# Students

Just a simple application to see how to build something in Go.

## Requirements

* Go 1.21.1

## Build
```
go build students.go
```
And run the executable:
```
./students
```
## Run
Using the interpreter:
```
go run students.go
```

Type `help` for the list of commands.
When you type `end` or `quit` it will show the averages per course per student.

# Evaluation

The good:

  * It works?

The bad:

  * Having to use to statements for printing out in a cross platform (although I did not test that) way (Printf followed by Println(""))
  * Reading input *until* a \n this is just...

The ugly:

  * You can fake OO by creating function in a different way; works but not my taste
  * Capitalization of the method names is just not my thing
  * No stream like API on List
  * The check on almost anything to see if the value is nil (error handling)
