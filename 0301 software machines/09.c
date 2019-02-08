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
  if (s == NULL) exit(1);
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

int sum(STACK *tos) {
  int a = 0;
  for (int p = 0; tos != NULL;) {
    tos = pop(tos, &p);
    a += p;
  }
  return a;
}

void print_sum(STACK *s) {
  printf("%d items: sum = %d\n", depth(s), sum(s));
}

int main() {
  STACK *s1 = push(NULL, 7);
  STACK *s2 = push(push(s1, 7), 11);

  s1 = push(push(push(s1, 2), 9), 4);
  STACK *s3 = push(s1, 17);

  s1 = push(s1, 3);
  print_sum(s1);
  print_sum(s2);
  print_sum(s3);
}
