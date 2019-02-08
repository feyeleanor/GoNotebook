#include <stdlib.h>
#include <string.h>

int accumulate(int x) {
  static int y;

  y += x;
  return y;
}

int main(int argc, char **argv) {
  int i;

  for (i = 1; i < argc; i++) {
    accumulate(atoi(*(++argv)));
  }
  return accumulate(0);
}
