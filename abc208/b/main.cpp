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
    int p;
    cin >> p;

    int coin[10];
    coin[0] = 1;
    FOR(i, 1, 10) { coin[i] = coin[i - 1] * (i + 1); }

    int res = 0;
    REPR(i, 10) {
        if (p < coin[i]) {
            continue;
        }
        res += int(p / coin[i]);
        p = p % coin[i];
    }

    cout << res << endl;

    return 0;
}
