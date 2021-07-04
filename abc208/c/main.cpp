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
    int n;
    long k;
    cin >> n >> k;

    int a[n], b[n];
    REP(i, n) {
        cin >> a[i];
        b[i] = a[i];
    }

    long base = k / n, mod = k % n;
    if (mod == 0) {
        REP(i, n) { cout << base << endl; }
        return 0;
    }

    sort(b, b + n);

    long tar = b[mod - 1];
    REP(i, n) {
        if (a[i] <= tar) {
            cout << base + 1 << endl;
        } else {
            cout << base << endl;
        }
    }

    return 0;
}
