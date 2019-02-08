#include <stdlib.h>
#include <string.h>

struct assoc_array {
  char *key;
  void *value;
  struct assoc_array *next;
};
typedef struct assoc_array assoc_array_t;

assoc_array_t *assoc_array_new(char *k, char *v) {
  assoc_array_t *a = malloc(sizeof(assoc_array_t));
  a->key = strdup(k);
  a->value = strdup(v);
  a->next = NULL;
  return a;
}

char *assoc_array_get_if(assoc_array_t *a, char *k) {
  char *r = NULL;
  if (a != NULL && strcmp(a->key, k) == 0) {
    r = strdup(a->value);
  }
  return r;
}
