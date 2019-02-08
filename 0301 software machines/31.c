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

#define COMPILE(body) \
  COMPILE_NEXT_OPCODE \
  body
#define COMPILE_NEXT_OPCODE \
  if (words < 1) \
    return program; \
  goto *compiler[*PC++];
#define DESCRIBE_PRIMITIVE(name, body) \
name: \
  body; \
  COMPILE_NEXT_OPCODE
#define WRITE_OPCODE(value) \
  *cp++ = value; \
  words--;

void **compile(int *PC, int words, void *despatch_table[]) {
  static void *compiler[] = {
    &&push,
    &&add,
    &&print,
    &&exit
  };

  void **program = malloc(sizeof(void *) * words);
  void **cp = program;
  COMPILE( \
    DESCRIBE_PRIMITIVE(push, \
      WRITE_OPCODE(despatch_table[PUSH]) \
      WRITE_OPCODE((void *)(long)*PC++)) \
    DESCRIBE_PRIMITIVE(add, \
      WRITE_OPCODE(despatch_table[ADD])) \
    DESCRIBE_PRIMITIVE(print, \
      WRITE_OPCODE(despatch_table[PRINT])) \
    DESCRIBE_PRIMITIVE(exit, \
      WRITE_OPCODE(despatch_table[EXIT])) \
  )
}

#define EXECUTE(body) \
  EXECUTE_OPCODE \
  body
#define READ_OPCODE *program++
#define EXECUTE_OPCODE goto *READ_OPCODE;
#define PRIMITIVE(name, body) \
name: \
  body; \
  EXECUTE_OPCODE

void interpret(int *PC, int words) {
  static void *despatch_table[] = {
    &&push,
    &&add,
    &&print,
    &&exit
  };

  int l, r;
  void **program = compile(PC, words, despatch_table);
  if (program == NULL)
    exit(1);
  EXECUTE(\
    PRIMITIVE(push, S = push(S, (int)(long)READ_OPCODE)) \
    PRIMITIVE(add, \
      S = pop(S, &l); \
      S = pop(S, &r); \
      S = push(S, l + r))
    PRIMITIVE(print, printf("%d + %d = %d\n", l, r, S->data)) \
    PRIMITIVE(exit, return) \
  )
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
