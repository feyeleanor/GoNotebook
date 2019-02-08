#include <stdlib.h>
#include <string.h>

int add(int x, int y) {
  return x + y;
}

int main(int argc, char **argv) {
  switch (argc) {
  case 1:
    return 0;
  case 2:
    return add(0, atoi(*(++argv)));
  case 3:
    return add(atoi(*(++argv)), atoi(*(++argv)));
  default:
    abort();
  }
}
