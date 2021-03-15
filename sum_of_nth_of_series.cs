using System;
using System.Linq;

public class NthSeries {
	
	public static string seriesSum (int n) {
		double sum = 0.0;
    foreach (int i in Enumerable.Range(0, n)) {
      sum += 1.0/(1.0+i*3);
    }
    return string.Format("{0:N2}", sum);
	}
}
