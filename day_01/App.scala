import scala.io.Source

@main def part1(filename: String): Unit = println(s"Answer to part 1 is: ${part1(loadData(filename))}")

@main def part2(filename: String): Unit = println(s"Answer to part 2 is: ${part2(loadData(filename))}")

def loadData(filename: String): Iterator[String] = Source.fromFile(filename).getLines

def part1(lines: Iterator[String]): String = 
    def getNumber(s: String): Int = 
        val firstDigit = s.find(_.isDigit).get.toInt - 48
        val lastDigit = s.findLast(_.isDigit).get.toInt - 48
        firstDigit * 10 + lastDigit

    lines.map(getNumber(_)).sum.toString()

def part2(lines: Iterator[String]): String = 
    def isNumber(s: String): Option[Int] = 
        if (s.isEmpty()) None
        else if (s.charAt(0).isDigit) Some(s.charAt(0).toInt - 48)
        else if (s.startsWith("zero")) Some(0)
        else if (s.startsWith("one")) Some(1)
        else if (s.startsWith("two")) Some(2)
        else if (s.startsWith("three")) Some(3)
        else if (s.startsWith("four")) Some(4)
        else if (s.startsWith("five")) Some(5)
        else if (s.startsWith("six")) Some(6)
        else if (s.startsWith("seven")) Some(7)
        else if (s.startsWith("eight")) Some(8)
        else if (s.startsWith("nine")) Some(9)
        else None

    def getNumber(s: String): Int =
        val matches = s.tails.flatMap(isNumber(_)).toList
        val firstDigit = matches.head
        val lastDigit = matches.last
        firstDigit * 10 + lastDigit

    lines.map(getNumber(_)).sum.toString()


    