from sys import argv as args #get command line arguments

def fizzbuzz(max_num = 100 ):
    if len(args) > 1:
        max_num = int(args[1])
        
    for i in range(1, max_num + 1):
        if i % 3 == 0 and i % 5 == 0:
            print(i,"- FizzBuzz")
        elif i % 3 == 0:
            print(i,"- Fizz")
        elif i % 5 == 0:
            print(i,"- Buzz")
        else:
            print(i)

fizzbuzz()
