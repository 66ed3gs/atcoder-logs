#include <bits/stdc++.h>

#define REP(i, n) for (int i = 0; i < n; i++)
#define REPR(i, n) for (int i = n - 1; i >= 0; i--)
#define FOR(i, m, n) for (int i = m; i < n; i++)
#define ALL(v) v.begin(), v.end()

typedef long long ll;

// sort(ALL(a))
// sortALL(a), greater<>())
// int max = *max_element(ALL(a));
// int min = *min_element(ALL(a));

using namespace std;

signed main() {
    int a, b;
    cin >> a >> b;

    if (a * 6 < b || a * 1 > b) {
        cout << "No" << endl;
    } else {
        cout << "Yes" << endl;
    }

    return 0;
}
