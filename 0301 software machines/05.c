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

int depth(STACK *s) {
  int r = 0;
  for (STACK *t = s; t != NULL; t = t->next) {
    r++;
  }
  return r;
}

void gc(STACK **old, int items) {
  STACK *t;
  for (; items > 0 && *old != NULL; items--) {
    t = *old;
    *old = (*old)->next;
    free(t);
  }
}

int main() {
  int l, r;
  STACK *s = push(NULL, 1);
  s = push(s, 3);
  printf("depth = %d\n", depth(s));
  STACK *t = pop(pop(s, &r), &l);
  printf("%d + %d = %d\n", l, r, l + r);
  printf("depth pre-gc  = %d\n", depth(s));
  gc(&s, 2);
  printf("depth post-gc = %d\n", depth(s));
}
