const process = require('process');
const argv = process.argv;

var max_num = 50;
const fizzbuzz = (max_num) => {
    if (argv.length > 2 && parseInt(argv[2], 10) > 0){
        max_num = argv[2];
    }

    for (i = 1; i < (max_num + 1); i++){
        if (i % 3 == 0 && i % 5 == 0){
            console.log(i,"- FizzBuzz");
        }

        else if (i % 3 == 0){
            console.log(i,"- Fizz");
        }

        else if (i % 5 == 0){
            console.log(i,"- Buzz");
        }

        else{
            console.log(i);
        }
    }
}

fizzbuzz(max_num);