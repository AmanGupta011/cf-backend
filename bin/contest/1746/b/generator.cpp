#include <bits/stdc++.h>
using namespace std;

int randomNumberBetween(int a, int b)
{
    if (a > b)
        swap(a, b);
    return a + rand() % (b - a + 1);
}
int main(int argc, char *argv[])
{
    srand(atoi(argv[1]));
    cout << "1" << endl;
    int n = randomNumberBetween(1, 50);
    int k = randomNumberBetween(1, 1000000000);
    cout << n << ' ' << k << endl;
    for (int i = 1; i <= n; i++)
    {
        int x = randomNumberBetween(1, 1000000000);
        int y = randomNumberBetween(1, 1000000000);
        cout << x << ' ' << y;
        if (i != n)
        {
            cout << endl;
        }
    }
    cout << endl;
}