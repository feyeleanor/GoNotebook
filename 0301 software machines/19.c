#include <stdio.h>
#include "stack.c"

typedef enum { PUSH = 0, ADD, PRINT, EXIT } opcodes;

STACK *S;

void interpret(int *PC) {
  int l, r;
  while (1) {
    switch(*PC++) {
      case PUSH:
        S = push(S, *PC++);
        break;
      case ADD:
        S = pop(S, &l);
        S = pop(S, &r);
        S = push(S, l + r);
        break;
      case PRINT:
        printf("%d + %d = %d\n", l, r, S->data);
        break;
      case EXIT:
        return;
    }
  }
}

int main() {
  int program [] = {
    (int)PUSH, 13,
    (int)PUSH, 28,
    (int)ADD,
    PRINT,
    EXIT,
  };
  interpret(program);
}
