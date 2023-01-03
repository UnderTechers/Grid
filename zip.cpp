#include<stdlib.h>
#include<stdio.h>
using namespace std;
int main(){
    FILE *f = fopen("test.cmd","r+");
    char str[100];

    fgets(str,100,f);
    system(str);
    
}