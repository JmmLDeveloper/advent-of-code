#include <iostream>
#include <fstream>
#include <string>

int correct_length( std::string& line ); 
std::string encode_string( std::string& line ); 

int main () {
    std::ifstream input1;
    input1.open("input.txt");

    std::string line;

    size_t result1 = 0;
    size_t result2 = 0;


    while (  input1 >> line  ) { 
        result1 += line.length() - correct_length(line);
        result2 += encode_string(line).length() - line.length();
    }



    std::cout << "Answer to part 1 is :" <<  result1 << std::endl;
    std::cout << "Answer to part 2 is :" <<  result2 << std::endl;

}

//this function assumes the input is correct of it my crash
int correct_length( std::string& line ) {
    int count = 0;
    bool scaping_hex = false;
    bool scaping_backslash = false;

    // we iterate through 1 to len - 1 to ignore the ""
    for (int i = 1; i < line.length() - 1 ; i++) {
        count++;
        char ch = line[i];

        if ( ch == '\\'   ) {
            // the first check prevent a crash
            if ( line.length() > i + 1  && line[i + 1] == 'x' ) {
                i += 3;
            } else if ( line.length() > i + 1  && (line[i + 1] == '\\' || line[i + 1] == '"' )  ) {
                i++;
            }
        }
    }
    return count;
}

std::string encode_string( std::string& line ){
    std::string new_string = "\"";

    for (char& c : line) {
        if ( c == '\\' || c == '"') {
            new_string.push_back('\\');
        }
        new_string.push_back(c);
    }

    new_string.append("\"");

    return new_string;
} 


