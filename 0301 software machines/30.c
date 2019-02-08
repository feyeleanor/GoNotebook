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

#define COMPILE_NEXT_OPCODE goto *compiler[*PC++];
#define COMPILE_NUMBER \
  *cp++ = (void *)(long)*PC++; \
  words--;
#define PRIMITIVE_ADDRESS(op) *cp++ = despatch_table[op];
#define COMPILE_PRIMITIVE(name, body) \
  c_##name: \
    PRIMITIVE_ADDRESS(name) \
    body; \
    words--; \
    if (words == 0) return program; \
    COMPILE_NEXT_OPCODE

typedef enum { PUSH = 0, ADD, PRINT, EXIT } opcodes;

STACK *S;

void **compile(int *PC, int words, void *despatch_table[]) {
  static void *compiler [] = {
    &&c_PUSH, &&c_ADD, &&c_PRINT, &&c_EXIT
  };

  if (words < 1) return NULL;
  void **program = malloc(sizeof(void *) * words);
  void **cp = program;
  COMPILE_NEXT_OPCODE

  COMPILE_PRIMITIVE(PUSH, COMPILE_NUMBER)
  COMPILE_PRIMITIVE(ADD, ;)
  COMPILE_PRIMITIVE(PRINT, ;)
  COMPILE_PRIMITIVE(EXIT, ;)
}

#define NEXT_OPCODE goto **program++;
#define PRIMITIVE(name, body) \
  op_##name: \
    body; \
    NEXT_OPCODE

void interpret(int *PC, int words) {
  static void *despatch_table[] = {
    &&op_push, &&op_add, &&op_print, &&op_exit
  };

  int l, r;
  void **program = compile(PC, words, despatch_table);
  if (program == NULL) exit(1);
  NEXT_OPCODE

  PRIMITIVE(push, S = push(S, (int)(long)*program++))
  PRIMITIVE(add, \
    S = pop(S, &l); \
    S = pop(S, &r); \
    S = push(S, l + r); \
  )

  PRIMITIVE(print, printf("%d + %d = %d\n", l, r, S->data))
  PRIMITIVE(exit, return)
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
