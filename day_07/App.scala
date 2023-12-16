import scala.io.Source
import scala.util.control.Breaks._

@main def part1(filename: String): Unit = println(s"Answer to part 1 is: ${part1(loadData(filename))}")

@main def part2(filename: String): Unit = println(s"Answer to part 2 is: ${part2(loadData(filename))}")

def loadData(filename: String): Iterator[String] = Source.fromFile(filename).getLines

object HandType extends Enumeration {
    type HandType = Value
    val HighCard = Value(0)
    val OnePair = Value(1)
    val TwoPair = Value(2)
    val ThreeKind = Value(3)
    val FullHouse = Value(4)
    val FourKind = Value(5)
    val FiveKind = Value(6)
}

class Hand(var hand: String, var value: Long) {
    def getTypePart1(): HandType.HandType = 
        getType(this.hand.groupBy(duplicates(hand, _)).map(set))

    def set(m: (Int, String)): (Int, Set[Char]) =
        val v = m._2.toSet
        (m._1, v)

    def getTypePart2(): HandType.HandType = 
        val m = this.hand.replace("J", "").groupBy(duplicates(hand, _)).map(set)
        var hType = getType(m)
        if (hand == "JJJJJ")
            hType = HandType.FiveKind
        else if (this.hand.count(_ == 'J') > 0)
            var max = m.get(m.keySet.max).get
            val c = max.toList(0)
            val newHand = hand.replace('J', c)
            hType = getType(newHand.groupBy(duplicates(newHand, _)).map(set))
        hType

    def duplicates(s: String, c :Char): Int =
        s.count(_ == c)

    def getType(m: Map[Int, Set[Char]]): HandType.HandType = 
        if (m.contains(5))
            HandType.FiveKind
        else if (m.contains(4))
            HandType.FourKind
        else if (m.contains(3) && m.contains(2))
            HandType.FullHouse
        else if (m.contains(3))
            HandType.ThreeKind
        else if (m.contains(2) && m.get(2).get.size == 2)
            HandType.TwoPair
        else if (m.contains(2))
            HandType.OnePair
        else 
            HandType.HighCard
}

def part1(lines: Iterator[String]): String = 
    val l = lines.map(toHand).toList.groupBy(_.getTypePart1()).toList.sortBy(_._1).map(_._2)
    val cards = Map('A' -> 14, 'K' -> 13, 'Q' -> 12, 'J' -> 11, 'T' -> 10)
    var start = 1
    var total = 0.toLong
    for (i <- l)
        val v = getValues(i, start, cards)
        start = start + v.length
        total = total + v.sum
    total.toString()

    
def part2(lines: Iterator[String]): String = 
    val l = lines.map(toHand).toList.groupBy(_.getTypePart2()).toList.sortBy(_._1).map(_._2)
    val cards = Map('A' -> 14, 'K' -> 13, 'Q' -> 12, 'J' -> 1, 'T' -> 10)
    var start = 1
    var total = 0.toLong
    for (i <- l)
        val v = getValues(i, start, cards)
        start = start + v.length
        total = total + v.sum
    total.toString()

def toHand(s: String): Hand =
    val words = s.split(" ")
    val hand = words(0)
    val value = words(1).toLong
    new Hand(hand, value)

def getValues(hands: List[Hand], start: Int, cards: Map[Char, Int]): List[Long] =
    hands.sortWith(handSort(_, _, cards)).zipWithIndex.map{ case (v, i) => (i.toLong + start) * v.value }

def handSort(h1: Hand, h2: Hand, cards: Map[Char, Int]): Boolean =
    val s1 = h1.hand
    val s2 = h2.hand
    var b = false
    breakable {
        for (s <- 0 until s1.length()) 
            if (s1.charAt(s) != s2.charAt(s))
                val n1 = if (s1.charAt(s).isDigit) s1.charAt(s).toInt - 48 else cards.get(s1.charAt(s)).get
                val n2 = if (s2.charAt(s).isDigit) s2.charAt(s).toInt - 48 else cards.get(s2.charAt(s)).get
                b = n1 < n2
                break
    }
    b