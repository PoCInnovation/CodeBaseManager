int main(int ac, char **av)
{
    av += ac;
    while (*++av)
        puts(*av);
}
