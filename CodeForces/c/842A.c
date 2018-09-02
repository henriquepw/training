#include <stdio.h>

int main() {
  int l, r, x, y, ef, count = 0;
  scanf("%d %d %d %d %d", &l, &r, &x, &y, &ef);

  int xp_mid, c_mid, find = 0;
  while (l <= r) {
    xp_mid = (l + r) / 2;
    while (x <= y) {
      c_mid = (x + y) / 2;

      if (ef > xp_mid / c_mid) {
        y = c_mid - 1;
      } else if (ef < xp_mid / c_mid) {
        x = c_mid + 1;
      } else if (ef == xp_mid / c_mid) {
        printf("YES\n");
        find = 1;
        break;
      }
    }

    if (find)
      break;
    else {
      if (ef > xp_mid / c_mid) {
        l = xp_mid + 1;
      } else if (ef < xp_mid / c_mid) {
        r = xp_mid - 1;
      } else if (ef == xp_mid / c_mid) {
        printf("YES\n");
        find = 1;
        break;
      }
    }
  }

  if (!find)
    printf("NO\n");

  return 0;
}
