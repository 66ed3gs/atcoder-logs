# Sort

Vectorを並び替え
```cpp
sort(ALL(a))
sort(ALL(a), greater<>())
```

# Max / Min

Vectorの要素
```cpp
int max = *max_element(ALL(a));
int min = *min_element(ALL(a));
```


# Priority Queue

昇順
```cpp
priority_queue<int, vector<int>, greater<int> > q;
```

ソート条件の指定
```cpp
auto c = [](int l, int r) { if (l % 2 == 0) return r % 2 != 0 || l >= r; else
return r % 2 != 0 && l < r; }; priority_queue<int, vector<int>, decltype(c)> q(c);
```

値を追加
```cpp
q.push()    // 値を追加
q.top()     // 先頭を取得
q.pop()     // 先頭を削除
```

