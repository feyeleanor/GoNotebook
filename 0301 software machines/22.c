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

typedef void (*opcode)();

STACK *S;
opcode *PC;

void op_push() {
  S = push(S, (int)(long)(*PC++));
}

void op_add_and_print() {
  int l, r;
  S = pop(S, &l);
  S = pop(S, &r);
  S = push(S, l + r);
  printf("%d + %d = %d\n", l, r, S->data);
}

void op_exit() {
  exit(0);
}

int main() {
  opcode program [] = {
    op_push, (opcode)(long)13,
    op_push, (opcode)(long)28,
    op_add_and_print,
    op_exit
  };
  PC = program;
  while (1) {
    (*PC++)();
  }
}
