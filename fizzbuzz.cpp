#include <iostream>

int getNum(){
    std::cout << "Enter max NO: " ;
    int num;
    std::cin >> num;
    return num;
}

int main(){
    int num = getNum();
    // std::cout << "num: " << num << std::endl; // debug
    
    for (int i =0, i < num, i++) {
        
        if ((i % 3 = 0) && (i % 5 =0)){
            std::cout << "FizzBuzz" << std::endl;

        }
    }
    return 0;
}