using System.Collections;
using System.Collections.Generic;

public class ListFilterer
{
   public static IEnumerable<int> GetIntegersFromList(List<object> listOfItems)
   {
     var l = new List<int>();
     foreach (var o in listOfItems) {
       if (o is int) {
         l.Add((int)o);
       }
     }
     return l;
   }
}
