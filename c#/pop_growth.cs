using System;

class Arge {

    public static int NbYear(int p0, double percent, int aug, int p) {
      int pop = p0;
      int count = 0;
      while (pop < p) {
        count++;
        pop = pop + (int)(pop * (percent/100)) + aug;
      }
      return count;
    }
}
