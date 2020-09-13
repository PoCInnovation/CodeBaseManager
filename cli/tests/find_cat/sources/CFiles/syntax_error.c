/*
** EPITECH PROJECT, 2019
** Minishell_1
** File description:
** Source file that manage syntax error in user input.
*/

#include "user_synthax.h"
#include "printf.h"
#include "globals.h"

static int fill_infos(char c, sc_t *info);
static int check_input_synthax(sc_t *info, char *ptr);
static int check_quotes(sc_t *info);
static int check_input_synthax(sc_t *info, char *ptr);

int check_user_input(char **input)
{
    char *ptr = *input;
    sc_t info = fill_info_struct();

    for (size_t i = 0; *(ptr + i); i += 1) {
        if (special_chars(*(ptr + i)))
            fill_infos(*(ptr + i), &info);
    }
    if (check_input_synthax(&info, ptr) == ERROR) {
        SET_RTC(LAST_CMD_ERROR);
        return ERROR;
    }
    return SUCCESS;
}

static int fill_infos(char c, sc_t *info)
{
    QUOTE += (quote(c) ? 1 : 0);
    DBQUOTE += (dbl_quote(c) ? 1 : 0);
    LPAR += (l_parenth(c) ? 1 : 0);
    RPAR += (r_parenth(c) ? 1 : 0);
    if (ORDER == '\0' && (quote(c) || dbl_quote(c)))
        ORDER = c;
    return (84);
}

static int check_quotes(sc_t *info)
{
    if ((ORDER == '\'') && ((QUOTE % 2) != 0)) {
        my_printf("Unmatched '''.\n");
        return ERROR;
    }
    if ((ORDER == '\"') && ((DBQUOTE % 2) != 0)) {
        my_printf("Unmatched '\"'.\n");
        return ERROR;
    }
    if ((QUOTE % 2) != 0) {
        my_printf("Unmatched '''.\n");
        return ERROR;
    }
    if (((DBQUOTE % 2) != 0)) {
        my_printf("Unmatched '\"'.\n");
        return ERROR;
    }
    return SUCCESS;
}

static int check_paren(sc_t *info)
{
    if (LPAR > RPAR) {
        my_printf("Too many ('s.\n");
        return ERROR;
    }
    if (RPAR > LPAR) {
        my_printf("Too many )'s.\n");
        return ERROR;
    }
    return SUCCESS;
}

static int check_input_synthax(sc_t *info, char *ptr)
{
    if (check_quotes(info)) {
        free(ptr);
        return ERROR;
    }
    if (check_paren(info)) {
        free(ptr);
        return ERROR;
    }
    return SUCCESS;
}
