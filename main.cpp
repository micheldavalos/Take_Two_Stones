#include <iostream>

using namespace std;

int main() {
 long long stones;
 cin >> stones;
 
 stones % 2 != 0 ? cout << "Alice" : cout << "Bob";
 
 return 0;
    
}
