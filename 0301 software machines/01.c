#include <stdio.h>
#include <stdlib.h>
#define STACK_MAX 1024

typedef enum {
  STACK_OK = 0,
  STACK_OVERFLOW,
  STACK_UNDERFLOW
} STACK_STATUS;

typedef struct stack STACK;
struct stack {
  int data[STACK_MAX];
  int depth;
};

STACK *NewStack() {
  STACK *s;
  s = malloc(sizeof(STACK));
  s->depth = 0;
  return s;
}

STACK_STATUS push(STACK *s, int data) {
  if (s->depth < STACK_MAX) {
    s->data[s->depth++] = data;
    return STACK_OK;
  }
  return STACK_OVERFLOW;
}

STACK_STATUS pop(STACK *s, int *r) {
  if (s->depth > 0) {
    *r = s->data[s->depth - 1];
    s->depth--;
    return STACK_OK;
  }
  return STACK_UNDERFLOW;
}

int main() {
  int l, r;
  STACK *s = NewStack();
  push(s, 1);
  push(s, 3);
  printf("depth = %d\n", s->depth);
  pop(s, &l);
  pop(s, &r);
  printf("%d + %d = %d\n", l, r, l + r);
  printf("depth = %d\n", s->depth);
}
