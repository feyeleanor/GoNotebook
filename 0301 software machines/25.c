#include <stdio.h>
#include <stdlib.h>

typedef struct stack STACK;
struct stack {
  int data;
  STACK *next;
};

STACK *push(STACK *s, int data) {
  STACK *r = malloc(sizeof(STACK));
  r->data = data;
  r->next = s;
  return r;
}

STACK *pop(STACK *s, int *r) {
  if (s == NULL)
    exit(1);
  *r = s->data;
  return s->next;
}

typedef enum { PUSH = 0, ADD, PRINT, EXIT } opcodes;

STACK *S;

void interpret(int *program) {
  static void *opcodes [] = {
    &&op_push,
    &&op_add,
    &&op_print,
    &&op_exit
  };

  int l, r;
  int *PC = program;
  goto *opcodes[*PC++];

op_push:
  S = push(S, *PC++);
  goto *opcodes[*PC++];

op_add:
  S = pop(S, &l);
  S = pop(S, &r);
  S = push(S, l + r);
  goto *opcodes[*PC++];

op_print:
  printf("%d + %d = %d\n", l, r, S->data);
  goto *opcodes[*PC++];

op_exit:
  return;
}

int main() {
  int program [] = {
    PUSH, 13,
    PUSH, 28,
    ADD,
    PRINT,
    EXIT
  };
  interpret(program);
}
