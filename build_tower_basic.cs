using System.Collections.Generic;
using System;

public class Kata
{
  public static string[] TowerBuilder(int nFloors)
  {
    List<string> str = new List<string>();
    for (int i=nFloors-1; i>=0; i--) {
      string newStr = new String('*', 1+(2*(nFloors-1-i)));
      string padding = new string(' ', i);
      str.Add(padding + newStr + padding);
    }
    return str.ToArray();
  }
}
