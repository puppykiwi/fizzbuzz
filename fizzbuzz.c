#include <stdio.h>
#include <stdlib.h>

int max_num;
int main(int argc, char* argv[]){
    if (argc>1){max_num = atoi(argv[1]);}
    else{max_num = 100;}

    for (int i=1; i < (max_num+1); i++){
        if((i % 3 == 0) && (i % 5 == 0)){printf("%d - FizzBuzz\n",i);}
        else if (i % 3 == 0){printf("%d - Fizz\n",i);}
        else if(i % 5 == 0){printf("%d - Buzz\n",i);}
        else{printf("%d\n",i);}
    }

    return 1;
}