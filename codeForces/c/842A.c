#include <stdio.h>

int main() {
  int l, r, x, y, ef;
  scanf("%d %d %d %d %d", &l, &r, &x, &y, &ef);

  float xp_mid, c_mid, find = 0;
  for (float i = ef; i <= r; i += ef) {
    if (i < l)
      continue;
    else {
      int xx = x, yy = y;
      while (xx <= yy) {
        c_mid = (xx + yy) / 2;

        if (ef > i / c_mid) {
          yy = c_mid - 1;
        } else if (ef < i / c_mid) {
          xx = c_mid + 1;
        } else {
          find = 1;
          break;
        }
      }
      if (find)
        break;
    }
  }

  printf(find ? "YES\n" : "NO\n");

  return 0;
}
