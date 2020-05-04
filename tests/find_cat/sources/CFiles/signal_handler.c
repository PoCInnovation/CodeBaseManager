/*
** EPITECH PROJECT, 2019
** Minishell_1
** File description:
** SOurce file that handle signals ?
*/

#include <signal.h>
#include "minishell.h"
#include "printf.h"

static void catch_sigint(int signal);

void signal_handler(UNUSED char *cmd)
{
    signal(SIGINT, catch_sigint);
}

static void catch_sigint(int signal)
{
    if (signal == SIGINT) {
        my_printf("\n");
        display_prompt();
    }
}
