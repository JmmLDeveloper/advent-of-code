#include <iostream>
#include <string>

std::string look_and_say(const std::string& input);


int main() {
    const std::string input = "1113122113";
    std::string result1;
    std::string result2;

    for ( int i = 0 ; i < 40 ;i++ ) {
        result1 = look_and_say(  i == 0 ? input : result1 );
    }   

    for ( int i = 0 ; i < 10 ;i++ ) {
        result2 = look_and_say(  i == 0 ? result1 : result2 );
    }   

    std::cout <<  "The answer of part 1 is :" << result1.length() << std::endl;
    std::cout <<  "The answer of part 2 is :" << result2.length() << std::endl;

}

std::string look_and_say(const std::string& input) {

    std::string output = "";
    char prev = '\0';
    int cont = 0;

    for (auto ch : input ) {
        if ( prev != ch && cont != 0 ) {
            output.push_back( '0' + cont ); // adds the number of numbers
            output.push_back( prev ); // adds the number repetead n times
            cont = 0;
        }

        prev = ch;
        cont++;
    }

    if ( cont != 0 ) {
        output.push_back( '0' + cont ); // adds the number of numbers
        output.push_back( prev ); // adds the number repetead n times
    }



    return output;
}