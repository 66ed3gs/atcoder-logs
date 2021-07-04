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
    long n;
    int k;
    cin >> n >> k;

    long ans;
    vector<int> q = {1, 2, 3, 4, 5, 6, 7, 8, 9};
    for (long count = 0; count < n; count++) {
        long x = q[0];
        q.erase(q.begin());

        if (x <= k) {
            REP(l, 10) { q.push_back(x * l); }
            ans += 1;
        }
    }

    cout << ans << endl;

    return 0;
}
