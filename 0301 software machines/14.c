#include <stdlib.h>
#include "assoc_array.c"

struct search {
  char *term, *value;
  assoc_array_t *cursor, *memo;
};
typedef struct search search_t;

search_t *search_new(assoc_array_t *a, char *k) {
  search_t *s = malloc(sizeof(search_t));
  s->term = k;
  s->value = NULL;
  s->cursor = a;
  s->memo = NULL;
  return s;
}

void search_step(search_t *s) {
  s->value = assoc_array_get_if(s->cursor, s->term);
}

int searching(search_t *s) {
  return s->value == NULL && s->cursor != NULL;
}

search_t *search_find(assoc_array_t *a, char *k) {
  search_t *s = search_new(a, k);
  for (search_step(s); searching(s); search_step(s)) {
    s->memo = s->cursor;
    s->cursor = s->cursor->next;
  }
  return s;
}
