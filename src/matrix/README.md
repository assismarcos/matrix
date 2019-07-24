##How to build
To build this project run command 'go install' in the project directory.
This command withh build and create a binary 

##How to run
You can run this project in many ways.
1. Navigate to where the directory bin is created. 
Run command './matrix' which fires the binary
you should see 'Server is starting...' in the console logs

2. Go to project directory and run command 'go run main.go string_ops.go math_ops.go'

##How to test
From the terminal run command 
curl -F 'file=@<path>/matrix.csv' "localhost:8080/<function>"

For example
curl -F 'file=@<path>/matrix.csv' "localhost:8080/invert"
curl -F 'file=@<path>/matrix.csv' "localhost:8080/flatten"
curl -F 'file=@<path>/matrix.csv' "localhost:8080/sum"
curl -F 'file=@<path>/matrix.csv' "localhost:8080/multiply"

##Error handling
1. If the input csv has elements other than numbers you will get an error as follows for the math operations:
"Invalid input. Matrix does not contain valid integers"
2. If you send an unhandeled URI in the request you will get the following error:
"Invalid operation".
 