fizzbuzz = unda(nambari=100){
    fanya i = 1
    wakati (i < (nambari + 1)){
            kama (i % 15 == 0){
                andika(i, "- FizzBuzz")
            }

            au kama(i % 3 == 0){
                andika(i,"- Fizz")
            }

            au kama(i % 5 == 0){
                andika(i,"- Buzz")
            }

            sivyo{
                andika(i)
            }
            i++
    }
}

fizzbuzz()