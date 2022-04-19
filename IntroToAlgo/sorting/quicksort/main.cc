#include <iostream>
#include <vector>
using namespace std;

void printVector(vector<int> v) {
  for (int i : v) {
    cout << i << " ";
  }
  cout << endl;
}

int partition(vector<int> &v, int l, int r) {
  // A: less than p [l, i+1)
  // B: greater than p [i+1, j)
  // C: unsorted [j, r)
  // pivot: r
  // l     i       j       r
  // 0 1 2 3 5 6 7 3 5 6 2 4
  // len(A) == len(B) == 0

  int p = v[r];
  int i = l - 1, j = l, tmp;
  for (j = l; j < r;) {
    if (v[j] <= p) {
      i++;
      tmp = v[i];
      v[i] = v[j];
      v[j] = tmp;
    }
    j++;
  }
  i++;
  tmp = v[i];
  v[i] = v[r];
  v[r] = tmp;

  return i;
}

void quicksort(vector<int> &v, int l, int r) {
  if (l >= r) {
    return;
  }
  int m = partition(v, l, r);
  quicksort(v, l, m - 1);
  quicksort(v, m + 1, r);
}

int main() {
  vector<int> origin = {2,    4,  1,      4,    2,   6,  2,   326,  0,
                        -352, 26, -23434, 5312, 42,  5,  436, -346, 34,
                        -4,   36, 34,     6,    346, 34, 63};
  printVector(origin);
  quicksort(origin, 0, origin.size() - 1);
  printVector(origin);
  return 0;
}
