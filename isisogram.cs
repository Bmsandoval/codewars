using System;
using System.Collections.Generic;

public class Kata
{
  public static bool IsIsogram(string str) 
  {
    string s = str.ToLower();
    var set = new HashSet<char>();
    for (int i = 0; i < s.Length; i++)
        if (!set.Add(s[i]))
            return false;
    return true;
  }
}
