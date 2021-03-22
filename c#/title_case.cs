using System.Collections.Generic;
using System;

public class Kata
{
  public static string TitleCase(string title, string minorWords="")
  {
    string[] _minorWords = null;
    if (minorWords != null) {
      _minorWords = minorWords.ToLower().Split(' ');
    }
    string[] _title = title.ToLower().Split(' ');
    List<string> _result = new List<string>();
    
    bool first = true;
    foreach (string s in _title) {
      string _word = s;
      if (_word.Length < 1) {
        continue;
      }
      else if ( first ) {
        _word = char.ToUpper(_word[0]) + _word.Remove(0,1).ToLower();
        first = false;
      }
      else if (_minorWords != null && Array.Exists(_minorWords, element => element == s)) {
        _word=_word.ToLower();
      } 
      else {
        _word = char.ToUpper(_word[0]) + _word.Remove(0,1).ToLower();
      }
      _result.Add(_word);
    }
    
    return string.Join(" ", _result);
  }
}
