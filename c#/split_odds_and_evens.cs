using System;
using System.Collections.Generic;

public class Kata
{
  public static long[] SplitOddAndEven(long n)
  {
    string digits = n.ToString();
    bool evens = false;
    List<long> groups = new List<long>{};
    string grp = "";
    for (int i=0; i<digits.Length; i++) {
      char digit = digits[i];
      bool isEven = digit%2==0;

      if (i==0) {
        evens=isEven;
        grp = digit.ToString();
      } else if (evens == isEven) {
        grp += digit.ToString();
      } else {
        groups.Add(Convert.ToInt64(grp));
        evens=isEven;
        grp = digit.ToString();
      }
    }
    if (grp.Length > 0) {
      groups.Add(Convert.ToInt64(grp));
    }
    return groups.ToArray();
  }
}