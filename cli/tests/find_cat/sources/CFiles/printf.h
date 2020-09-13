/*
** EPITECH PROJECT, 2019
** Minimal printf (Usable)
** File description:
** header for printf_lib
*/

//<===============================================>
// Essentials Defines.
#ifndef UNUSED
#define UNUSED __attribute__ ((unused))
#endif /* UNUSED */

#ifndef HIDDEN
#define HIDDEN __attribute__ ((visibility ("hidden")))
#endif /* HIDDEN */


#ifndef DESTRUCTOR
#define DESTRUCTOR __attribute__((destructor))
#endif /* DESTRUCTOR */

#ifndef CONSTRUCTOR
#define CONSTRUCTOR __attribute__((constructor))
#endif /* CONSTRUCTOR */

#ifndef _PRINTF_
#define _PRINTF_

//<===============================================>
//<===============================================>
// Necessary Header Inclusion.
#include <stdarg.h>
#include <stddef.h>
#include <unistd.h>

//<===============================================>
//<===============================================>
// Printf flag functions.
static inline size_t printf_int(va_list list);
static inline size_t printf_percent(va_list list);
static inline size_t printf_char(va_list list);
static inline size_t printf_str(va_list list);
static inline size_t printf_uint(va_list list);

//<===============================================>
//<===============================================>
// Library (trivial) Functions Prototypes (definition is below).
static inline size_t my_print_strlen(char const *str);
static inline size_t my_putchar(char c);
static inline size_t my_putstr(char const *str);
static inline size_t my_putstr_err(char const *str);
static inline size_t my_put_nbr(long long nb);
static inline size_t my_put_nbr_base(unsigned int nb, char const *base);
static inline size_t printf_word_array(va_list argv_list);

//<===============================================>
//<===============================================>
// Associative Array.
typedef struct printf_array_s {
    char flag;
    size_t (*fptr)(va_list argv_list);
} printf_array_t;

static const printf_array_t printf_array[] = {
    {'d', printf_int},
    {'i', printf_int},
    {'u', printf_uint},
    {'%', printf_percent},
    {'c', printf_char},
    {'s', printf_str},
    {'w', printf_word_array}
};

static const int PRINTF_ARRAY = sizeof(printf_array) / sizeof(printf_array_t);

static inline size_t check_flag(va_list argv_list, char const flag)
{
    size_t printed = 0;

    for (size_t i = 0; i < (size_t)PRINTF_ARRAY; i += 1) {
        if (printf_array[i].flag == flag)
            printed += printf_array[i].fptr(argv_list);
    }
    return printed;
}

//<===============================================>
//<===============================================>
// Printf flag functions.
static inline size_t printf_int(va_list argv_list)
{
    size_t char_printed = my_put_nbr(va_arg(argv_list, int));
    return (char_printed);
}

static inline size_t printf_percent(va_list argv_list)
{
    (void)va_arg(argv_list, void *);

    my_putchar('%');
    return (1);
}

static inline size_t printf_char(va_list argv_list)
{
    char c;

    c = (char)va_arg(argv_list, int);
    my_putchar(c);
    return (1);
}

static inline size_t printf_str(va_list argv_list)
{
    size_t char_printed = my_putstr(va_arg(argv_list, char *));
    return (char_printed);
}

static inline size_t printf_word_array(va_list argv_list)
{
    size_t char_printed = 0;
    char **word_array = va_arg(argv_list, char **);
    for (; *word_array; word_array++) {
        char_printed += my_putstr(*word_array);
        my_putchar('\n');
    }
    return (char_printed);
}

static inline size_t printf_uint(va_list argv_list)
{
    size_t char_printed = my_put_nbr((long long)va_arg(argv_list,
                                                    unsigned int));
    return (char_printed);
}

static inline size_t my_printf(char const *str, ...)
{
    va_list argv_list;
    size_t printed = 0;

    va_start(argv_list, str);
    for (; *str; str++)
        printed += *str == '%' ? check_flag(argv_list, *++str) :
            (size_t)write(1, str, 1);
    va_end(argv_list);
    return printed;
}

//<===============================================>
//<===============================================>
// Library (trivial) Functions Definitions.
static inline size_t my_putchar(char c)
{
    write(1, &c, 1);
    return (1);
}

static inline size_t my_putstr(char const *str)
{
    size_t i = 0;

    for (i = 0; str[i] != '\0'; i++);
    write(1, str, i);
    return (i);
}

static inline size_t my_putstr_err(char const *str)
{
    size_t i = 0;

    for (i = 0; str[i] != '\0'; i++);
    write(2, str, i);
    return (i);
}

static inline size_t my_put_nbr(long long nb)
{
    size_t printed = 0;

    if (nb > 9)
        printed += my_put_nbr(nb / 10);
    else if (nb < 0) {
        nb *=-1;
        write(1, "-", 1);
        printed += 1;
        printed += my_put_nbr(nb / 10);
    }
    printed += my_putchar(nb % 10 + '0');
    return printed;
}

static inline size_t my_put_nbr_base(unsigned int nb, char const *base)
{
    size_t base_len = my_print_strlen(base);
    size_t divisor = 1;
    size_t count = 0;

    while ((nb / divisor) >= base_len)
        divisor *= base_len;
    while (divisor > 0) {
        count += my_putchar(base[(nb / divisor) % base_len]);
        divisor /= base_len;
    }
    return (count);
}

static inline size_t my_print_strlen(char const *str)
{
    size_t i = 0;

    for (i = 0; str[i] != '\0'; i++);
    return (i);
}

#endif /* _PRINTF_ */
