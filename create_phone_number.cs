using System.Linq;

public class Kata
{
  public static string CreatePhoneNumber(int[] numbers)
  {
    int sub1 = int.Parse(numbers.Take(3).Aggregate("", (current, next) => current + next.ToString()));
    int sub2 = int.Parse(numbers.Skip(3).Take(3).Aggregate("", (current, next) => current + next.ToString()));
    int sub3 = int.Parse(numbers.Skip(6).Take(4).Aggregate("", (current, next) => current + next.ToString()));
    return $"({sub1:000}) {sub2:000}-{sub3:0000}";
  }
}
