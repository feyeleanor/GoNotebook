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

void **compile(int *PC, int words, void *despatch_table[]) {
  static void *compiler [] = {
    &&c_push,
    &&c_add,
    &&c_print,
    &&c_exit
  };

  if (words < 1) return NULL;
  void **program = malloc(sizeof(void *) * words);
  void **cp = program;
  goto *compiler[*PC++];

c_push:
  *cp++ = despatch_table[PUSH];
  *cp++ = (void *)(long)*PC++;
  words -= 2;
  if (words == 0) return program;
  goto *compiler[*PC++];

c_add:
  *cp++ = despatch_table[ADD];
  words--;
  if (words == 0) return program;
  goto *compiler[*PC++];

c_print:
  *cp++ = despatch_table[PRINT];
  words--;
  if (words == 0) return program;
  goto *compiler[*PC++];

c_exit:
  *cp++ = despatch_table[EXIT];
  words--;
  if (words == 0) return program;
  goto *compiler[*PC++];
}

void interpret(int *PC, int words) {
  static void *despatch_table[] = {
    &&op_push,
    &&op_add,
    &&op_print,
    &&op_exit
  };

  int l, r;
  void **program = compile(PC, words, despatch_table);
  if (program == NULL) exit(1);
  goto **program++;

op_push:
  S = push(S, (int)(long)*program++);
  goto **program++;

op_add:
  S = pop(S, &l);
  S = pop(S, &r);
  S = push(S, l + r);
  goto **program++;

op_print:
  printf("%d + %d = %d\n", l, r, S->data);
  goto **program++;

op_exit:
  return;
}

int main() {
  int program[] = {
    PUSH, 13,
    PUSH, 28,
    ADD,
    PRINT,
    EXIT
  };
  interpret(program, 7);
}
