#include <stdlib.h>
#include <string.h>

int add(int x, int y) {
  return x + y;
}

int main(int argc, char **argv) {
  int i;
  int sum = 0;

  for (i = 1; i < argc; i++) {
    sum = add(sum, atoi(*(++argv)));
  }
  return sum;
}
