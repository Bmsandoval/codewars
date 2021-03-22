using System;
using System.Text.RegularExpressions;

public class Kata
{
  public static string PigIt(string str)
  {
    Regex r = new Regex(@"^[a-zA-Z][a-z]+$"); 
    string[] spl = str.Split(' ');
    string newStr = "";
    foreach (string s in spl) {
      string n = s;
      if (r.IsMatch(n)) {
        n = s.Substring(1,s.Length-1) + s.Substring(0,1) + "ay";
      }
      newStr += " " + n;
    }
    return newStr.Trim(' ');
  }
}
